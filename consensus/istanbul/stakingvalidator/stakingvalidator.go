package stakingvalidator

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
)

func New(addr common.Address) istanbul.Validator {
	return &defaultValidator{
		address: addr,
	}
}

func NewSet(addrs []common.Address, policy *istanbul.ProposerPolicy) istanbul.ValidatorSet {
	return newDefaultSet(addrs, policy)
}

func ExtractValidators(extraData []byte) []common.Address {
	var emptyAddrs []common.Address
	return emptyAddrs
}

// Check whether the extraData is presented in prescribed form
func ValidExtraData(extraData []byte) bool {
	return true
}

func SortedAddresses(validators []istanbul.Validator) []common.Address {
	addrs := make([]common.Address, len(validators))
	for i, validator := range validators {
		addrs[i] = validator.Address()
	}

	for i := 0; i < len(addrs); i++ {
		for j := i + 1; j < len(addrs); j++ {
			if bytes.Compare(addrs[i][:], addrs[j][:]) > 0 {
				addrs[i], addrs[j] = addrs[j], addrs[i]
			}
		}
	}

	return addrs
}
