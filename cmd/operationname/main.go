package main

import (
	"github.com/gqlgo/gqlanalysis/multichecker"
	"github.com/gqlgo/operationname"
)

func main() {
	multichecker.Main(
		operationname.Analyzer(),
	)
}
