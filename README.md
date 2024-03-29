# Simple App

This repo is for presenting the clean architecture using golang. I separate the app into multiple layer. There is `service`,`usecase`,`application` (from bottom to top). The service layer will have responsibilities to store or get data from any resource such as `database`,`api` etc. The usecase layer will do any business logic by consuming the needed service. And application layer will serve any usecase using many app service like `http`,`cli`,`grpc` etc.

In short, `service` is responsible as a `data layer`, `usecase` is responsible as `business layer`, and the `application` is responsible as a `server layer`.

## Development Guideline

### Requirement

1. Docker
2. Go 1.14+

### Docker

This simple-app using docker for easy development, please make sure you already install the docker compose. Then just run:

```bash
~$ make dev
```

You can add the needed service by editing the `.dev/docker-compose.dev.yaml` files.

author: @reyhanfahlevi