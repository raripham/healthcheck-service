# healthcheck-service

```bash
cd healthcheck-service/

go get .
go run .

```

Register service

```bash
curl -d '{
    "service_name": "abcos",
    "status": "Up",
    "metadata": {
        "server": "34"
    }
}' -H "Content-Type: application/json" -X POST http://localhost:8080/api
```