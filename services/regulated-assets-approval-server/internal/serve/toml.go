package serve

import (
	"fmt"
	"net/http"
)

type stellarTOMLHandler struct {
	assetCode         string
	approvalServer    string
	issuerAddress     string
	networkPassphrase string
	kycThreshold      int64
}

func (h stellarTOMLHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	approvalCriteria := fmt.Sprintf("Currently %s is not approving any %s transactions", h.approvalServer, h.assetCode)
	fmt.Fprintf(rw, "NETWORK_PASSPHRASE=%q\n", h.networkPassphrase)
	fmt.Fprintf(rw, "[[CURRENCIES]]\n")
	fmt.Fprintf(rw, "code=%q\n", h.assetCode)
	fmt.Fprintf(rw, "issuer=%q\n", h.issuerAddress)
	fmt.Fprintf(rw, "regulated=true\n")
	fmt.Fprintf(rw, "approval_server=%q\n", h.approvalServer)
	fmt.Fprintf(rw, "approval_criteria=%q\n", approvalCriteria)
}
