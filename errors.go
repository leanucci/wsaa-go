package wsaa

import "fmt"

// CredentialsError tells that the certificate or the private key is missing,
// unreadable, or invalid, or that the two do not match.
// The client returns this error before it calls WSAA.
// The message never contains the private key, the TOKEN, or the SIGN.
type CredentialsError struct {
	// Message describes the problem.
	Message string
}

// Error returns the error message.
func (e *CredentialsError) Error() string {
	return "wsaa: credentials error: " + e.Message
}

// SoapFaultError tells that WSAA returned a SOAP fault.
type SoapFaultError struct {
	// FaultCode is the faultcode value from WSAA, for example "ns1:cms.cert.expired".
	FaultCode string
	// FaultString is the faultstring value from WSAA.
	FaultString string
}

// Error returns the error message. The message includes the fault code and the fault string.
func (e *SoapFaultError) Error() string {
	return fmt.Sprintf("wsaa: WSAA returned a SOAP fault (%s): %s", e.FaultCode, e.FaultString)
}

// ResponseError tells that the WSAA response is not valid, or that the HTTP request failed.
// The message never contains the response body, because the body can contain the TOKEN and the SIGN.
type ResponseError struct {
	// Message describes the problem.
	Message string
}

// Error returns the error message.
func (e *ResponseError) Error() string {
	return "wsaa: response error: " + e.Message
}
