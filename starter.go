package actuatorStarter

import (
	"github.com/dennesshen/photon-actuator-starter/actuator"
	"github.com/dennesshen/photon-core-starter/core"
)

func init() {
	core.RegisterCoreDependency(actuator.Start)
}
