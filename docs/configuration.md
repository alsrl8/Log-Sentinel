## Configuration

이 문서는 **Log Sentinel Application**에서 설정 관리를 위해 사용하는 `config.yaml` 파일의 내용을 설명합니다.

## Kube

Kubernetes 환경에서 **로그 파일**을 수집하기 위한 설정 항목입니다.

### fields

1. **name**
   - **required**
   - 수집한 로그를 **로컬 디렉터리에 저장**할 때 사용할 **이름**을 지정합니다.
   - 예) `fetched_log`

2. **format**
   - **required**
   - 로컬 디렉터리에 저장될 **파일 형식**을 지정합니다.
   - 예) `ndjson`

3. **namespace**
   - **required**
   - 로그를 수집할 대상 Pod가 속한 **네임스페이스**를 입력합니다.
   - 예) `my_namespace`

4. **serviceLabel**
   - **required**
   - Pod에 할당된 **label**을 입력합니다.
   - 동일한 label을 가진 Pod가 여러 개 있을 경우, **현재 로직**상 **가장 먼저 조회된** Pod만 대상으로 로그를 수집합니다. (추후 개선 필요)
   - 예) `app=my-service`

5. **logPath**
   - **required**
   - Pod 내부에 존재하는 **로그 파일 경로**를 입력합니다.
   - 예) `/app/logs/mylog.log`
