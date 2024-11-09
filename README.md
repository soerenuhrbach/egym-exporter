# VisualStudio Marketplace Prometheus Exporter

[![Build](https://github.com/soerenuhrbach/egym-exporter/actions/workflows/ci.yml/badge.svg)](https://github.com/soerenuhrbach/egym-exporter/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/soerenuhrbach/egym-exporter?status.png)](https://godoc.org/github.com/soerenuhrbach/egym-exporter)
[![GoReportCard](https://goreportcard.com/badge/github.com/soerenuhrbach/egym-exporter)](https://goreportcard.com/report/github.com/soerenuhrbach/egym-exporter)

Prometheus exporter to scrape metrics from egym.

## Usage

### Using release binaries 

You can download the latest version of the binary built for your architecture here:

* Architecture **i386** [
    [Linux](https://github.com/soerenuhrbach/egym-exporter/releases/latest/download/egym_exporter-linux-386) /
    [Windows](https://github.com/soerenuhrbach/egym-exporter/releases/latest/download/egym_exporter-windows-386.exe)
]
* Architecture **amd64** [
    [Darwin](https://github.com/soerenuhrbach/egym-exporter/releases/latest/download/egym_exporter-darwin-amd64) /
    [Linux](https://github.com/soerenuhrbach/egym-exporter/releases/latest/download/egym_exporter-linux-amd64) /
    [Windows](https://github.com/soerenuhrbach/egym-exporter/releases/latest/download/egym_exporter-windows-amd64.exe)
]
* Architecture **arm** [
    [Darwin](https://github.com/soerenuhrbach/egym-exporter/releases/latest/download/egym_exporter-darwin-arm64) /
    [Linux](https://github.com/soerenuhrbach/egym-exporter/releases/latest/download/egym_exporter-linux-arm)
]

You can run it using the following example:

```bash
./egym_exporter -egym-brand "<brand>" -egym-username "<username>" -egym-password "<password>"
```

### Docker 

The exporter is also available as a [Docker image](https://github.com/soerenuhrbach/egym-exporter/pkgs/container/egym-exporter).
You can run it using the following example and pass configuration environment variables:

```bash
docker run \
  -e 'EGYM_BRAND=<brand>' \
  -e 'EGYM_USERNAME=<brand>' \
  -e 'EGYM_PASSWORD=<password>' \
  -p 9391:9391 \
  ghcr.io/soerenuhrbach/egym-exporter:latest
```

## Scrape metrics

Once the exporter is running, you can configure your collector to scrape the metrics. 

Adjust your `prometheus.yml` configuration to let it scrape the exporter like this:

```yaml
scrape_configs:
  - job_name: 'egym'
    static_configs:
      - targets: ['localhost:9391']
```

## Available configurations

|Configuration name|Description|Required|Argument|Environment variable|Default|
|---|---|---|---|---|---|
|Extensions|Comma-separated list of extensions that should be scraped|Required|extensions|EXTENSIONS|-|
|Metric path|Port to be used for the exporter|false|metricsPath|METRICSPATH|/metrics|
|Port|Port to be used for the exporter|false|port|PORT|9391|
|Binding Address|Address to be used for the exporter|false|bindAddress|BINDADDRESS|0.0.0.0|

Examples with all possible configurations:

```bash
./egym_exporter \
  -egym-brand "<brand>" \ 
  -egym-username "<username>" \
  -egym-password "<password>" \
  -metricsPath "/metrics" \
  -port 9391 \
  -bindAddress 0.0.0.0
```
or using docker:

```bash
docker run \
  -e 'EGYM_BRAND=<brand>' \
  -e 'EGYM_USERNAME=<brand>' \
  -e 'EGYM_PASSWORD=<password>' \
  -e 'METRICSPATH=/metrics' \
  -e 'PORT=9391' \
  -e 'BINDADDRESS=0.0.0.0' \
  -p 9391:9391 \
  ghcr.io/soerenuhrbach/egym-exporter:latest
```

## Available metrics

|Metric|Description|
|---|---|
|egym_bio_age_total|Total bio age|
|egym_bio_age_cardio|Cardio bio age|
|egym_bio_age_flexibility|Flexibility bio age|
|egym_bio_age_metabolic|Metabolic bio age|
|egym_bio_age_muscles|Total muscles bio age|
|egym_bio_age_muscles_core|Core muscles bio age|
|egym_bio_age_muscles_lower_body|Lower body muscles bio age|
|egym_bio_age_muscles_upper_body|Upper body muscles bio age|
|egym_activity_points|Amount of activity points within the current period|
|egym_activity_points_goal|Amount of activity points to reach the next activity level|
|egym_activity_maintain_points|Amount of required activity points to maintain the activity level|
|egym_body|Various metrics about the body like BMI, muscles or fat.|
|egym_strength|Strength metrics about various muscles and muscle groups.|
|egym_flexibility|Flexibility metrics about different body parts like neck or hips.|
|egym_muscle_imbalance|Proportions and (im)balances of different muscle pairs.|
|egym_exercise_activity_points|Collected activity points of a specific exercise|
|egym_exercise_activity_distance|Distance left behind with the specific exercise|
|egym_exercise_activity_duration|Duration of a specific exercise|
|egym_exercise_activity_calories|Burned calories with a specific exercise|
|egym_exercise_average_speed|Average speed within a specific exercise|
|egym_exercise_sets|Number of sets within a specific exercise|
|egym_exercise_reps|number of reps over all sets of a specific exercise|
|egym_exercise_weight_total|Total weight across all reps of all sets of the exercise|

## Release Notes

### 0.7.0

Added metrics about exercises. 

### 0.6.0

Added muscle imbalance metrics.

### 0.5.0

Added flexibility metrics about different body parts like neck or hips.

### 0.4.0

Added strength metrics about various muscles and muscle groups.

### 0.3.0

Added body metrics 

### 0.2.0

Added metrics about the users activity level

### 0.1.0

Initial release

[MIT LICENSE](LICENSE)