package main

import (
	"github.com/wearelumenai/clusauth/internal"
	"github.com/wearelumenai/clusauth/internal/conf"

	"log"
)

func main() {
	var conf, err = conf.Load("configs")

	if err == nil {
		err = internal.ApplySentry(conf)
	}

	if err == nil {
		err = internal.Serve(conf)
	}

	if err != nil {
		log.Fatal(err)
	}
}
