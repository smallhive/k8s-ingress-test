REGISTRY_PATH?=belazar13/tic-tak-toe
RELEASE=echo-service-3

GO_CMD=CGO_ENABLED=0 GO111MODULE=on go
GO_PLATFORM_FLAGS=GOOS=linux GOARCH=amd64
LDFLAGS:=-s -w

build:
	$(GO_PLATFORM_FLAGS) $(GO_CMD) build -ldflags "$(LDFLAGS)" -o ./echo-service ./cmd/echo

docker-build:
	docker build -t ${REGISTRY_PATH}:${RELEASE} .

docker-push:
	docker push ${REGISTRY_PATH}:${RELEASE}
