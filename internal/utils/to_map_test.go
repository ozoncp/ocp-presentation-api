package utils

import (
	"testing"

	"github.com/ozoncp/ocp-presentation-api/internal/models"
	"github.com/stretchr/testify/assert"
)

type presentationsToMapInput struct {
	one []models.Presentation
}

type presentationsToMapOutput struct {
	one map[uint64]models.Presentation
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
