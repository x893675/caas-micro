# build api
CGO_ENABLED=0 GOOS=linux go build -o cmd/api/api -a -installsuffix cgo -ldflags '-w' ./cmd/api/

# build auth srv
CGO_ENABLED=0 GOOS=linux go build -o cmd/auth/auth -a -installsuffix cgo -ldflags '-w' cmd/auth/auth.go
