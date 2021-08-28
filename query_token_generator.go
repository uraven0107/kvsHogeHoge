package main

type QueryTokenGenerator struct{}

func (QueryTokenGenerator) Generate(runes *[]rune) []string {
	tokens := make([]string, len(*runes)) // ルーンの数 >= 文字列の数
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
