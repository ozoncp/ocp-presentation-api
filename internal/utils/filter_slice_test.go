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
	{
		in: filterSliceInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []models.Presentation{{ID: 1}},
		},
		out: filterSliceOutput{
			one: []models.Presentation{{ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
		},
	},
	{
		in: filterSliceInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []models.Presentation{{ID: 1}, {ID: 2}},
		},
		out: filterSliceOutput{
			one: []models.Presentation{{ID: 3}, {ID: 4}, {ID: 5}},
		},
	},
	{
		in: filterSliceInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}},
		},
		out: filterSliceOutput{
			one: []models.Presentation{{ID: 4}, {ID: 5}},
		},
	},
	{
		in: filterSliceInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}},
		},
		out: filterSliceOutput{
			one: []models.Presentation{{ID: 5}},
		},
	},
	{
		in: filterSliceInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
		},
		out: filterSliceOutput{
			one: nil,
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
