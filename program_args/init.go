package program_args

import (
	"flag"
	"sync"
)

type Args struct {
	DisableUI bool
}

var argsInstance *Args
var once sync.Once

func GetArgs() *Args {
	once.Do(func() {
		noUI := flag.Bool("no-ui", false, "Disable the user interface")
		flag.Parse()
		argsInstance = &Args{
			DisableUI: *noUI,
		}
	})
	return argsInstance
}
