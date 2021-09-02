package process

import "github.com/uraven0107/kvsHogeHoge/alias"

type QueryTokenGenerator struct{}

func (QueryTokenGenerator) Generate(runes *alias.Runes) alias.Tokens {
	tokens := alias.Tokens{}
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
