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
			t.Fatalf("v: %v\n\nexpected: %v\nactual: %v",
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
