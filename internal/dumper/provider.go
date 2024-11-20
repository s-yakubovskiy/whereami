package dumper

import (
	"github.com/google/wire"
	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

func ProvideDumper(d contracts.LKInterface) (*DumperJSON, error) {
	return NewDumperJSON(d)
}

var ProviderSet = wire.NewSet(
	ProvideDumper,
)
