package main

import (
	"flag"
	"fmt"
)

func main() {
	cfg := flag.String("c", "cfg.dev.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	commonCfg := flag.String("cc", "cfg.common.json", "common configuration file")
	flag.Parse()

	fmt.Println(*cfg)
	fmt.Println(*version)
	fmt.Println(*commonCfg)
}
