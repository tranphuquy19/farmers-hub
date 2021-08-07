# Farmer Hub project

## Run

```bash
docker-compose up -d
go run .
```

## Build

`go build .`

## Build cross-platform

`bash ./build.sh`

```text
# platforms.txt
linux/amd64
linux/arm
darwin/amd64
darwin/arm64
windows/amd64
windows/arm
```

## Hot-reload

- Requires Node.js installed
- Install nodemon: `[sudo] npm install -g nodemon`
- Start in deamon: `nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run .`