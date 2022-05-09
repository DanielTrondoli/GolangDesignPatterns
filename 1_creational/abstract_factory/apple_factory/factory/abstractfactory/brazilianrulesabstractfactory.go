package abstractfactory

import (
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/certificate"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/packing"
)

type BrazilianRulesAbstractFactory struct {
	ICountryRulesAbstractFactory
}

func (r BrazilianRulesAbstractFactory) GetCertificates() certificate.ICertificate {
	return new(certificate.BrazilianCertificate)
}

func (r BrazilianRulesAbstractFactory) GetPacking() packing.IPacking {
	return new(packing.BrazilianPacking)
}
