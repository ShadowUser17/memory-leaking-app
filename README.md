#### Validate project files:
```bash
golangci-lint run
```

#### Build binary file:
```bash
go mod tidy
```
```bash
go build -ldflags="-s -w" -o ./app ./cmd/main.go
```

#### Build docker image:
- Stable version:
```bash
docker build -t "shadowuser17/memory-leaking-app:latest" .
```

#### Scan docker image:
```bash
dockle "shadowuser17/memory-leaking-app:latest"
```
```bash
trivy image "shadowuser17/memory-leaking-app:latest"
```

#### Publish docker image:
```bash
docker login -u "${DOCKERHUB_LOGIN}" -p "${DOCKERHUB_TOKEN}"
```
```bash
docker push "shadowuser17/memory-leaking-app:latest"
```
