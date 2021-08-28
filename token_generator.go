package main

type TokenGenerator interface {
	Generate(runes *[]rune) []string
}

func NewTokenGenerator(tokenizer_type Tokenizer_type) TokenGenerator {
	return newDSTokenGenerator()
}

func newDSTokenGenerator() TokenGenerator {
	return DSTokenGenerator{}
}
