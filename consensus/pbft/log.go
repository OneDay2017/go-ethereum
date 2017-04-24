// Copyright 2017 AMIS Technologies
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package pbft

import (
	"math/big"
	"reflect"
)

func NewLog(preprepare *Preprepare) *Log {
	return &Log{
		ViewNumber:  preprepare.View.ViewNumber,
		Sequence:    preprepare.View.Sequence,
		Preprepare:  preprepare,
		Prepares:    NewMessageSet(preprepare.View, reflect.TypeOf(&Subject{})),
		Commits:     NewMessageSet(preprepare.View, reflect.TypeOf(&Subject{})),
		Checkpoints: NewMessageSet(preprepare.View, reflect.TypeOf(&Checkpoint{})),
	}
}

type Log struct {
	ViewNumber  *big.Int
	Sequence    *big.Int
	Preprepare  *Preprepare
	Prepares    MessageSet
	Commits     MessageSet
	Checkpoints MessageSet
}
