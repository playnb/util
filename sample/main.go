package main

import (
	_ "github.com/playnb/util/config"
	"github.com/playnb/util/log"
)

func main() {
	log.InitPanic("../tmp")
	log.Init(log.DefaultLogger("../tmp", "run"))
	defer log.Flush()

	log.Trace("--")

}
