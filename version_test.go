package version

import (
	"reflect"
	"testing"
)

func TestNewVersion(t *testing.T) {
	cases := []struct {
		version  string
		expected Version
	}{
		// Test 0
		{"0", Version{0, "0", ""}},
		{"0:0", Version{0, "0", ""}},
		{"0:0-", Version{0, "0", ""}},
		{"0:0-0", Version{0, "0", "0"}},
		{"0:0.0-0.0", Version{0, "0.0", "0.0"}},
		// Test epoch
		{"1:0", Version{1, "0", ""}},
		{"5:1", Version{5, "1", ""}},
		// Test multiple hyphens
		{"0:0-0-0", Version{0, "0", "0-0"}},
		{"0:0-0-0-0", Version{0, "0", "0-0-0"}},
		// Test multiple colons
		{"0:0:0-0", Version{0, "0:0", "0"}},
		{"0:0:0:0-0", Version{0, "0:0:0", "0"}},
		// Test multiple hyphens and colons
		{"0:0:0-0-0", Version{0, "0:0", "0-0"}},
		{"0:0-0:0-0", Version{0, "0", "0:0-0"}},
		// Test version with leading and trailing spaces
		{"  	0:0-1", Version{0, "0", "1"}},
		{"0:0-1  	", Version{0, "0", "1  	"}},
		{"	  0:0-1  	", Version{0, "0", "1  	"}},
		// Test empty version
		{"", Version{}},
		{" ", Version{0, " ", ""}},
		{"0:", Version{}},
		// Test version with embedded spaces
		{"0:0 0-1", Version{0, "0 0", "1"}},
		// Test version with negative epoch
		{"-1:0-1", Version{-1, "0", "1"}},
		// Test invalid characters in epoch
		{"a:0-0", Version{0, "0", "0"}},
		{"A:1-2", Version{0, "1", "2"}},
		// Test version not starting with a digit
		{"0:abc3-0", Version{0, "abc3", "0"}},
		// Test actual version
		{"1.2.3", Version{0, "1.2.3", ""}},
		{"1:1.2.3", Version{1, "1.2.3", ""}},
		{"A:1.2.3", Version{0, "1.2.3", ""}},
		{"-1:1.2.3", Version{-1, "1.2.3", ""}},
		{"6.0-4.el6.x86_64", Version{0, "6.0", "4.el6.x86_64"}},
		{"c105b9de-4e0fd3a3", Version{0, "c105b9de", "4e0fd3a3"}},
		{"4.999.9-0.5.beta.20091007git.el6", Version{0, "4.999.9", "0.5.beta.20091007git.el6"}},
	}

	for _, tc := range cases {
		actual := NewVersion(tc.version)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf(
				"version: %s\nexpected: %v\nactual: %v",
				tc.version, tc.expected, actual)
		}
	}
}

func TestEqual(t *testing.T) {
	for _, tc := range cases {
		expected := (tc.expected == 0)

		// compare 'version'
		a := Version{0, tc.v1, ""}
		b := Version{0, tc.v2, ""}
		if actual := a.Equal(b); actual != expected {
			t.Errorf("[Version] v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, expected, actual)
		}

		// compare 'release'
		a = Version{0, "", tc.v1}
		b = Version{0, "", tc.v2}
		if actual := a.Equal(b); actual != expected {
			t.Errorf("[Release] v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, tc.expected, actual)
		}
	}
}

func TestGreaterThan(t *testing.T) {
	for _, tc := range cases {
		expected := tc.expected > 0

		// compare 'version'
		a := Version{0, tc.v1, ""}
		b := Version{0, tc.v2, ""}
		if actual := a.GreaterThan(b); actual != expected {
			t.Errorf("[Version] v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, expected, actual)
		}

		// compare 'release'
		a = Version{0, "", tc.v1}
		b = Version{0, "", tc.v2}
		if actual := a.GreaterThan(b); actual != expected {
			t.Errorf("[Release] v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, tc.expected, actual)
		}
	}
}

func TestLessThan(t *testing.T) {
	for _, tc := range cases {
		expected := tc.expected < 0

		// compare 'version'
		a := Version{0, tc.v1, ""}
		b := Version{0, tc.v2, ""}
		if actual := a.LessThan(b); actual != expected {
			t.Errorf("[Version] v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, expected, actual)
		}

		// compare 'release'
		a = Version{0, "", tc.v1}
		b = Version{0, "", tc.v2}
		if actual := a.LessThan(b); actual != expected {
			t.Errorf("[Release] v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, tc.expected, actual)
		}
	}
}

func TestCompare(t *testing.T) {
	for _, tc := range cases {
		// compare 'version'
		a := Version{0, tc.v1, ""}
		b := Version{0, tc.v2, ""}
		if actual := a.Compare(b); actual != tc.expected {
			t.Errorf("[Version] v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, tc.expected, actual)
		}

		// compare 'release'
		a = Version{0, "", tc.v1}
		b = Version{0, "", tc.v2}
		if actual := a.Compare(b); actual != tc.expected {
			t.Errorf("[Release] v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, tc.expected, actual)
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		v        Version
		expected string
	}{
		{Version{2, "7.4.052", "1.el6"}, "2:7.4.052-1.el6"},
		{Version{2, "7.4.052", "1"}, "2:7.4.052-1"},
		{Version{0, "7.4.052", "1"}, "7.4.052-1"},
	}
	for _, tc := range cases {
		actual := tc.v.String()
		if actual != tc.expected {
			t.Errorf("v: %v\n\nexpected: %v\nactual: %v",
				tc.v, tc.expected, actual)
		}
	}
}

func TestRpmVerCmp(t *testing.T) {
	// Import cases from version_testcase.go
	for _, tc := range cases {
		actual := rpmvercmp(tc.v1, tc.v2)
		if tc.expected != actual {
			t.Fatalf("v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, tc.expected, actual)
		}
	}
}

func TestParseAndCompare(t *testing.T) {
	cases := []struct {
		v1       string
		expected int
		v2       string
	}{
		// Oracle Linux corner cases.
		{"2.9.1-6.0.1.el7_2.3", GREATER, "2.9.1-6.el7_2.3"},
		{"3.10.0-327.28.3.el7", GREATER, "3.10.0-327.el7"},
		{"3.14.3-23.3.el6_8", GREATER, "3.14.3-23.el6_7"},
		{"2.23.2-22.el7_1", LESS, "2.23.2-22.el7_1.1"},

		// Tests imported from tests/rpmvercmp.at
		{"1.0", EQUAL, "1.0"},
		{"1.0", LESS, "2.0"},
		{"2.0", GREATER, "1.0"},
		{"2.0.1", EQUAL, "2.0.1"},
		{"2.0", LESS, "2.0.1"},
		{"2.0.1", GREATER, "2.0"},
		{"2.0.1a", EQUAL, "2.0.1a"},
		{"2.0.1a", GREATER, "2.0.1"},
		{"2.0.1", LESS, "2.0.1a"},
		{"5.5p1", EQUAL, "5.5p1"},
		{"5.5p1", LESS, "5.5p2"},
		{"5.5p2", GREATER, "5.5p1"},
		{"5.5p10", EQUAL, "5.5p10"},
		{"5.5p1", LESS, "5.5p10"},
		{"5.5p10", GREATER, "5.5p1"},
		{"10xyz", LESS, "10.1xyz"},
		{"10.1xyz", GREATER, "10xyz"},
		{"xyz10", EQUAL, "xyz10"},
		{"xyz10", LESS, "xyz10.1"},
		{"xyz10.1", GREATER, "xyz10"},
		{"xyz.4", EQUAL, "xyz.4"},
		{"xyz.4", LESS, "8"},
		{"8", GREATER, "xyz.4"},
		{"xyz.4", LESS, "2"},
		{"2", GREATER, "xyz.4"},
		{"5.5p2", LESS, "5.6p1"},
		{"5.6p1", GREATER, "5.5p2"},
		{"5.6p1", LESS, "6.5p1"},
		{"6.5p1", GREATER, "5.6p1"},
		{"6.0.rc1", GREATER, "6.0"},
		{"6.0", LESS, "6.0.rc1"},
		{"10b2", GREATER, "10a1"},
		{"10a2", LESS, "10b2"},
		{"1.0aa", EQUAL, "1.0aa"},
		{"1.0a", LESS, "1.0aa"},
		{"1.0aa", GREATER, "1.0a"},
		{"10.0001", EQUAL, "10.0001"},
		{"10.0001", EQUAL, "10.1"},
		{"10.1", EQUAL, "10.0001"},
		{"10.0001", LESS, "10.0039"},
		{"10.0039", GREATER, "10.0001"},
		{"4.999.9", LESS, "5.0"},
		{"5.0", GREATER, "4.999.9"},
		{"20101121", EQUAL, "20101121"},
		{"20101121", LESS, "20101122"},
		{"20101122", GREATER, "20101121"},
		{"2_0", EQUAL, "2_0"},
		{"2.0", EQUAL, "2_0"},
		{"2_0", EQUAL, "2.0"},
		{"a", EQUAL, "a"},
		{"a+", EQUAL, "a+"},
		{"a+", EQUAL, "a_"},
		{"a_", EQUAL, "a+"},
		{"+a", EQUAL, "+a"},
		{"+a", EQUAL, "_a"},
		{"_a", EQUAL, "+a"},
		{"+_", EQUAL, "+_"},
		{"_+", EQUAL, "+_"},
		{"_+", EQUAL, "_+"},
		{"+", EQUAL, "_"},
		{"_", EQUAL, "+"},
		{"1.0~rc1", EQUAL, "1.0~rc1"},
		{"1.0~rc1", LESS, "1.0"},
		{"1.0", GREATER, "1.0~rc1"},
		{"1.0~rc1", LESS, "1.0~rc2"},
		{"1.0~rc2", GREATER, "1.0~rc1"},
		{"1.0~rc1~git123", EQUAL, "1.0~rc1~git123"},
		{"1.0~rc1~git123", LESS, "1.0~rc1"},
		{"1.0~rc1", GREATER, "1.0~rc1~git123"},

		// Test epoch
		{"1:1.0~rc1", GREATER, "0:1.0~rc1"},
		{"1.0~rc1", LESS, "2:1.0~rc1"},
		{"3:1.0~rc1", EQUAL, "3:1.0~rc1"},
	}

	for _, tc := range cases {
		v1 := NewVersion(tc.v1)
		v2 := NewVersion(tc.v2)
		actual := v1.Compare(v2)
		if actual != tc.expected {
			t.Errorf("v1: %s\nv2: %s\nexpected: %v\nactual: %v",
				tc.v1, tc.v2, tc.expected, actual)
		}
	}
}
