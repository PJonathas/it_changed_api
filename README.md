CONTRIBUTING:
```
docker compose --env-file .env-dev up
```

post change:
```
curl -X POST -H "Content-Type: application/json" -d '{"description":"some manual change","type":"manual"}' localhost:8080/changes
```
