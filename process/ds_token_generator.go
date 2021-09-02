package process

import (
	"github.com/uraven0107/kvsHogeHoge/alias"
)

type DSTokenGenerator struct{}

func (DSTokenGenerator) Generate(runes *alias.Runes) alias.Tokens {
	tokens := alias.Tokens{}
	str := ""
	for _, r := range *runes {
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
	return tokens
}
