package fizzbuzzer

import (
	. "gopkg.in/check.v1"
)

func (this *FizzbuzzerTestSuite) TestBasicFizzbuzzer(c *C) {
	fizzbuzzer := NewBasicFizzbuzzer()
	output := fizzbuzzer.Compute(2, 3, "fizz", "buzz", 6)

	c.Assert(output, HasLen, 6)
	c.Assert(output[0], Equals, 1)
	c.Assert(output[1], Equals, "fizz")
	c.Assert(output[2], Equals, "buzz")
	c.Assert(output[3], Equals, "fizz")
	c.Assert(output[4], Equals, 5)
	c.Assert(output[5], Equals, "fizzbuzz")
}
