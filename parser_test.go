package main

import (
	"testing"
)

func Test_Expr(t *testing.T) {
	tests := []struct {
		name  string
		wants []*DatastoreSource
	}{
		{
			name: "canParse",
			wants: []*DatastoreSource{
				&DatastoreSource{
					name: "test",
					k_v_map: map[string]string{
						"hoge": "fuga",
						"foo":  "bar",
					},
				},
				&DatastoreSource{
					name: "baka",
					k_v_map: map[string]string{
						"aho":  "manuke",
						"unko": "brbr",
						"aaa":  "bbb",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer, err := NewTokenizer(Type_DS, "test={hoge=fuga;foo=bar;};baka={aho=manuke;unko=brbr;aaa=bbb;};")
			if err != nil {
				t.Errorf(":( Error has occured at NewTokenizer(), error = %v", err)
				return
			}
			par := NewParser(tokenizer)
			gots, err := par.Expr()
			if err != nil {
				t.Errorf("Expr() error = %v", err)
				return
			}
			for _, want := range tt.wants {
				is_matched := false
				for _, got := range gots {
					if got.name == want.name {
						is_matched = true
						for k, v := range want.k_v_map {
							if got.k_v_map[k] != v {
								t.Errorf("Expr() want = %v, but got = %v", v, got.k_v_map[k])
							}
						}
					}
				}
				if !is_matched {
					t.Errorf("Nothing matched name = %v. Expr() is not working", want.name)
				}
			}
		})
	}
}
