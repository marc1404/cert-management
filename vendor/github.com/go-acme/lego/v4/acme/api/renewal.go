package api

import (
	"errors"
	"net/http"

	"github.com/go-acme/lego/v4/acme"
)

// ErrNoARI is returned when the server does not advertise a renewal info endpoint.
var ErrNoARI = errors.New("renewalInfo[get/post]: server does not advertise a renewal info endpoint")

// GetRenewalInfo GETs renewal information for a certificate from the renewalInfo endpoint.
// This is used to determine if a certificate needs to be renewed.
//
// Note: this endpoint is part of a draft specification, not all ACME servers will implement it.
// This method will return api.ErrNoARI if the server does not advertise a renewal info endpoint.
//
// https://datatracker.ietf.org/doc/draft-ietf-acme-ari
func (c *CertificateService) GetRenewalInfo(certID string) (*http.Response, error) {
	if c.core.GetDirectory().RenewalInfo == "" {
		return nil, ErrNoARI
	}

	if certID == "" {
		return nil, errors.New("renewalInfo[get]: 'certID' cannot be empty")
	}

	return c.core.HTTPClient.Get(c.core.GetDirectory().RenewalInfo + "/" + certID)
}

// UpdateRenewalInfo POSTs updated renewal information for a certificate to the renewalInfo endpoint.
// This is used to indicate that a certificate has been replaced.
//
// Note: this endpoint is part of a draft specification, not all ACME servers will implement it.
// This method will return api.ErrNoARI if the server does not advertise a renewal info endpoint.
//
// https://datatracker.ietf.org/doc/draft-ietf-acme-ari
func (c *CertificateService) UpdateRenewalInfo(req acme.RenewalInfoUpdateRequest) (*http.Response, error) {
	if c.core.GetDirectory().RenewalInfo == "" {
		return nil, ErrNoARI
	}

	if req.CertID == "" {
		return nil, errors.New("renewalInfo[post]: 'certID' cannot be empty")
	}

	if !req.Replaced {
		return nil, errors.New("renewalInfo[post]: 'replaced' cannot be false")
	}

	return c.core.post(c.core.GetDirectory().RenewalInfo, req, nil)
}
