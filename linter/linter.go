package linter

// Example rules set installation

import (
	dcRules "github.com/delivery-club/delivery-club-rules"
	"github.com/quasilyte/go-ruleguard/dsl"

	"github.com/peakle/dc-rules-example/linter/local"
)

func init() {
	dsl.ImportRules("", local.Bundle)   // import local rules defined only for this project
	dsl.ImportRules("", dcRules.Bundle) // import dc-rules
}
