package timeUtc

import (
	"time"
)

func warning1() {
	n := time.Now() // want `\QtimeUtc: maybe UTC() call was forgotten`

	print(n)
}

func warning2() {
	t, _ := time.Parse("", "") // false negative // it's because we call UTC() on line 20 // TODO need fix
	print(t)
	print(time.Now().Round(time.Minute))              // want `\QtimeUtc: maybe UTC() call was forgotten`
	print(time.Now().Add(time.Second))                // want `\QtimeUtc: maybe UTC() call was forgotten`
	print(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)) // want `\QtimeUtc: maybe UTC() call was forgotten`

	print(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local).Add(time.Second).UTC()) // want `\QtimeUtc: maybe UTC() call was forgotten`
	_ = time.Date(0, 0, 0, 0, 0, 0, 0, time.Local).UTC().Add(time.Second)    // want `\QtimeUtc: maybe UTC() call was forgotten`
}

func warning3() time.Time {
	t, _ := time.Parse("", "") // want `\QtimeUtc: maybe UTC() call was forgotten`
	print(t)

	return t
}
func good() {
	k, _ := time.Parse("", "")
	print(k.UTC())

	_ = time.Now().UTC()
	_ = time.Now().UTC().Add(time.Second)

	print(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC).Add(time.Second))
	_ = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC).Add(time.Second)
	time.ParseInLocation("", "", nil)
}
