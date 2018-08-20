package migrations

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"

	"github.com/polyswarm/token_example/bindings"
)

type FooDeployer struct{}

func (d *FooDeployer) Deploy(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
	address, transaction, contract, err := bindings.DeployFoo(auth, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	val, ok := backend.(*backends.SimulatedBackend)
	if ok {
		val.Commit()
	}

	session := &bindings.FooSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *FooDeployer) Bind(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend, address common.Address) (interface{}, error) {
	contract, err := bindings.NewFoo(address, backend)
	if err != nil {
		return nil, err
	}

	session := &bindings.FooSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("Foo", &FooDeployer{})

	migration.AddMigration(&migration.Migration{
		Number: 2,
		F: func(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend) error {
			if err := contract.Deploy(ctx, "Foo", auth, backend); err != nil {
				return err
			}

			return nil
		},
	})
}
