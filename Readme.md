# 🚀 Log Generator (Go + Docker)

A lightweight (`790 kb`), configurable log generator built in Go.  
It emits structured logs continuously and is useful for testing log pipelines, monitoring systems, and ingestion services.

---

## 🧠 Overview

This service generates logs based on environment configuration.

It supports:
- Multiple log levels (INFO, WARN, ERROR, etc.)
- Custom messages
- Configurable delay between logs
- Service tagging

---

## ⚙️ Environment Variables

| Variable        | Description                                      | Default Value |
|----------------|--------------------------------------------------|--------------|
| `LOG_CONFIG`   | Comma-separated log definitions (`LEVEL:message`) | `INFO:welcome tester,WARN:don't push Node modules` |
| `LOG_DELAY_MS` | Delay between log emissions (in milliseconds)     | `2000` |
| `SERVICE_NAME` | Name of the service emitting logs                | `demo-log-service` |

---

## 🧩 LOG_CONFIG Format

```text
LEVEL  TIME  SERVICE_NAME : MESSAGE
```
---

## Run container
```bash
docker run -d \
  -e LOG_CONFIG="INFO:login success,ERROR:db fail,WARN:memory high" \
  -e LOG_DELAY_MS=1000 \
  -e SERVICE_NAME=auth-service \
  rushilab/logemitter:latest
  
```
---

## 🧱 Use Cases

- Testing log ingestion pipelines
- Simulating production logs
- Debugging monitoring systems
- Load testing logging infrastructure
---
## 🧠 Notes
1. Logs are written to stdout (Docker logs)
2. Suitable for containerized environments
3. Can be used with log agents (Fluent Bit, Vector, etc.)
