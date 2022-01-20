package local

// Example local defined rules

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

func timeUtc(m dsl.Matcher) {
	m.Match(`time.Now().$method`, `time.Unix($*_).$method`, `time.UnixMicro($*_).$method`, `time.UnixMilli($*_).$method`).
		Where(!m["method"].Text.Matches(`UTC`)).
		Report("maybe UTC() call was forgotten")

	m.Match(`$x := time.Now();`, `var $x = time.Now();`, `var $x _ = time.Now();`, `$x = time.Now();`,
		`$x := time.Unix($*_);`, `var $x = time.Unix($*_);`, `var $x _ = time.Unix($*_);`, `$x = time.Unix($*_);`,
		`$x := time.UnixMicro($*_);`, `var $x = time.UnixMicro($*_);`, `var $x _ = time.UnixMicro($*_);`, `$x = time.UnixMicro($*_);`,
		`$x := time.UnixMilli($*_);`, `var $x = time.UnixMilli($*_);`, `var $x _ = time.UnixMilli($*_);`, `$x = time.UnixMilli($*_);`).
		Report("maybe UTC() call was forgotten").
		At(m["x"])

	m.Match(`time.Date($_, $_, $_, $_, $_, $_, $_, $timezone)`).
		Where(!m["timezone"].Text.Matches(`time\.UTC`)).
		Report("maybe UTC() call was forgotten").
		At(m["timezone"])

	m.Match(`$x, $_ := time.Parse($*args); $*body`,
		`$x, $_ = time.Parse($*args); $*body`,
		`var $x, $_ = time.Parse($*args); $*body`,
	).
		Where(!m["body"].Text.Matches(`\.UTC\(\)`)).
		Report("maybe UTC() call was forgotten").
		At(m["x"])
}
