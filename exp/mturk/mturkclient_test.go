package mturk_test

import (
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/exp/mturk"
	. "launchpad.net/gocheck"
	"net/http"
	"net/url"
)

// http.RoundTripper which sets a flag before invoking a delegate
// http.RoundTripper
type FlaggingRoundTripper struct {
	Transported bool
	delegate    http.RoundTripper
}

// Create a FlaggingRoundTripper with the specified delgate
func (r *FlaggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	r.Transported = true
	return r.delegate.RoundTrip(req)
}

var _ = Suite(&ClientSuite{})

type ClientSuite struct {
	HTTPSuite
	auth aws.Auth
}

func (s *ClientSuite) SetUpSuite(c *C) {
	s.HTTPSuite.SetUpSuite(c)
	s.auth = aws.Auth{"abc", "123"}
}

// Test if MTurkWithClient returns an MTurk
func (s *ClientSuite) TestWithClient(c *C) {
	instance := mturk.MTurkWithClient(s.auth, &http.Client{})
	c.Assert(instance, NotNil)
}

// Test if performing a request invokes the MTurkss configured
// http.Client
func (s *ClientSuite) TestClientUsed(c *C) {
	transport := &FlaggingRoundTripper{false, http.DefaultTransport}
	instance := mturk.MTurkWithClient(s.auth, &http.Client{Transport: transport})
	u, err := url.Parse(testServer.URL)
	if err != nil {
		panic(err.Error())
	}
	instance.URL = u
	testServer.PrepareResponse(200, nil, BasicHitResponse)
	question := mturk.ExternalQuestion{
		ExternalURL: "http://www.amazon.com",
		FrameHeight: 200,
	}
	reward := mturk.Price{
		Amount:       "0.01",
		CurrencyCode: "USD",
	}
	hit, err := instance.CreateHIT("title", "description", question, reward, 1, 2, "key1,key2", 3, nil, "annotation")
	testServer.WaitRequest()
	c.Assert(err, IsNil)
	c.Assert(hit, NotNil)
	c.Assert(transport.Transported, Equals, true)
}
