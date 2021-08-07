package main

import (
	"testing"
)

func Test_next(t *testing.T) {
	tests := []struct {
		name  string
		wants []string
	}{
		{
			name: "canIterateToken",
			wants: []string{
				"test",
				"=",
				"{",
				"hoge",
				"=",
				"fuga",
				";",
				"foo",
				"=",
				"bar",
				";",
				"}",
				";",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := "test={hoge=fuga;foo=bar;};"
			tokenizer := NewTokenizer(source)
			for _, want := range tt.wants {
				if !tokenizer.hasNext() {
					t.Errorf("hasNext() should return true.")
				}
				got, err := tokenizer.next()
				if err != nil {
					t.Errorf("Error has occured at next(). error = %v", err)
				}
				if got != want {
					t.Errorf("next() = %v, want %v", got, want)
				}
			}
			if tokenizer.hasNext() {
				t.Errorf("hasNext() should return false.")
			}

		})
	}
}
