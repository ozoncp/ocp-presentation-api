package utils

import (
	"testing"

	"github.com/ozoncp/ocp-presentation-api/internal/models"
	"github.com/stretchr/testify/assert"
)

type filterSliceInput struct {
	one []models.Presentation
	two []models.Presentation
}

type filterSliceOutput struct {
	one []models.Presentation
}

type filterSliceTestCase struct {
	in  filterSliceInput
	out filterSliceOutput
}

var filterSliceTestCases = []filterSliceTestCase{
	{
		in: filterSliceInput{
			one: nil,
			two: nil,
		},
		out: filterSliceOutput{
			one: nil,
		},
	},
	{
		in: filterSliceInput{
			one: []models.Presentation{},
			two: nil,
		},
		out: filterSliceOutput{
			one: []models.Presentation{},
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
