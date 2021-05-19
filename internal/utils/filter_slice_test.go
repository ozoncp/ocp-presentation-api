package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterSlice(t *testing.T) {
	ast := assert.New(t)

	type input struct {
		one []string
		two []string
	}

	type output struct {
		one []string
	}

	type question struct {
		in  input
		out output
	}

	questions := []question{
		{
			in: input{
				one: []string{},
				two: []string{},
			},
			out: output{
				one: nil,
			},
		},
		{
			in: input{
				one: []string{""},
				two: []string{""},
			},
			out: output{
				one: nil,
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: []string{"a", "b", "c", "d", "e"},
			},
			out: output{
				one: nil,
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: nil,
			},
			out: output{
				one: []string{"a", "b", "c", "d", "e"},
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: []string{"a", "c", "e"},
			},
			out: output{
				one: []string{"b", "d"},
			},
		},
		{
			in: input{
				one: []string{"a", "a", "a", "b", "b"},
				two: []string{"a"},
			},
			out: output{
				one: []string{"b", "b"},
			},
		},
	}

	for _, q := range questions {
		out, in := q.out, q.in
		ast.Equal(out.one, FilterSlice(in.one, in.two), "Test Case: %v %v", in, out)
	}
}
