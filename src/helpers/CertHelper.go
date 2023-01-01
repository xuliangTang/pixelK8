package helpers

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/models"
)

func getCertType(alg x509.PublicKeyAlgorithm) string {
	switch alg {
	case x509.RSA:
		return "RSA"
	case x509.DSA:
		return "DSA"
	case x509.ECDSA:
		return "ECDSA"
	case x509.UnknownPublicKeyAlgorithm:
		return "Unknown"
	}
	return "Unknown"
}

// ParseCert 解析证书
func ParseCert(crt []byte) *models.CertModel {
	var cert tls.Certificate
	// 解码证书
	certBlock, restPEMBlock := pem.Decode(crt)
	if certBlock == nil {
		return nil
	}
	cert.Certificate = append(cert.Certificate, certBlock.Bytes)
	// 处理证书链
	certBlockChain, _ := pem.Decode(restPEMBlock)
	if certBlockChain != nil {
		cert.Certificate = append(cert.Certificate, certBlockChain.Bytes)
	}

	// 解析证书
	x509Cert, err := x509.ParseCertificate(certBlock.Bytes)

	if err != nil {
		return nil
	}

	return &models.CertModel{
		CN:        x509Cert.Subject.CommonName,
		Issuer:    x509Cert.Issuer.CommonName,
		Algorithm: getCertType(x509Cert.PublicKeyAlgorithm),
		BeginTime: x509Cert.NotBefore.Format(athena.DateTimeFormat),
		EndTime:   x509Cert.NotAfter.Format(athena.DateTimeFormat),
	}
}
