// Copyright 2015 The elementalcore Authors
// This file is part of the elementalcore library.
//
// The elementalcore library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The elementalcore library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the elementalcore library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"math/big"

	"github.com/Elemental-core/elementalcore/core/state"
	"github.com/Elemental-core/elementalcore/core/types"
	"github.com/Elemental-core/elementalcore/core/vm"
)

// Validator is an interface which defines the standard for block validation. It
// is only responsible for validating block contents, as the header validation is
// done by the specific consensus engines.
//
type Validator interface {
	// ValidateBody validates the given block's content.
	ValidateBody(block *types.Block) error

	// ValidateState validates the given statedb and optionally the receipts and
	// gas used.
	ValidateState(block, parent *types.Block, state *state.StateDB, receipts types.Receipts, usedGas *big.Int) error
	// ValidateDposState validates the given dpos state
	ValidateDposState(block *types.Block) error
}

// Processor is an interface for processing blocks using a given initial state.
//
// Process takes the block to be processed and the statedb upon which the
// initial state is based. It should return the receipts generated, amount
// of gas used in the process and return an error if any of the internal rules
// failed.
type Processor interface {
	Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) (types.Receipts, []*types.Log, *big.Int, error)
}
