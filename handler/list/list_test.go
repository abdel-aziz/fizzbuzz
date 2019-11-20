package list

import (
	"fizzbuzz/app"
	"fizzbuzz/service/fizzbuzzer"

	"github.com/bitly/go-simplejson"
	. "gopkg.in/check.v1"
)

func (this *ListHandlerTestSuite) TestListHandler200(c *C) {
	mockFizzbuzzer := &fizzbuzzer.MockFizzbuzzer{}
	mockFizzbuzzer.On("Compute", 2, 3, "fizz", "buzz", 6).Return([]interface{}{1, "fizz", "buzz", "fizz", 5, "fizzbuzz"})
	this.App.SetContext(app.AppContext{
		Fizzbuzzer: mockFizzbuzzer,
	})

	recorder, req := this.newTestRequest(c, "GET", "/list?int1=2&int2=3&string1=fizz&string2=buzz&limit=6", nil)
	this.App.ServeHTTP(recorder, req)

	c.Assert(recorder.Code, Equals, 200)
	response, err := simplejson.NewJson(recorder.Body.Bytes())
	c.Assert(err, IsNil)

	c.Assert(response.MustArray(), HasLen, 6)
	c.Assert(response.GetIndex(0).MustInt(), Equals, 1)
	c.Assert(response.GetIndex(1).MustString(), Equals, "fizz")
	c.Assert(response.GetIndex(2).MustString(), Equals, "buzz")
	c.Assert(response.GetIndex(3).MustString(), Equals, "fizz")
	c.Assert(response.GetIndex(4).MustInt(), Equals, 5)
	c.Assert(response.GetIndex(5).MustString(), Equals, "fizzbuzz")

}

func (this *ListHandlerTestSuite) TestListHandler400(c *C) {
	recorder, req := this.newTestRequest(c, "GET", "/list", nil)
	this.App.ServeHTTP(recorder, req)
	c.Assert(recorder.Code, Equals, 400)
}
