apps = 'api' 'auth'

.PHONY: wire
wire:
	wire ./...
.PHONY: clean-image
clean-image:
	docker-compose stop
	docker-compose rm -f
	docker images | grep caas-micro | awk '{cmd="docker rmi "\$1":"\$2;system(cmd)}'

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o cmd/api/api -a -installsuffix cgo -ldflags '-w -s' ./cmd/api
