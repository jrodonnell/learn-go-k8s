package kubecsr

import (

)

func genCSRPem() {
		pubbytes, _ := x509.MarshalPKIXPublicKey(&ecdsakey.PublicKey)
		pubblock := &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubbytes,
		}
		pubpem := string(pem.EncodeToMemory(pubblock))

		reference for creating kube CSR object

		x509csr := x509.CertificateRequest{
			SignatureAlgorithm: x509.ECDSAWithSHA256,
			PublicKeyAlgorithm: x509.ECDSA,
			PublicKey:          ecdsakey.PublicKey,
			Subject: pkix.Name{
				Organization: []string{"g8s"},
				CommonName:   sstls.Spec.AppName,
			},
			DNSNames: sstls.Spec.SANs,
		}

		csrbytes, _ := x509.CreateCertificateRequest(rand.Reader, &x509csr, ecdsakey)
		csrblock := &pem.Block{
			Type:  "CERTIFICATE REQUEST",
			Bytes: csrbytes,
		}
		csrpem := pem.EncodeToMemory(csrblock)
}
