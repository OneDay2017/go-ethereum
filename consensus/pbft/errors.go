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

import "errors"

var (
	ErrFromSelf          = errors.New("message comes from myself")
	ErrNotFromProposer   = errors.New("message does not come from proposer")
	ErrNilProposal       = errors.New("nil proposal")
	ErrIgnored           = errors.New("message is ignored")
	ErrSubjectNotMatched = errors.New("subjects are not matched")
	ErrInvalidSignature  = errors.New("Invalid signature")
	ErrInvalidPeerID     = errors.New("Invalid peer ID")

	// ErrNoMatchingValidator is returned when validating a peer message by verifying its signature
	// and the associated public is not found in the validator set.
	ErrNoMatchingValidator = errors.New("Cannot find matching validator of the given signature")
)
