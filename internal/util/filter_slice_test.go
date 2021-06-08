// Package util implement a simple internal library for Ozon Code Platform Presentation API.
package util

import (
	"testing"

	"github.com/ozoncp/ocp-presentation-api/internal/model"
	"github.com/stretchr/testify/assert"
)

type filterSliceInput struct {
	one []model.Presentation
	two []model.Presentation
}

type filterSliceOutput struct {
	one []model.Presentation
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
			one: []model.Presentation{},
			two: nil,
		},
		out: filterSliceOutput{
			one: []model.Presentation{},
		},
	},
	{
		in: filterSliceInput{
			one: []model.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []model.Presentation{{ID: 1}},
		},
		out: filterSliceOutput{
			one: []model.Presentation{{ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
		},
	},
	{
		in: filterSliceInput{
			one: []model.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []model.Presentation{{ID: 1}, {ID: 2}},
		},
		out: filterSliceOutput{
			one: []model.Presentation{{ID: 3}, {ID: 4}, {ID: 5}},
		},
	},
	{
		in: filterSliceInput{
			one: []model.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []model.Presentation{{ID: 1}, {ID: 2}, {ID: 3}},
		},
		out: filterSliceOutput{
			one: []model.Presentation{{ID: 4}, {ID: 5}},
		},
	},
	{
		in: filterSliceInput{
			one: []model.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []model.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}},
		},
		out: filterSliceOutput{
			one: []model.Presentation{{ID: 5}},
		},
	},
	{
		in: filterSliceInput{
			one: []model.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: []model.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
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
