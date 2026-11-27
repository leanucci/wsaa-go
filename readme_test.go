package wsaa

import (
	"os"
	"strings"
	"testing"
)

func readREADME(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("read README.md: %v", err)
	}
	return string(b)
}

func TestREADMEHasNoPortfolioReference(t *testing.T) {
	readme := readREADME(t)
	for _, s := range []string{"leanucci/portfolio", "APPROACH.md"} {
		if strings.Contains(readme, s) {
			t.Errorf("README.md contains %q", s)
		}
	}
}

func TestREADMEKeepsClientDescription(t *testing.T) {
	readme := readREADME(t)
	for _, s := range []string{"WSAA", "TRA", "signs", "loginCms", "TOKEN", "SIGN"} {
		if !strings.Contains(readme, s) {
			t.Errorf("README.md does not contain %q", s)
		}
	}
}
