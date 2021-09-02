package process

import "github.com/uraven0107/kvsHogeHoge/alias"

type TokenGenerator interface {
	Generate(runes *alias.Runes) alias.Tokens
}

func NewTokenGenerator(tokenizer_type Tokenizer_type) TokenGenerator {
	switch tokenizer_type {
	case Type_DS:
		return DSTokenGenerator{}
	default:
		return QueryTokenGenerator{}
	}
}
