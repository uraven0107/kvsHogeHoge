package main

import "errors"

type Tokenizer struct {
	tokens []string
	p      int
}

func NewTokenizer(source string) *Tokenizer {
	// 文字列からトークンスライス生成する
	runes := []rune(source)
	tokens := make([]string, len(runes)) // ルーンの数 >= 文字列の数
	str := ""
	for _, r := range runes {
		switch s := string(r); s {
		case "=", "{", "}", ";":
			if str != "" {
				tokens = append(tokens, str)
				str = ""
			}
			tokens = append(tokens, s)
		default:
			str = str + s
		}
	}

	// 空文字以外の文字列数を出す
	tokens_len := 0
	for _, t := range tokens {
		if t == "" {
			continue
		}
		tokens_len = tokens_len + 1
	}

	// トークンスライスから空文字を除去する
	new_tokens := make([]string, tokens_len)
	i := 0
	for _, t := range tokens {
		if t == "" {
			continue
		}
		new_tokens[i] = t
		i = i + 1
	}

	return &Tokenizer{
		tokens: new_tokens,
		p:      0,
	}
}

func (t *Tokenizer) next() (string, error) {
	if !t.hasNext() {
		return "", errors.New(":( No such Element")
	}
	str := t.tokens[t.p]
	t.p = t.p + 1
	return str, nil
}

func (t *Tokenizer) hasNext() bool {
	return len(t.tokens) >= t.p+1
}

func (t *Tokenizer) hasPrev() bool {
	return t.p-1 >= 0
}

func (t *Tokenizer) prev() (string, error) {
	if !t.hasPrev() {
		return "", errors.New(":( No such Element")
	}
	t.p = t.p - 1
	str := t.tokens[t.p]
	return str, nil
}
