# Payments API DEMO

This is a demo project where I applied Hexagonal Architecture and DDD on a Golang project.

## Requirements
* Golang v1.18
* Docker
* Docker Compose

## Run Tests
1. `docker-compose up -d`
2. `go test ./... -v`

## ToDo List
- [x] Initial setup
- [ ] Basic auth authorization
- [x] Find balance use case
- [ ] Update balance based on new transactions
- [x] Create transaction use case
- [ ] List transactions use case
- [x] MySQL setup
- [ ] Think on how to monitor the application
- [ ] Think on how to send logs to Prometheus
- [ ] Command/Query Bus implementation

## Improvements
* JWT Tokens
* JSON:API