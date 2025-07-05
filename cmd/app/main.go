package main

import (
	"fmt"

	"github.com/elghazx/my-dompetku/common/config"
)

func main() {
	cnf := config.Get()
	fmt.Print(cnf.Server.Host + ":" + cnf.Server.Port)
}
