//  Copyright (C) 2021-2023 Chronicle Labs, Inc.
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Affero General Public License as
//  published by the Free Software Foundation, either version 3 of the
//  License, or (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Affero General Public License for more details.
//
//  You should have received a copy of the GNU Affero General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.

package median

import (
	"context"
	"math/big"
	"time"

	"github.com/defiweb/go-eth/types"
)

// Median is an interface for the Medianizer oracle contract:
// https://github.com/makerdao/median/
//
// Contract documentation:
// https://docs.makerdao.com/smart-contract-modules/oracle-module/median-detailed-documentation
type Median interface {
	// Address returns medianizer contract address.
	Address() types.Address

	// Age returns the value from contract's age method. The age is the block
	// timestamp of last price val update.
	Age(ctx context.Context) (time.Time, error)

	// Bar returns the value from contract's bar method. The bar method returns
	// the minimum number of prices necessary to accept a new median value.
	Bar(ctx context.Context) (int64, error)

	// Val returns current asset price form the contract's storage.
	Val(ctx context.Context) (*big.Int, error)

	// Wat returns asset name.
	Wat(ctx context.Context) (string, error)

	// Feeds returns a list of all Ethereum addresses that are authorized to update
	// Oracle prices (orcls).
	Feeds(ctx context.Context) ([]types.Address, error)

	// Poke sends transaction to the smart contract which invokes contract's
	// poke method, which updates asset price (val).  If simulateBeforeRun is
	// set to true, then transaction will be simulated on the EVM before actual
	// transaction will be sent.
	Poke(ctx context.Context, prices []*Price, simulateBeforeRun bool) (*types.Hash, error)

	// Lift sends transaction to the smart contract which invokes contract's
	// lift method, which sends  adds given addresses to the feeds list (orcls).
	// If simulateBeforeRun is set to true, then transaction will be simulated
	// on the EVM before actual transaction will be sent.
	Lift(ctx context.Context, addresses []types.Address, simulateBeforeRun bool) (*types.Hash, error)

	// Drop sends transaction to the smart contract which invokes contract's
	// drop method, which removes given addresses from the feeds list (orcls).
	// If simulateBeforeRun is set to true, then transaction will be simulated
	// on the EVM before actual transaction will be sent.
	Drop(ctx context.Context, addresses []types.Address, simulateBeforeRun bool) (*types.Hash, error)

	// SetBar sends transaction to the smart contract which invokes contract's
	// setBar method, which sets bar variable (quorum).  If simulateBeforeRun is
	// set to true, then transaction will be simulated on the EVM before actual
	// transaction will be sent.
	SetBar(ctx context.Context, bar *big.Int, simulateBeforeRun bool) (*types.Hash, error)
}
