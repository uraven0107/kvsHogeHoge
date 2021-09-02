package process

import (
	"errors"
)

type Parser struct {
	tokenizer *Tokenizer
}

type DatastoreSource struct {
	Name    string
	K_V_map map[string]string
}

func NewParser(tokenizer *Tokenizer) *Parser {
	return &Parser{tokenizer}
}

func (par *Parser) Expr() ([]*DatastoreSource, error) {
	var datastoreSource_list []*DatastoreSource = []*DatastoreSource{}

	name := ""
	expect := "*"
	eq_after_name := false
	is_key := true
	key := ""
	value := ""
	is_k_v_generated := false
	is_completed := false
	is_closed := false
	var datastoreSource *DatastoreSource = nil

	var initializer = func() {
		name = ""
		expect = "*"
		eq_after_name = false
		is_key = true
		key = ""
		value = ""
		is_k_v_generated = false
		is_closed = false
		is_completed = false
		datastoreSource = nil
	}
	t := par.tokenizer
	for t.HasNext() {
		token, err := t.Next()
		if err != nil {
			return nil, err
		}

		if token != expect && expect != "*" {
			return nil, errors.New("Parse error occured. expected = " + expect + ", but got = " + token)
		}

		if name == "" {
			name = token
			datastoreSource = &DatastoreSource{
				Name:    name,
				K_V_map: make(map[string]string),
			}
			expect = "="
			eq_after_name = true
			continue
		}

		switch token {
		case "=":
			if eq_after_name {
				expect = "{"
				eq_after_name = false
			} else {
				expect = "*"
			}
		case "{":
			expect = "*"
		case "}":
			expect = ";"
		case ";":
			if is_closed {
				is_completed = true
				break
			}
			next, err := t.Next()
			if err != nil {
				return nil, err
			}
			if next == "}" {
				expect = ";"
				is_closed = true
			} else {
				key = next
				is_key = false
				expect = "="
			}
		default:
			if is_key {
				key = token
				is_key = false
				expect = "="
			} else {
				value = token
				expect = ";"
				is_k_v_generated = true
			}
		}

		if is_k_v_generated {
			datastoreSource.K_V_map[key] = value
			is_k_v_generated = false
		}

		if is_completed {
			datastoreSource_list = append(datastoreSource_list, datastoreSource)
			initializer()
		}

	}
	return datastoreSource_list, nil
}
