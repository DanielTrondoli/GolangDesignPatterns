package certificate

type USCertificate struct {
	ICertificate
}

func (c USCertificate) ApplyCertificate() string {
	return "\t- Calibrating US rules"
}
