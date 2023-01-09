package semver

import "testing"

// test the VersionSatisfiesRange function
func TestVersionSatisfiesRange(t *testing.T) {
	// >=
	if !VersionSatisfiesRange("1.2.3", ">=1.2.0") {
		t.Error("VersionSatisfiesRange failed")
	}
	if VersionSatisfiesRange("1.2.3", ">=1.2.4") {
		t.Error("VersionSatisfiesRange failed")
	}
	// <=
	if !VersionSatisfiesRange("1.2.3", "<=1.2.4") {
		t.Error("VersionSatisfiesRange failed")
	}
	if VersionSatisfiesRange("1.2.3", "<=1.2.0") {
		t.Error("VersionSatisfiesRange failed")
	}
	// >
	if !VersionSatisfiesRange("1.2.3", ">1.2.0") {
		t.Error("VersionSatisfiesRange failed")
	}
	if VersionSatisfiesRange("1.2.3", ">1.2.4") {
		t.Error("VersionSatisfiesRange failed")
	}
	// <
	if !VersionSatisfiesRange("1.2.3", "<1.2.4") {
		t.Error("VersionSatisfiesRange failed")
	}
	if VersionSatisfiesRange("1.2.3", "<1.2.0") {
		t.Error("VersionSatisfiesRange failed")
	}
	// =
	if !VersionSatisfiesRange("1.2.3", "=1.2.3") {
		t.Error("VersionSatisfiesRange failed")
	}
	if VersionSatisfiesRange("1.2.3", "=1.2.4") {
		t.Error("VersionSatisfiesRange failed")
	}
	// ~
	if !VersionSatisfiesRange("1.2.3", "~1.2.0") {
		t.Error("VersionSatisfiesRange failed")
	}
	if VersionSatisfiesRange("1.2.3", "~1.3.0") {
		t.Error("VersionSatisfiesRange failed")
	}
	// ^
	if !VersionSatisfiesRange("1.2.3", "^1.2.0") {
		t.Error("VersionSatisfiesRange failed")
	}
	if VersionSatisfiesRange("1.2.3", "^1.3.0") {
		t.Error("VersionSatisfiesRange failed")
	}
}
