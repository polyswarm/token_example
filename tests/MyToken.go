package tests

import (
	. "gopkg.in/check.v1"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/swarmdotmarket/perigord/contract"
	"github.com/swarmdotmarket/perigord/testing"

	"github.com/swarmdotmarket/token_example/bindings"
)

type MyTokenSuite struct {
	auth    *bind.TransactOpts
	backend bind.ContractBackend
}

var _ = Suite(&MyTokenSuite{})

func (s *MyTokenSuite) SetUpTest(c *C) {
	auth, backend := testing.SetUpTest()

	s.auth = auth
	s.backend = backend
}

func (s *MyTokenSuite) TearDownTest(c *C) {
	testing.TearDownTest()
}

func (s *MyTokenSuite) TestName(c *C) {
	session := contract.Session("MyToken")
	c.Assert(session, NotNil)

	token_session, ok := session.(*bindings.MyTokenSession)
	c.Assert(ok, Equals, true)
	c.Assert(token_session, NotNil)

	ret, _ := token_session.Name()
	c.Assert(ret, Equals, "FOO")
}
