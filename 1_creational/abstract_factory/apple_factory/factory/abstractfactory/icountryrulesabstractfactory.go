package abstractfactory

import (
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/certificate"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/packing"
)

type ICountryRulesAbstractFactory interface {
	GetCertificates() certificate.ICertificate
	GetPacking() packing.IPacking
}
