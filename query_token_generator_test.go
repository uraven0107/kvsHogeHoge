package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueryTokenGenerate(t *testing.T) {
	t.Run("canGenerateQueryToken", func(t *testing.T) {
		assert := assert.New(t)
		query := "use hoge fuga"
		expected := Tokens{"use", "hoge", "fuga"}
		g := QueryTokenGenerator{}
		runes := []rune(query)
		r_p := &runes
		got := g.Generate(r_p)
		assert.Equal(expected, got, fmt.Sprintf("Tokens not equal, expected = %v, but got = %v", expected, got))
	})
}
