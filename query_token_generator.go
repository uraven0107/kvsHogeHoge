package main

type QueryTokenGenerator struct{}

func (QueryTokenGenerator) Generate(runes *[]rune) Tokens {
	tokens := Tokens{}
	str := ""
	for _, r := range *runes {
		switch s := string(r); s {
		case " ":
			if str != "" {
				tokens = append(tokens, str)
				str = ""
			}
		default:
			str = str + s
		}
	}
	if str != "" {
		tokens = append(tokens, str)
	}
	return tokens
}
