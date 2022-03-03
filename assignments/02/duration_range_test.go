package main

import (
	"testing"
	"time"
)

func TestDurationRange_RandomDuration(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		from, to time.Duration
		allSame  bool
	}{
		{
			name:    "Empty range returns always the same value",
			from:    time.Millisecond,
			to:      time.Millisecond,
			allSame: true,
		},
		{
			name:    "Non-empty range returns a value within the interval",
			from:    time.Millisecond,
			to:      100 * time.Millisecond,
			allSame: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			durationRange := NewDurationRange(tc.from, tc.to)

			set := set{}
			for i := 0; i < 10; i++ {
				duration := durationRange.RandomDuration()
				if tc.from > duration || duration > tc.to {
					t.Fatalf("Duration %v is outside of range [%v, %v]", duration, tc.from, tc.to)
				}

				set.Add(duration)
			}

			uniqValues := set.Values()
			if tc.allSame && len(uniqValues) != 1 {
				t.Fatalf("Expected all values the same, but got multiple different values: %v", uniqValues)
			}
			if !tc.allSame && len(uniqValues) < 2 {
				t.Fatalf("Expected multiple different values, but got: %v", uniqValues)
			}
		})
	}
}

type set map[interface{}]bool

func (s set) Add(item interface{}) {
	s[item] = true
}

func (s set) Values() []interface{} {
	items := []interface{}{}

	for key := range s {
		items = append(items, key)
	}

	return items
}
