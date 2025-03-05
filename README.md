``` shell
go run cmd/web/start.go & curl -X POST -H "Content-type:application/json" --data "{\"expression\":\"2*2\"}" http://localhost:8080/api/v1/calculate & curl localhost:8080/api/v1/expressions
```
