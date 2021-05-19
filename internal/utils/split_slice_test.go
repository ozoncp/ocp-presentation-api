package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitSlice(t *testing.T) {
	ast := assert.New(t)

	type input struct {
		one []string
		two int
	}

	type output struct {
		one [][]string
	}

	type question struct {
		in  input
		out output
	}

	questions := []question{
		{
			in: input{
				one: []string{},
				two: 0,
			},
			out: output{
				one: nil,
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: -1,
			},
			out: output{
				one: nil,
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: 0,
			},
			out: output{
				one: nil,
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: 1,
			},
			out: output{
				one: [][]string{
					{"a"},
					{"b"},
					{"c"},
					{"d"},
					{"e"},
				},
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: 2,
			},
			out: output{
				one: [][]string{
					{"a", "b"},
					{"c", "d"},
					{"e"},
				},
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: 3,
			},
			out: output{
				one: [][]string{
					{"a", "b", "c"},
					{"d", "e"},
				},
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: 4,
			},
			out: output{
				one: [][]string{
					{"a", "b", "c", "d"},
					{"e"},
				},
			},
		},
		{
			in: input{
				one: []string{"a", "b", "c", "d", "e"},
				two: 5,
			},
			out: output{
				one: [][]string{
					{"a", "b", "c", "d", "e"},
				},
			},
		},
	}

	for _, q := range questions {
		out, in := q.out, q.in
		ast.Equal(out.one, SplitSlice(in.one, in.two), "Test Case: %v %v", in, out)
	}
}
