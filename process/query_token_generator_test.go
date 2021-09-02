package process

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uraven0107/kvsHogeHoge/alias"
)

func Test_QueryTokenGenerate(t *testing.T) {
	t.Run("canGenerateQueryToken", func(t *testing.T) {
		assert := assert.New(t)
		query := "use hoge fuga"
		expected := alias.Tokens{"use", "hoge", "fuga"}
		g := QueryTokenGenerator{}
		runes := alias.Runes(query)
		r_p := &runes
		got := g.Generate(r_p)
		assert.Equal(expected, got, fmt.Sprintf("Tokens not equal, expected = %v, but got = %v", expected, got))
	})
}
