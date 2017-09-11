package migrations

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/swarmdotmarket/perigord/contract"
	"github.com/swarmdotmarket/perigord/migration"

	"github.com/swarmdotmarket/token/bindings"
)

type MyTokenDeployer struct{}

func (d *MyTokenDeployer) Deploy(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
	address, transaction, contract, err := bindings.DeployMyToken(auth, backend, big.NewInt(1337), "FOO", 0, "BAR")
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	val, ok := backend.(*backends.SimulatedBackend)
	if ok {
		val.Commit()
	}

	session := &bindings.MyTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *MyTokenDeployer) Bind(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend, address common.Address) (interface{}, error) {
	contract, err := bindings.NewMyToken(address, backend)
	if err != nil {
		return nil, err
	}

	session := &bindings.MyTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("MyToken", &MyTokenDeployer{})

	migration.AddMigration(&migration.Migration{
		Number: 3,
		F: func(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend) error {
			if err := contract.Deploy(ctx, "MyToken", auth, backend); err != nil {
				return err
			}

			return nil
		},
	})
}
