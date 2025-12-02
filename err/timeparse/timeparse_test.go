package timeparse

import "testing"

func TestParsetime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:44", false},
		{"aa:bb:cc", false},
		{"5:23:", false},
		{"", false},
		{"11:15", false},
	}
	for _, data := range table {
		_, err := PareTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v : %v,err shoud be nil", data.time, err)
		}
	}
}
