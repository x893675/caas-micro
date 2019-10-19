apps = 'api' 'auth'

.PHONY: wire
wire:
	wire ./...
.PHONY: clean-image
clean-image:
	docker-compose stop
	docker-compose rm -f
	docker images | grep caas-micro | awk '{cmd="docker rmi "$$1":"$$2;system(cmd)}'

.PHONY: build
build:
	for app in $(apps) ;\
	do \
		CGO_ENABLED=0 go build -o dist/$$app -a -installsuffix cgo -ldflags '-w -s' ./cmd/$$app;\
	done

.PHONY: docker
docker:
	docker-compose up --build -d
