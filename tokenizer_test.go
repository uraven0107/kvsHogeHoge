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
			tokenizer, err := NewTokenizer(Type_DS, source)
			if err != nil {
				t.Errorf(":( Error has occured at NewTokenizer(), error = %v", err)
			}
			for _, want := range tt.wants {
				if !tokenizer.HasNext() {
					t.Errorf("hasNext() should return true.")
				}
				got, err := tokenizer.Next()
				if err != nil {
					t.Errorf("Error has occured at next(). error = %v", err)
				}
				if got != want {
					t.Errorf("next() = %v, want %v", got, want)
				}
			}
			if tokenizer.HasNext() {
				t.Errorf("hasNext() should return false.")
			}

		})
	}
}

func Test_prev(t *testing.T) {
	tests := []struct {
		name  string
		wants []string
	}{
		{
			name: "canRevertToken",
			wants: []string{
				";",
				"}",
				";",
				"bar",
				"=",
				"foo",
				";",
				"fuga",
				"=",
				"hoge",
				"{",
				"=",
				"test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := "test={hoge=fuga;foo=bar;};"
			tokenizer, err := NewTokenizer(Type_DS, source)
			if err != nil {
				t.Errorf(":( Error has occured at NewTokenizer(), error = %v", err)
			}
			for tokenizer.HasNext() {
				tokenizer.Next()
			}
			for _, want := range tt.wants {
				if !tokenizer.HasPrev() {
					t.Errorf("hasPrev() should return true.")
				}
				got, err := tokenizer.Prev()
				if err != nil {
					t.Errorf("Error has occured at prev(). error = %v", err)
				}
				if got != want {
					t.Errorf("prev() = %v, want %v", got, want)
				}
			}
			if tokenizer.HasPrev() {
				t.Errorf("hasPrev() should return false.")
			}
		})
	}
}
