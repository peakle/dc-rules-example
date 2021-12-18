package linter

import (
	rules "github.com/delivery-club/delivery-club-rules"
	"github.com/quasilyte/go-ruleguard/dsl"
)

func init() {
	dsl.ImportRules("", rules.Bundle)
}
