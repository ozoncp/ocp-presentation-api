// Package utils implement a simple internal library for Ozon Code Platform Presentation API.
package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type reverseMapInput struct {
	one map[string]string
}

type reverseMapOutput struct {
	one map[string]string
}

type reverseMapTestCase struct {
	in  reverseMapInput
	out reverseMapOutput
}

var reverseMapTestCases = []reverseMapTestCase{
	{
		in: reverseMapInput{
			one: map[string]string{},
		},
		out: reverseMapOutput{
			one: map[string]string{},
		},
	},
	{
		in: reverseMapInput{
			one: map[string]string{
				"C":          "1972",
				"SQL":        "1978",
				"C++":        "1980",
				"Python":     "1990",
				"Java":       "1995",
				"Go":         "2009",
				"Rust":       "2010",
				"Kotlin":     "2011",
				"TypeScript": "2012",
			},
		},
		out: reverseMapOutput{
			one: map[string]string{
				"1972": "C",
				"1978": "SQL",
				"1980": "C++",
				"1990": "Python",
				"1995": "Java",
				"2009": "Go",
				"2010": "Rust",
				"2011": "Kotlin",
				"2012": "TypeScript",
			},
		},
	},
}

func TestReverseMap(t *testing.T) {
	ast := assert.New(t)

	for _, testCase := range reverseMapTestCases {
		out, in := testCase.out, testCase.in
		ast.Equal(out.one, ReverseMap(in.one), "Test Case: %v %v", in, out)
	}
}
