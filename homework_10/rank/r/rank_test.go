package rank

import "testing"

func isSame(map1 map[string]int, map2 map[string]int) bool {
	return map1["hundreds"] == map2["hundreds"] && map1["dozens"] == map2["dozens"] &&
	map1["units"] == map2["units"]
}

func TestNumber_NumberRank(t *testing.T) {
	TestCase := []struct {
		name     string
		input    int
		expected map[string]int
	}{
		{name: "case one",
			input: 123,
			expected: map[string]int{
				"hundreds": 1,
				"dozens":   2,
				"units":    3,
			},
		},
		{
			name:  "case two",
			input: 752,
			expected: map[string]int{
				"hundreds": 7,
				"dozens":   5,
				"units":    2,
			},
		},
	}

	for _, cse := range TestCase {
		t.Run(cse.name, func(t *testing.T) {
			testMap := NumberRank(cse.input)
			if !isSame(testMap, cse.expected) {
				t.Errorf("expected %v got %v", cse.expected, testMap)
			}
		})
	}
}
