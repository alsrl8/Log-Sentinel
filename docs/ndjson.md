# NDJSON File Format

## Description

--- 

- Elastic Search와 fibana를 활용한 log 수집 및 분석을 위해 기존의 `.log` 형식으로 유지되던 log file을 체계화 된 format으로 변경
- ndjson format으로 변경하여 아래와 같은 내용을 쉽게 파악할 수 있도록 변경
    - env
        - dev, prod, on-prem
    - timestamp
        - 로그 생성 일시
    - lvl
        - logging level
        - DEBUG, INFO, ERROR
    - msg
        - fibana ui에 title로 출력할 문자열
    - svc
        - dclo-api
        - dclo-admin-api
        - dclo-key-validation
    - tb
        - (Optional)
        - error trace back
        - 에러, 디버깅 용도로 사용
    - user_id, tenant_id
        - (Optional)
        - API Server의 경우 요청한 고객을 특정하기 위해 사용
    - request_id
        - (Optional)
        - API Server의 경우 API Call 마다 unique id를 할당하여 특정하기 위해 사용

## Examples

---

```json
{
  "env": "dev",
  "timestamp": "2025-03-12T01:41:09",
  "lvl": "DEBUG",
  "msg": "Publishing message to exchange:cspm with routing key(engine)",
  "svc": "dclo-api",
  "user_id": "dragon",
  "tenant_id": "dragon",
  "request_id": "4b6835d5-0936-45e8-b142-bb89219aed6a"
},
{"env": "dev", "timestamp": "2025-03-12T01:41:09", "lvl": "DEBUG", "msg": "Publishing message to exchange:cspm with routing key(engine)", "svc": "dclo-api", "user_id": "dragon", "tenant_id": "dragon", "request_id": "4b6835d5-0936-45e8-b142-bb89219aed6a"}
```

위 예시는 가독성을 위해 여러 줄에 걸쳐 log message를 출력했지만, 실제 ndjson log file은 **하나의 log가 하나의 line 규칙**을 반드시 지킵니다.