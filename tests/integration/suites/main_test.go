package suites

import (
	"github.com/TemaKut/messenger-apigateway/tests/integration/suites/userregister"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSuites(t *testing.T) {
	suite.Run(t, new(userregister.Suite))
}
