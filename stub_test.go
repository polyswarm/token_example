package main

import (
	"testing"

	_ "github.com/swarmdotmarket/token/tests"
	_ "github.com/swarmdotmarket/token/migrations"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
