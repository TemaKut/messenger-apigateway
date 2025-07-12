package factory

import "github.com/google/wire"

var HttpSet = wire.NewSet(
	ProvideHttpProvider,
)

type HttpProvider struct{}

func ProvideHttpProvider() HttpProvider {
	return HttpProvider{}
}
