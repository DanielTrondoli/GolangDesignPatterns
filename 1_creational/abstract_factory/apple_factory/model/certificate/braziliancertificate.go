package certificate

type BrazilianCertificate struct {
	ICertificate
}

func (c BrazilianCertificate) ApplyCertificate() string {
	return "\t- Calibrating Brasilian rules\n\t- Applying Anatel's Stamp"
}
