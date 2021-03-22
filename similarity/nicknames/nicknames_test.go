package nicknames

import (
	"fmt"
	"testing"
)

func TestNicknames(test *testing.T) {
	var tests = []struct {
		name     string
		nickname string
		expected bool
	}{{"", "", false}, // no name supplied
		{"David", "David", true},  // full name is its own nickname
		{"Robert", "Rob", true},   // name and nickname correct
		{"Peter", "Paul", false},  // name list but not a known nickname
		{"Edward", "John", false}, // name not listed as having any nicknames
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("Nicknames:%s with %s ... \n", tt.name, tt.nickname)

		test.Run(testname, func(t *testing.T) {
			found := FindNicknames(tt.name, tt.nickname)

			if found != tt.expected{
				t.Errorf("got %v, want %v", found, tt.expected)
			}
		})

	}

}
