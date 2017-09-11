// Invokes the perigord driver application

package main

import (
	_ "github.com/swarmdotmarket/token/migrations"
	"github.com/swarmdotmarket/perigord/stub"
)

func main() {
	stub.StubMain()
}
