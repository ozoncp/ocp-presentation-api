package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseMap(t *testing.T) {
	ast := assert.New(t)

	type input struct {
		one map[string]string
	}

	type output struct {
		one map[string]string
	}

	type question struct {
		in  input
		out output
	}

	questions := []question{
		{
			in: input{
				one: map[string]string{},
			},
			out: output{
				one: map[string]string{},
			},
		},
		{
			in: input{
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
			out: output{
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

	for _, q := range questions {
		out, in := q.out, q.in
		ast.Equal(out.one, ReverseMap(in.one), "Test Case: %v %v", in, out)
	}
}
