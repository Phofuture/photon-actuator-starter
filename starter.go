package actuatorStarter

import (
	"github.com/Phofuture/photon-actuator-starter/actuator"
	"github.com/Phofuture/photon-core-starter/core"
)

func init() {
	core.RegisterCoreDependency(actuator.Start)
}
