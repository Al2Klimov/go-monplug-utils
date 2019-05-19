package go_monplug_utils

import (
	. "github.com/Al2Klimov/go-test-utils"
	"testing"
)

func TestOptionalThreshold_Set(t *testing.T) {
	{
		var actual OptionalThreshold
		actualErr := actual.Set("")

		AssertCallResult(
			t,
			"var t OptionalThreshold; e := t.Set(\"\"); (t, e)",
			[]interface{}{},
			[]interface{}{OptionalThreshold{}, invalidThreshold("")},
			[]interface{}{actual, actualErr},
			true,
		)
	}

	assertOptionalThreshold_Set(t, "1.5", OptionalThreshold{true, false, 0, 1.5})
	assertOptionalThreshold_Set(t, ":1.5", OptionalThreshold{true, false, 0, 1.5})
	assertOptionalThreshold_Set(t, "-1.5:", OptionalThreshold{true, false, -1.5, posInf})
	assertOptionalThreshold_Set(t, "~:", OptionalThreshold{true, false, negInf, posInf})
	assertOptionalThreshold_Set(t, "-1.5:1.5", OptionalThreshold{true, false, -1.5, 1.5})
	assertOptionalThreshold_Set(t, "@-1.5:1.5", OptionalThreshold{true, true, -1.5, 1.5})
}

func assertOptionalThreshold_Set(t *testing.T, in string, out OptionalThreshold) {
	t.Helper()

	var actual OptionalThreshold
	actualErr := actual.Set(in)

	AssertCallResult(
		t,
		"var t OptionalThreshold; e := t.Set(%#v); (t, e)",
		[]interface{}{in},
		[]interface{}{out, nil},
		[]interface{}{actual, actualErr},
		true,
	)
}

func TestOptionalThreshold_String(t *testing.T) {
	assertOptionalThreshold_String(t, OptionalThreshold{}, "")
	assertOptionalThreshold_String(t, OptionalThreshold{IsSet: true}, "0")
	assertOptionalThreshold_String(t, OptionalThreshold{true, false, 0, 1.5}, "1.5")
	assertOptionalThreshold_String(t, OptionalThreshold{true, false, -1.5, posInf}, "-1.5:")
	assertOptionalThreshold_String(t, OptionalThreshold{true, false, negInf, posInf}, "~:")
	assertOptionalThreshold_String(t, OptionalThreshold{true, false, -1.5, 1.5}, "-1.5:1.5")
	assertOptionalThreshold_String(t, OptionalThreshold{true, true, -1.5, 1.5}, "@-1.5:1.5")
}

func assertOptionalThreshold_String(t *testing.T, in OptionalThreshold, out string) {
	t.Helper()

	AssertCallResult(
		t,
		"(&%#v).String()",
		[]interface{}{in},
		[]interface{}{out},
		[]interface{}{in.String()},
		true,
	)
}
