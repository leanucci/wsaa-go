package wsaa

// Environment identifies a WSAA deployment: testing or production.
type Environment int

const (
	// Testing is the WSAA homologation environment (wsaahomo.afip.gov.ar).
	Testing Environment = iota
	// Production is the WSAA production environment (wsaa.afip.gov.ar).
	Production
)

// TestingURL is the LoginCms endpoint URL of the testing environment.
const TestingURL = "https://wsaahomo.afip.gov.ar/ws/services/LoginCms"

// ProductionURL is the LoginCms endpoint URL of the production environment.
const ProductionURL = "https://wsaa.afip.gov.ar/ws/services/LoginCms"

// URL returns the LoginCms endpoint URL of the environment.
// It returns an empty string for an unknown environment.
func (e Environment) URL() string {
	switch e {
	case Testing:
		return TestingURL
	case Production:
		return ProductionURL
	default:
		return ""
	}
}

// String returns the name of the environment: "testing", "production", or "unknown".
func (e Environment) String() string {
	switch e {
	case Testing:
		return "testing"
	case Production:
		return "production"
	default:
		return "unknown"
	}
}
