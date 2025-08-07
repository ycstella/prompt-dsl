module web-backend

go 1.24.4

require (
	codegen v0.0.0
	service v0.0.0
)

require (
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/sashabaranov/go-openai v1.40.5 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
)

replace codegen => ../codegen

replace service => ../service
