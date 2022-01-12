package linter

import (
	dcRules "github.com/delivery-club/delivery-club-rules"
	"github.com/quasilyte/go-ruleguard/dsl"

	"github.com/peakle/dc-rules-example/linter/rules"
)

func init() {
	dsl.ImportRules("", rules.Bundle)
	dsl.ImportRules("", dcRules.Bundle)
}
