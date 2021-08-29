package main

import "errors"

type Tokenizer struct {
	tokens Tokens
	p      int
}

type Tokenizer_type int

const (
	Type_DS Tokenizer_type = iota
	Type_Query
)

func NewTokenizer(tokenizer_type Tokenizer_type, source string) (*Tokenizer, error) {
	tokens := convertStringToTokens(tokenizer_type, source)

	if len(tokens) == 0 {
		return nil, errors.New("token length is 0, Tokenizer.NewTokenizer() received invalid string = " + source)
	}

	return &Tokenizer{
		tokens: tokens,
		p:      0,
	}, nil
}

func convertStringToTokens(tokenizer_type Tokenizer_type, source string) Tokens {
	// 文字列からトークンスライス生成する
	runes := []rune(source)
	token_generator := NewTokenGenerator(tokenizer_type)
	return token_generator.Generate(&runes)
}

func (t *Tokenizer) Next() (string, error) {
	if !t.HasNext() {
		return "", errors.New(":( No such Element")
	}
	str := t.tokens[t.p]
	t.p = t.p + 1
	return str, nil
}

func (t *Tokenizer) HasNext() bool {
	return len(t.tokens) >= t.p+1
}

func (t *Tokenizer) HasPrev() bool {
	return t.p-1 >= 0
}

func (t *Tokenizer) Prev() (string, error) {
	if !t.HasPrev() {
		return "", errors.New(":( No such Element")
	}
	t.p = t.p - 1
	str := t.tokens[t.p]
	return str, nil
}
