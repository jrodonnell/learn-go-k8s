package kubecsr

import (

)

func newCSR(ktls *internalv1alpha1.SelfSignedTLSBundle) {
	csrpem := <-ktls.CSRPEM
	name := strings.ToLower(ktls.TypeMeta.GetObjectKind().GroupVersionKind().Kind + "-" + ktls.ObjectMeta.Name)
	kubecsr := certsv1.CertificateSigningRequest{
		TypeMeta:   ktls.TypeMeta,
		ObjectMeta: internalv1alpha1.NewG8sObjectMeta(ktls, name),
		Spec: certsv1.CertificateSigningRequestSpec{
			Request:    csrpem,
			SignerName: certsv1.KubeletServingSignerName,
			Usages:     []certsv1.KeyUsage{certsv1.UsageServerAuth, certsv1.UsageDigitalSignature, certsv1.UsageKeyEncipherment},
		},
	}
	pendingcsr, err := ktls.CertificateSigningRequests().Create(context.TODO(), &kubecsr, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err)
	}

	pendingcsr.Status.Conditions = append(pendingcsr.Status.Conditions, certsv1.CertificateSigningRequestCondition{
		Type:           certsv1.CertificateApproved,
		Status:         "True",
		Reason:         "G8s Approved",
		Message:        "This CSR was generated as part of a SelfSignedTLSBundle Object and approved by g8s-controller",
		LastUpdateTime: metav1.Now(),
	})
	approvedcsr, _ := ktls.CertificateSigningRequests().UpdateApproval(context.TODO(), pendingcsr.ObjectMeta.Name, pendingcsr, metav1.UpdateOptions{})
	certpem := approvedcsr.Status.Certificate

	for {
		if certpem != nil {
			ktls.CertPEM <- certpem
			break
		} else {
			approvedcsr, _ = ktls.CertificateSigningRequests().Get(context.TODO(), pendingcsr.ObjectMeta.Name, metav1.GetOptions{})
			certpem = approvedcsr.Status.Certificate
		}
	}
}
