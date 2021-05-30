// Package utils implement a simple internal library for Ozon Code Platform Presentation API.
package utils

import (
	"testing"

	"github.com/ozoncp/ocp-presentation-api/internal/models"
	"github.com/stretchr/testify/assert"
)

type splitPresentationsToBulksInput struct {
	one []models.Presentation
	two uint
}

type splitPresentationsToBulksOutput struct {
	one [][]models.Presentation
}

type splitPresentationsToBulksTestCase struct {
	in  splitPresentationsToBulksInput
	out splitPresentationsToBulksOutput
}

var splitPresentationsToBulksTestCases = []splitPresentationsToBulksTestCase{
	{
		in: splitPresentationsToBulksInput{
			one: []models.Presentation{},
			two: 0,
		},
		out: splitPresentationsToBulksOutput{
			one: nil,
		},
	},
	{
		in: splitPresentationsToBulksInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: 1,
		},
		out: splitPresentationsToBulksOutput{
			one: [][]models.Presentation{
				{{ID: 1}},
				{{ID: 2}},
				{{ID: 3}},
				{{ID: 4}},
				{{ID: 5}},
			},
		},
	},
	{
		in: splitPresentationsToBulksInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: 2,
		},
		out: splitPresentationsToBulksOutput{
			one: [][]models.Presentation{
				{{ID: 1}, {ID: 2}},
				{{ID: 3}, {ID: 4}},
				{{ID: 5}},
			},
		},
	},
	{
		in: splitPresentationsToBulksInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: 3,
		},
		out: splitPresentationsToBulksOutput{
			one: [][]models.Presentation{
				{{ID: 1}, {ID: 2}, {ID: 3}},
				{{ID: 4}, {ID: 5}},
			},
		},
	},
	{
		in: splitPresentationsToBulksInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: 4,
		},
		out: splitPresentationsToBulksOutput{
			one: [][]models.Presentation{
				{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}},
				{{ID: 5}},
			},
		},
	},
	{
		in: splitPresentationsToBulksInput{
			one: []models.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			two: 5,
		},
		out: splitPresentationsToBulksOutput{
			one: [][]models.Presentation{
				{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
			},
		},
	},
}

func TestSplitPresentationsToBulks(t *testing.T) {
	ast := assert.New(t)

	for _, testCase := range splitPresentationsToBulksTestCases {
		out, in := testCase.out, testCase.in
		ast.Equal(out.one, SplitPresentationsToBulks(in.one, in.two), "Test Case: %v %v", in, out)
	}
}
