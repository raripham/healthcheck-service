# healthcheck-service

```bash
cd healthcheck-service/

go get .
go run .

```

Register service

```bash
curl -d '{
    "service_name": "is-trongpt",
    "status": "Up",
    "node_name": "dev-ftech",
    "node_ip": "192.168.1.1",
    "node_metadata": {
        "env": "dev"
    },
    "service_metadata": {
        "env": "dev"
    }
}' -H "Content-Type: application/json" -X POST http://localhost:8080/api
```