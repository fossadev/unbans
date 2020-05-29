package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type environmentSuite struct {
	suite.Suite
}

func (t *environmentSuite) TestSetsCorrectEnvironmentWhenValidPresent() {
	t.Equal("development", getEnvironment("development"))
	t.Equal("production", getEnvironment("production"))
}

func (t *environmentSuite) TestSetsValidEnvironmentWhenInvalidPresent() {
	t.Equal("production", getEnvironment("fake_environment_should_be_either_development_or_production"))
}

func TestEnvironment(t *testing.T) {
	suite.Run(t, &environmentSuite{})
}
