// Package utils implement a simple internal library for Ozon Code Platform Presentation API.
package utils

import (
	"errors"
	"testing"

	"github.com/ozoncp/ocp-presentation-api/internal/model"
	"github.com/stretchr/testify/assert"
)

type presentationsToMapInput struct {
	one []model.Presentation
}

type presentationsToMapOutput struct {
	one map[uint64]model.Presentation
	two error
}

type presentationsToMapTestCase struct {
	in  presentationsToMapInput
	out presentationsToMapOutput
}

var presentationsToMapTestCases = []presentationsToMapTestCase{
	{
		in: presentationsToMapInput{
			one: nil,
		},
		out: presentationsToMapOutput{
			one: nil,
			two: nil,
		},
	},
	{
		in: presentationsToMapInput{
			one: []model.Presentation{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
		},
		out: presentationsToMapOutput{
			one: map[uint64]model.Presentation{1: {ID: 1}, 2: {ID: 2}, 3: {ID: 3}, 4: {ID: 4}, 5: {ID: 5}},
			two: nil,
		},
	},
	{
		in: presentationsToMapInput{
			one: []model.Presentation{{ID: 1}, {ID: 1}},
		},
		out: presentationsToMapOutput{
			one: nil,
			two: errors.New("identifiers are not unique"),
		},
	},
}

func TestPresentationsToMap(t *testing.T) {
	ast := assert.New(t)

	for _, testCase := range presentationsToMapTestCases {
		out, in := testCase.out, testCase.in
		result, err := PresentationsToMap(in.one)
		ast.Equal(out.one, result, "Test Case: %v %v", in, out)
		ast.Equal(out.two, err, "Test Case: %v %v", in, out)
	}
}
