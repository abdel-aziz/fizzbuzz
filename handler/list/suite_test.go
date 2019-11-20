package list

import (
	"fizzbuzz/app"

	. "gopkg.in/check.v1"

	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuite(t *testing.T) { TestingT(t) }

type ListHandlerTestSuite struct {
	App app.App
}

func (this *ListHandlerTestSuite) SetUpSuite(c *C) {
	var err error
	this.App, err = app.NewApp(app.Config{
		Router: "tester",
	})
	c.Assert(err, IsNil)

	Register(this.App)
}

var _ = Suite(&ListHandlerTestSuite{})

func (this *ListHandlerTestSuite) newTestRequest(c *C, method, path string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	request, err := http.NewRequest(method, path, nil)
	c.Assert(err, IsNil)
	return httptest.NewRecorder(), request
}
