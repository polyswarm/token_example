package main

import (
	"testing"

	_ "github.com/swarmdotmarket/token_example/tests"
	_ "github.com/swarmdotmarket/token_example/migrations"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
