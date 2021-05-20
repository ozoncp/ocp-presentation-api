package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type filterSliceInput struct {
	one []string
	two []string
}

type filterSliceOutput struct {
	one []string
}

type filterSliceTestCase struct {
	in  filterSliceInput
	out filterSliceOutput
}

var filterSliceTestCases = []filterSliceTestCase{
	{
		in: filterSliceInput{
			one: []string{},
			two: []string{},
		},
		out: filterSliceOutput{
			one: nil,
		},
	},
	{
		in: filterSliceInput{
			one: []string{""},
			two: []string{""},
		},
		out: filterSliceOutput{
			one: nil,
		},
	},
	{
		in: filterSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: []string{"a", "b", "c", "d", "e"},
		},
		out: filterSliceOutput{
			one: nil,
		},
	},
	{
		in: filterSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: nil,
		},
		out: filterSliceOutput{
			one: []string{"a", "b", "c", "d", "e"},
		},
	},
	{
		in: filterSliceInput{
			one: []string{"a", "b", "c", "d", "e"},
			two: []string{"a", "c", "e"},
		},
		out: filterSliceOutput{
			one: []string{"b", "d"},
		},
	},
	{
		in: filterSliceInput{
			one: []string{"a", "a", "a", "b", "b"},
			two: []string{"a"},
		},
		out: filterSliceOutput{
			one: []string{"b", "b"},
		},
	},
}

func TestFilterSlice(t *testing.T) {
	ast := assert.New(t)

	for _, testCase := range filterSliceTestCases {
		out, in := testCase.out, testCase.in
		ast.Equal(out.one, FilterSlice(in.one, in.two), "Test Case: %v %v", in, out)
	}
}
