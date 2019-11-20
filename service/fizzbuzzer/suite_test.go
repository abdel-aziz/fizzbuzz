package fizzbuzzer

import (
	. "gopkg.in/check.v1"

	"testing"
)

func TestSuite(t *testing.T) { TestingT(t) }

type FizzbuzzerTestSuite struct{}

var _ = Suite(&FizzbuzzerTestSuite{})
