package usecase

import (
	"github.com/google/wire"
	"github.com/s-yakubovskiy/whereami/internal/usecase/keeper"
	"github.com/s-yakubovskiy/whereami/internal/usecase/locator"
	"github.com/s-yakubovskiy/whereami/internal/usecase/zosher"
)

var ProviderSet = wire.NewSet(locator.NewLocatorUserCase, keeper.NewLocationKeeperUseCase, zosher.NewZoshUseCase)
