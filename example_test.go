package wsaa_test

import (
	"fmt"

	wsaa "github.com/leanucci/wsaa-go"
)

// ExampleEnvironment runs the same code as the Usage section of README.md.
func ExampleEnvironment() {
	env := wsaa.Testing
	fmt.Println(env.String())
	fmt.Println(env.URL())

	env = wsaa.Production
	fmt.Println(env.String())
	fmt.Println(env.URL())
	// Output:
	// testing
	// https://wsaahomo.afip.gov.ar/ws/services/LoginCms
	// production
	// https://wsaa.afip.gov.ar/ws/services/LoginCms
}
