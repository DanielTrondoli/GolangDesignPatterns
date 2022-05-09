package abstractfactory

import (
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/certificate"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/packing"
)

type USRulesAbstractFactory struct {
	ICountryRulesAbstractFactory
}

func (r USRulesAbstractFactory) GetCertificates() certificate.ICertificate {
	return new(certificate.USCertificate)
}

func (r USRulesAbstractFactory) GetPacking() packing.IPacking {
	return new(packing.USPacking)
}
