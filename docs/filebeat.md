## Filebeat

이 문서는 Filebeat Docker 컨테이너에서 사용하는 `filebeat.yaml` 설정 파일의 내용에 대해 설명합니다.

### processors

1. **timestamp**
    - 이 프로세서는 로그 이벤트의 시간 필드를 지정하여 Elasticsearch(또는 Kibana)에서 시간 기반 검색/필터링이 가능하도록 합니다.
    - `layouts` 항목에 Go 언어 특유의 시간 포맷(예: `2006-01-02T15:04:05.000Z07:00`)을 지정하여, ISO8601/RFC3339 형태의 날짜/시간 문자열을 파싱하도록 설정합니다.

2. **rename**
    - 특정 JSON 필드를 Elasticsearch/Kibana로 전송하기 전에 다른 이름으로 바꿀 수 있도록 설정합니다.
    - `message` 필드는 Kibana UI에서 기본적으로 로그 본문을 표시하는 예약 필드입니다. 다른 JSON 필드의 값을 `message`로 바꿔서 Kibana에서 바로 보이도록 하고 싶다면 아래와 같이 설정합니다.

    ```yaml
    - rename:
        fields:
          - from: "msg"
            to: "message"
        ignore_missing: true
    ```
