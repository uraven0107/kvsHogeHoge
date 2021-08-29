package main

type DSTokenGenerator struct{}

func (DSTokenGenerator) Generate(runes *[]rune) []string {
	tokens := []string{}
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
