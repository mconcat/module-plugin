package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Module interface {
	Handler() sdk.Handler
}
