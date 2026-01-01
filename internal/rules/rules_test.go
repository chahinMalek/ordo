package rules

import "testing"

func TestResolve(t *testing.T) {
	rules := map[string]Rule{
		"images": {Extensions: []string{"jpg", "png"}},
	}
	r := NewResolver(rules)

	test_cases := []struct {
		name      string
		filename  string
		expected  string
		useGroups bool
	}{
		{"resolve_with_group", "photo.jpg", "images", true},
		{"resolve_without_groups", "photo.jpg", "jpg", false},
		{"resolve_without_extension", "README", "", true},
		{"resolve_no_rule", "data.csv", "csv", true},
	}

	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			result := r.Resolve(tc.filename, tc.useGroups)
			if result != tc.expected {
				t.Errorf("Resolve(%s) = %s; want %s", tc.filename, result, tc.expected)
			}
		})
	}
}
