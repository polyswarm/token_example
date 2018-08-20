// Example main file for a native dapp, replace with application code
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"

	"github.com/polyswarm/token_example/bindings"
	_ "github.com/polyswarm/token_example/migrations"
)

func main() {
	// Run our migrations
	migration.RunMigrations(context.Background())

	session, ok := contract.Session("MyToken").(*bindings.MyTokenSession)
	if !ok {
		fmt.Println("Did our migrations complete successfully?")
		os.Exit(1)
	}

	name, _ := session.Name()
	totalSupply, _ := session.TotalSupply()
	symbol, _ := session.Symbol()

	fmt.Printf("Let's spend some %s\n", name)
	fmt.Printf("There are %d %s in total\n", totalSupply, symbol)
}
