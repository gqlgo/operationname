package operationname

import (
	"fmt"
	"github.com/gqlgo/gqlanalysis"
)

func Analyzer() *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "operationname",
		Doc:  "operationname print operation name in your GraphQL query files",
		Run:  run(),
	}
}

func run() func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		for _, q := range pass.Queries {
			for _, op := range q.Operations {
				fmt.Printf("%s\n", op.Name)
			}
		}

		return nil, nil
	}
}
