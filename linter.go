//go:build ignore
// +build ignore

package main

import (
	dcRules "github.com/delivery-club/delivery-club-rules"
	"github.com/peakle/dc-rules-example/rules"
	"github.com/quasilyte/go-ruleguard/dsl"
)

func init() {
	dsl.ImportRules("", rules.Bundle)
	dsl.ImportRules("", dcRules.Bundle)
}
