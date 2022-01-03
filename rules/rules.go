package rules

import "github.com/quasilyte/go-ruleguard/dsl"

var Bundle = dsl.Bundle{}

func stringsSimplify(m dsl.Matcher) {
	// Some issues have simple fixes that can be expressed as
	// a replacement pattern. Rules can use Suggest() function
	// to add a quickfix action for such issues.
	m.Match(`strings.Replace($s, $old, $new, -1)`).
		Report(`this Replace call can be simplified`).
		Suggest(`strings.ReplaceAll($s, $old, $new)`)
}
