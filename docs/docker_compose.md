# Docker Compose

이 Application은 `docker-compose.yaml` 파일을 사용하여 로컬(Local) 환경에서 다음과 같은 세 가지 컨테이너를 배포 및 동작시킵니다:

1. **Elasticsearch**
2. **Kibana**
3. **Filebeat**

아래는 각 컨테이너별 설정 내용과 사용 방법에 대한 설명입니다.

---

## 1. Elasticsearch

- **이미지**
    - `docker.elastic.co/elasticsearch/elasticsearch:8.4.0`
- **환경 변수**
    - `discovery.type=single-node`
        - 단일 노드로 실행 (개발/테스트용)
    - `xpack.security.enabled=false`
        - 보안 설정 비활성화 (테스트용)
    - `ES_JAVA_OPTS=-Xms512m -Xmx512m`
        - 메모리를 최소/최대 512MB로 제한
- **포트**
    - `9200:9200`, `9300:9300`
    - 로컬에서 9200과 9300 포트를 개방
- **네트워크**
    - `elastic` 네트워크에 연결
- **재시작 정책**
    - `restart: always`

이 컨테이너는 **Elasticsearch**를 구동하며, Filebeat와 Kibana가 이를 참조하여 로그를 전송하거나 시각화할 수 있도록 합니다.

---

## 2. Kibana

- **이미지**
    - `docker.elastic.co/kibana/kibana:8.4.0`
- **환경 변수**
    - `ELASTICSEARCH_HOSTS=http://elasticsearch:9200`
        - Kibana가 Elasticsearch 컨테이너(`elasticsearch`)에 연결하도록 설정
- **포트**
    - `5601:5601`
    - 로컬에서 5601 포트를 개방 → [http://localhost:5601](http://localhost:5601)에서 Kibana UI 접근 가능
- **종속성**
    - `depends_on: elasticsearch`
        - Elasticsearch가 먼저 시작되어야 함
- **네트워크**
    - `elastic` 네트워크에 연결
- **재시작 정책**
    - `restart: always`

**Kibana**는 Elasticsearch에 적재된 로그 및 데이터를 시각화/분석하기 위한 UI를 제공합니다.

---

## 3. Filebeat

- **이미지**
    - `docker.elastic.co/beats/filebeat:8.4.0`
- **컨테이너 이름**
    - `filebeat`
- **유저 권한**
    - `root`
- **볼륨(Volumes)**
  ```yaml
  volumes:
    - ./filebeat.yml:/usr/share/filebeat/filebeat.yml:ro
    - D:/00_Development/Log-Sentinel/remote_logs:/logs:ro
  ```
  
