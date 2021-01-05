package zipcode2prefs

import (
	"reflect"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	specs := []struct {
		Input  string
		Expect []string
	}{
		{
			Input:  "1500013",
			Expect: []string{"東京都"},
		},
		{
			Input:  "4980000",
			Expect: []string{"三重県", "愛知県"},
		},
		{
			Input:  "111",
			Expect: nil,
		},
	}

	for _, s := range specs {
		if e, g := s.Expect, Get(s.Input); !reflect.DeepEqual(e, g) {
			t.Errorf("expected %v but got %v", e, g)
		}
	}
}

func TestDups(t *testing.T) {
	dups := map[string]string{}
	for k, v := range prefs {
		if len(strings.Split(v, ",")) > 1 {
			dups[k] = v
		}
	}

	expected := map[string]string{
		"4980000": "三重県,愛知県",
		"6180000": "京都府,大阪府",
		"8710000": "大分県,福岡県",
	}

	if !reflect.DeepEqual(expected, dups) {
		t.Errorf("expected %v but got %v", expected, dups)
	}
}
