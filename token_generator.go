package main

type TokenGenerator interface {
	Generate(runes *Runes) Tokens
}

func NewTokenGenerator(tokenizer_type Tokenizer_type) TokenGenerator {
	switch tokenizer_type {
	case Type_DS:
		return DSTokenGenerator{}
	default:
		return QueryTokenGenerator{}
	}
}
