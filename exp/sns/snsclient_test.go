package sns_test
import (
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/exp/sns"
	. "launchpad.net/gocheck"
	"net/http"
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
	auth   aws.Auth
	region aws.Region
}

func (s *ClientSuite) SetUpSuite(c *C) {
	s.HTTPSuite.SetUpSuite(c)
	s.auth = aws.Auth{"abc", "123"}
	s.region = aws.Region{SNSEndpoint: testServer.URL}
}

// Test if SNSWithClient returns an SNS
func (s *ClientSuite) TestWithClient(c *C) {
	instance := sns.SNSWithClient(s.auth, s.region, &http.Client{})
	c.Assert(instance, NotNil)
}

// Test if performing a request invokes the SNSs configured
// http.Client
func (s *ClientSuite) TestClientUsed(c *C) {
	transport := &FlaggingRoundTripper{false, http.DefaultTransport}
	instance := sns.SNSWithClient(s.auth, s.region, &http.Client{Transport: transport})
	testServer.Response(200, nil, TestListTopicsXmlOK)
	instance.ListTopics(nil)
	testServer.WaitRequest()
	c.Assert(transport.Transported, Equals, true)
}
