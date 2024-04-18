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
