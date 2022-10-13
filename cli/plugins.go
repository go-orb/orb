package cli

import (
	"jochum.dev/orb/orb/util/container"
)

var Plugins = container.NewPlugins(
	func() Cli { return nil }, // Plugin factory
	func() any { return nil }, // Config factory
)
