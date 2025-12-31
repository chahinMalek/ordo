package rules

import "testing"

func TestResolve(t *testing.T) {
	rules := []Rule{
		{TargetDir: "images", Extensions: []string{"jpg", "png"}},
	}
	r := NewResolver(rules, "unknown")

	test_cases := []struct {
		name        string
		filename    string
		expected    string
		useGroups   bool
		useFallback bool
	}{
		{"resolve_with_group", "photo.jpg", "images", true, true},
		{"resolve_without_groups", "photo.jpg", "jpg", false, true},
		{"resolve_without_fallback_no_extension", "README", "", true, false},
		{"resolve_with_fallback_no_extension", "README", "unknown", true, true},
		{"resolve_fallback_no_rule", "data.csv", "csv", true, true},
	}

	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			result := r.Resolve(tc.filename, tc.useGroups, tc.useFallback)
			if result != tc.expected {
				t.Errorf("Resolve(%s) = %s; want %s", tc.filename, result, tc.expected)
			}
		})
	}
}
