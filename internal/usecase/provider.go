package usecase

import (
	"github.com/google/wire"
	"github.com/s-yakubovskiy/whereami/internal/usecase/locator"
)

var ProviderSet = wire.NewSet(locator.NewLocatorUserCase)
