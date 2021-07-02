package api_test

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/api"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

func TestPanic(t *testing.T) {
	defer func() {
		err := recover().(error)
		if !errors.Is(err, api.ErrInvalidArgument) {
			t.Fatalf("Wrong panic message: %s", err.Error())
		}
	}()

	api.NewPresentationAPI(nil, 0)
}
