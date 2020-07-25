package cmd

import (
	"testing"
)

func TestFormatDate(t *testing.T) {
	var dates = []struct {
		in   string
		want string
	}{
		{"2006-01-02T15:04:05-07:00", "02.01.2006 um 23:04 Uhr"},
		{"2020-01-01T10:00:00+01:00", "01.01.2020 um 10:00 Uhr"},
		{"2020-05-31T10:04:06+02:00", "31.05.2020 um 10:04 Uhr"},
		{"2111-11-11T11:11:11+01:00", "11.11.2111 um 11:11 Uhr"},
		{"2020-05-31T", ""},
	}
	for _, date := range dates {
		out := formatDate(date.in)
		if out != date.want {
			t.Errorf("Error while formatting %q. Expected: %q, but got %q.", date.in, date.want, out)
		}
	}
}
