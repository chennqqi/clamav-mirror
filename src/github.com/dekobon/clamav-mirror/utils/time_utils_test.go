package utils

import (
	"testing"
	"time"
)

func TestValidParseClamAVTimeStamp(t *testing.T) {
	expected, err := time.Parse(time.RFC3339, "2006-01-02T15:04:00-07:00")

	if err != nil {
		t.Error(err)
	}

	parsed, err := ParseClamAVTimeStamp("02 Jan 2006 15:04 -0700")

	if err != nil {
		t.Error(err)
	}

	if !expected.Equal(parsed) {
		t.Errorf("Timestamps are not equal.\n"+
			"Expected: %v\n"+
			"Actual  : %v",
			expected, parsed)
	}
}

func TestInvalidParseClamAVTimeStamp(t *testing.T) {
	_, err := ParseClamAVTimeStamp("02 Jan 2006 15:04Z")

	if err == nil {
		t.Error("Expected time parsing exception not thrown")
	}
}
