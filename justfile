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


install-tool-deps:
    go install github.com/atombender/go-jsonschema/cmd/gojsonschema@latest

gen-tool-schema: install-tool-deps
    gojsonschema ./_build/tool.schema.json --package tools  -o ./_build/toolgen/tools/types.go
