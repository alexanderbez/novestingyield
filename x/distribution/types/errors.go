package types

import (
	"cosmossdk.io/errors"
	dstrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// x/distribution (extended) module sentinel errors
var (
	ErrForbiddenWithdrawal = errors.Register(dstrtypes.ModuleName, 99, "forbidden withdrawal")
)
