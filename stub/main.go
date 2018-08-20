// Invokes the perigord driver application

package main

import (
	_ "github.com/polyswarm/token_example/migrations"
	"github.com/polyswarm/perigord/stub"
)

func main() {
	stub.StubMain()
}
