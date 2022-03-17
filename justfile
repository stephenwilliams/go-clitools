default: generate format

format:
    #!/usr/bin/env bash
    go fmt ./...
    cd _build/toolgen
    go fmt ./...

generate:
    #!/usr/bin/env bash
    cd _build/toolgen
    go run main.go
