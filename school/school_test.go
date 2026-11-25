package school

import "testing"

func TestSchool(t *testing.T) {
	if SchoolName != "สุดยอด" {
		t.Errorf("SchoolName = %q", SchoolName)
	}
	if got := GetSchoolAddress(); got != "เมืองคอน" {
		t.Errorf("GetSchoolAddress() = %q", got)
	}
}
