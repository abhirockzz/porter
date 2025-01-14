package config

import "testing"

func TestIsCoreAction(t *testing.T) {
	testcases := map[string]bool{
		"install":   true,
		"upgrade":   true,
		"uninstall": true,
		"status":    false,
		"INSTALL":   false,
	}

	for action, want := range testcases {
		t.Run(action, func(t *testing.T) {
			got := IsCoreAction(Action(action))
			if want != got {
				t.Fatalf("IsCoreAction(%q) failed, want %t, got %t", action, want, got)
			}
		})
	}
}
