module github.com/stephenwilliams/go-clitools/_build/toolgen

go 1.15

require (
	github.com/dave/jennifer v1.4.1
	github.com/davecgh/go-spew v1.1.1
	github.com/kr/pretty v0.1.0 // indirect
	github.com/stephenwilliams/go-clitools v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/mod v0.3.0
	golang.org/x/tools v0.1.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/stephenwilliams/go-clitools => ../../
