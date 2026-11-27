package wsaa

import "testing"

func TestTestingEnvironmentURL(t *testing.T) {
	want := "https://wsaahomo.afip.gov.ar/ws/services/LoginCms"
	if got := Testing.URL(); got != want {
		t.Errorf("Testing.URL() = %q, want %q", got, want)
	}
}

func TestProductionEnvironmentURL(t *testing.T) {
	want := "https://wsaa.afip.gov.ar/ws/services/LoginCms"
	if got := Production.URL(); got != want {
		t.Errorf("Production.URL() = %q, want %q", got, want)
	}
}

func TestUnknownEnvironment(t *testing.T) {
	e := Environment(99)
	if got := e.URL(); got != "" {
		t.Errorf("URL() = %q, want empty string", got)
	}
	if got := e.String(); got != "unknown" {
		t.Errorf("String() = %q, want %q", got, "unknown")
	}
}

func TestEnvironmentString(t *testing.T) {
	if got := Testing.String(); got != "testing" {
		t.Errorf("Testing.String() = %q, want %q", got, "testing")
	}
	if got := Production.String(); got != "production" {
		t.Errorf("Production.String() = %q, want %q", got, "production")
	}
}
