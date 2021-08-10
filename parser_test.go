package main

import (
	"testing"
)

func Test_Expr(t *testing.T) {
	tests := []struct {
		name string
		want []*DatastoreSource
	}{
		{
			name: "canParse",
			want: []*DatastoreSource{
				&DatastoreSource{
					name: "test",
					k_v_map: map[string]string{
						"hoge": "fuga",
						"foo":  "bar",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			par := NewParser(NewTokenizer("test={hoge=fuga;foo=bar;};"))
			got, err := par.Expr()
			if err != nil {
				t.Errorf("Expr() error = %v", err)
				return
			}
			for _, g := range got {
				is_matched := false
				for _, w := range tt.want {
					if g.name == w.name {
						is_matched = true
						for k, v := range w.k_v_map {
							if g.k_v_map[k] != v {
								t.Errorf("Expr() want = %v, but got = %v", v, g.k_v_map[k])
							}
						}
					}
				}
				if !is_matched {
					t.Errorf("Nothing matched. Expr() is not working")
				}
			}
		})
	}
}
