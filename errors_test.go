package wsaa

import (
	"errors"
	"strings"
	"testing"
)

func TestCredentialsErrorMessage(t *testing.T) {
	var err error = &CredentialsError{Message: "private key does not match certificate"}
	if got := err.Error(); !strings.Contains(got, "private key does not match certificate") {
		t.Errorf("Error() = %q, want it to contain the message", got)
	}
}

func TestSoapFaultErrorMessage(t *testing.T) {
	var err error = &SoapFaultError{FaultCode: "ns1:cms.cert.expired", FaultString: "Certificado expirado"}
	got := err.Error()
	for _, want := range []string{"ns1:cms.cert.expired", "Certificado expirado"} {
		if !strings.Contains(got, want) {
			t.Errorf("Error() = %q, want it to contain %q", got, want)
		}
	}
}

func TestResponseErrorMessage(t *testing.T) {
	var err error = &ResponseError{Message: "missing loginCmsReturn element"}
	if got := err.Error(); !strings.Contains(got, "missing loginCmsReturn element") {
		t.Errorf("Error() = %q, want it to contain the message", got)
	}
}

func TestErrorTypesWorkWithErrorsAs(t *testing.T) {
	var err error = &SoapFaultError{FaultCode: "c", FaultString: "s"}
	var fault *SoapFaultError
	if !errors.As(err, &fault) || fault.FaultCode != "c" {
		t.Errorf("errors.As did not find *SoapFaultError")
	}
}
