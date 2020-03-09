package route

import (
	"pandog/infra/local/server"
)

type Base interface {
	Route() []server.Route
}
