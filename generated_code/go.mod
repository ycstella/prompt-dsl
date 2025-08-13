module workflow

go 1.24.4

require service v0.0.0

require github.com/sashabaranov/go-openai v1.40.5 // indirect

replace service => ../service

replace codegen => ../codegen
