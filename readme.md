a simple reverse proxy written in go

example for use:

```
go run . -whitelist mydomain.com -target http://192.168.1.244:2283 -agree
```