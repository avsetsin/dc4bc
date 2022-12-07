package entity

import fssz "github.com/prysmaticlabs/fastssz"

type BLSToExecutionChange struct {
	ValidatorIndex     uint64   `json:"validator_index"`
	FromBlsPubkey      [48]byte `json:"from_bls_pubkey" ssz-size:"48"`
	ToExecutionAddress [20]byte `json:"to_execution_address" ssz-size:"20"`
}

var _ fssz.Marshaler = (*BLSToExecutionChange)(nil)
var _ fssz.Unmarshaler = (*BLSToExecutionChange)(nil)
