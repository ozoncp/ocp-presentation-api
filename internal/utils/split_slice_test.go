package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type splitSliceInput struct {
	one []string
	two int
}

type splitSliceOutput struct {
	one [][]string
}

type splitSliceTestCase struct {
	in  splitSliceInput
	out splitSliceOutput
}

var splitSliceTestCases = []splitSliceTestCase{
	{
		in: splitSliceInput{
			one: []string{},
			two: 0,
		},
		out: splitSliceOutput{
			one: nil,
		},
	},
	{
		in: splitSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: -1,
		},
		out: splitSliceOutput{
			one: nil,
		},
	},
	{
		in: splitSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: 0,
		},
		out: splitSliceOutput{
			one: nil,
		},
	},
	{
		in: splitSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: 1,
		},
		out: splitSliceOutput{
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
		in: splitSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: 2,
		},
		out: splitSliceOutput{
			one: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"e"},
			},
		},
	},
	{
		in: splitSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: 3,
		},
		out: splitSliceOutput{
			one: [][]string{
				{"a", "b", "c"},
				{"d", "e"},
			},
		},
	},
	{
		in: splitSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: 4,
		},
		out: splitSliceOutput{
			one: [][]string{
				{"a", "b", "c", "d"},
				{"e"},
			},
		},
	},
	{
		in: splitSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: 5,
		},
		out: splitSliceOutput{
			one: [][]string{
				{"a", "b", "c", "d", "e"},
			},
		},
	},
	{
		in: splitSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: 25,
		},
		out: splitSliceOutput{
			one: [][]string{
				{"a", "b", "c", "d", "e"},
			},
		},
	},
}

func TestSplitSlice(t *testing.T) {
	ast := assert.New(t)

	for _, testCase := range splitSliceTestCases {
		out, in := testCase.out, testCase.in
		ast.Equal(out.one, SplitSlice(in.one, in.two), "Test Case: %v %v", in, out)
	}
}
