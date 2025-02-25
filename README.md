# LogSentinel

이 문서는 LogSentinel Application의 기능과 설계, 빌드 방법 및 환경 설정에 대한 내용을 다룹니다.

## Overview

이 Application은 D-CLO Service의 Log를 수집 및 관리하고 알림을 전송하는 기능을 수행합니다.

## Config

### Environment variables

- `LOG_SENTINEL_CONFIG_PATH`
    - config file path

### Config file

- 이 Application은 설정 관리를 위해 `config.yaml` file을 사용합니다.
- [config.yaml](./docs/configuration.md)

## Features

1. [Kibana UI](http://localhost:5601)
    - 수집한 log file을 ElasticSearch와 kibana를 사용하여 분석 및 모니터링 합니다.

## [Docker-compose](./docs/docker_compose.md)

## [Filebeat](./docs/filebeat.md)

## Deployment

- docker compose up on detach mode
  ```shell
  docker-compose up -d
    ```
- docker compose down with removing volumes
    ```shell
    docker-compose down -v
    ```