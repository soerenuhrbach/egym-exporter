package egym

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

type EgymClient struct {
	Brand    string
	Username string
	Password string

	userId         string
	cookies        string
	defaultHeaders map[string]string
	apiUrl         string
	brandApiUrl    string

	httpClient *http.Client
}

func NewEgymClient(brand, username, password string) (*EgymClient, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}

	c := &EgymClient{
		Brand:    brand,
		Username: username,
		Password: password,
		defaultHeaders: map[string]string{
			"x-np-user-agent":  "clientType=MOBILE_DEVICE; devicePlatform=IOS; deviceUid=0B7F0E30-9598-43EF-8DA6-7018BD289B3C; applicationName=EGYM Fitness; applicationVersion=3.11; applicationVersionCode=853; containerName=NetpulseFitness;",
			"user-agent":       "NetpulseFitness/3.11 (com.netpulse.netpulsefitness; build:853; iOS 17.2.0) Alamofire/5.4.4",
			"x-np-app-version": "3.11",
			"Accept":           "application/json",
		},
		brandApiUrl: fmt.Sprintf("https://%s.netpulse.com", brand),
		apiUrl:      "https://mobile-api.int.api.egym.com",
		httpClient:  httpClient,
	}
	loggedIn, err := c.login()
	if err != nil || !loggedIn {
		log.Fatal("Login failed:", err)
		return nil, err
	}
	return c, nil
}

func (c *EgymClient) login() (bool, error) {
	data := url.Values{}
	data.Set("username", c.Username)
	data.Set("password", c.Password)

	hasLogin := c.userId != ""
	data.Set("relogin", fmt.Sprintf("%t", hasLogin))

	loginUrl := fmt.Sprintf("%s/np/exerciser/login", c.brandApiUrl)
	req, err := http.NewRequest("POST", loginUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return false, err
	}
	for k, v := range c.defaultHeaders {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var result map[string]interface{}
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &result)
		c.userId = result["uuid"].(string)
		c.cookies = resp.Header.Get("Set-Cookie")
		return true, nil
	}
	return false, nil
}

func (c *EgymClient) fetch(url string, retryCount int) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", c.cookies)
	for k, v := range c.defaultHeaders {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 || resp.StatusCode == 403 && retryCount > 0 {
		c.login()
		return c.fetch(url, retryCount-1)
	}

	return io.ReadAll(resp.Body)
}
