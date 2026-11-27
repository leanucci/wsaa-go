# wsaa-go

A Go client for AFIP's WSAA (Web Service de Autenticación y Autorización).

A Go port of [wsaa-ruby](https://github.com/leanucci/wsaa-ruby) and
[wsaa-node](https://github.com/leanucci/wsaa-node). The client builds a login ticket request
(TRA), signs it as a CMS (PKCS#7) message with your certificate and private key, sends it to
WSAA `loginCms`, and returns the TOKEN and the SIGN. Other AFIP web services need these
credentials.
