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

type MigrationsDeployer struct{}

func (d *MigrationsDeployer) Deploy(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
	address, transaction, contract, err := bindings.DeployMigrations(auth, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	val, ok := backend.(*backends.SimulatedBackend)
	if ok {
		val.Commit()
	}

	session := &bindings.MigrationsSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *MigrationsDeployer) Bind(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend, address common.Address) (interface{}, error) {
	contract, err := bindings.NewMigrations(address, backend)
	if err != nil {
		return nil, err
	}

	session := &bindings.MigrationsSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("Migrations", &MigrationsDeployer{})

	migration.AddMigration(&migration.Migration{
		Number: 1,
		F: func(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend) error {
			if err := contract.Deploy(ctx, "Migrations", auth, backend); err != nil {
				return err
			}

			return nil
		},
	})
}
