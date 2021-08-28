package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueryTokenGenerate(t *testing.T) {
	t.Run("canGenerateQueryToken", func(t *testing.T) {
		query := "use hoge"
		expected := []string{"use", "hoge"}
		g := QueryTokenGenerator{}
		runes := []rune(query)
		r_p := &runes
		got := g.Generate(r_p)
		assert := assert.New(t)
		assert.Equal(expected, got, fmt.Sprintf("Tokens not equal, expected = %v, but got = %v", expected, got))
	})
}
