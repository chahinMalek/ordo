package rules

import "testing"

func TestResolve(t *testing.T) {
	rules := []Rule{
		{TargetDir: "images", Extensions: []string{"jpg", "png"}},
	}
	r := NewResolver(rules, "unknown")

	test_cases := []struct {
		name     string
		filename string
		expected string
	}{
		{"rule_match", "photo.jpg", "images"},
		{"fallback_no_extension", "README", "unknown"},
		{"fallback_no_rule", "data.csv", "csv"},
	}

	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			result := r.Resolve(tc.filename)
			if result != tc.expected {
				t.Errorf("Resolve(%s) = %s; want %s", tc.filename, result, tc.expected)
			}
		})
	}
}
