package agencies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"log"
)

func GetProvider(opts gophercloud.AuthOptions) (*gophercloud.ProviderClient, error) {
	return openstack.AuthenticatedClient(opts)
}

func AuthenticatedClient(opts gophercloud.AuthOptions, tenant, agency string) (*gophercloud.ProviderClient, error) {
	identity, err := getIdentityClient(opts)
	if err != nil {
		return nil, err
	}
	agencyOpts := newAgencyOptions(tenant, agency)
	token, err := create(identity, agencyOpts).ExtractTokenID()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v", token)
	return getTokenProvider(opts, token)
}

func getTokenProvider(opts gophercloud.AuthOptions, token string) (*gophercloud.ProviderClient, error) {
	opts.TokenID = token
	opts.Username = ""
    opts.Password = ""
	opts.UserID = ""
	opts.DomainID = ""
	opts.DomainName = ""
	provider, err := GetProvider(opts)
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func getIdentityClient(opts gophercloud.AuthOptions) (*gophercloud.ServiceClient, error) {
	opts.Scope = new(gophercloud.AuthScope)
	opts.Scope.DomainName = opts.DomainName
	provider, err := GetProvider(opts)
	if err != nil {
		return nil, err
	}

	client, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		return nil, err
	}

	return client, nil
}

type AgencyAuthOptions struct {
	DomainName string `json:"name,omitempty"`
	AgencyName string
}

func (opts *AgencyAuthOptions) ToTokenV3CreateMap(scope map[string]interface{}) (map[string]interface{}, error) {
	type agencyReq struct {
		Domain *string `json:"domain_name"`
		Agency *string `json:"xrole_name"`
	}

	type identityReq struct {
		Methods []string   `json:"methods"`
		Agency  *agencyReq `json:"assume_role"`
	}

	type authReq struct {
		Identity identityReq `json:"identity"`
	}

	type request struct {
		Auth authReq `json:"auth"`
	}

	var req request
	req.Auth.Identity.Methods = []string{"assume_role"}
	req.Auth.Identity.Agency = &agencyReq{
		Agency: &opts.AgencyName,
		Domain: &opts.DomainName,
	}

	b, err := gophercloud.BuildRequestBody(req, "")
	if err != nil {
		return nil, err
	}

	if len(scope) != 0 {
		b["auth"].(map[string]interface{})["scope"] = scope
	}
	return b, nil
}

func (opts *AgencyAuthOptions) ToTokenV3ScopeMap() (map[string]interface{}, error) {
	return map[string]interface{}{
		"domain": map[string]interface{}{
			"name": opts.DomainName,
		},
	}, nil
}

func (opts *AgencyAuthOptions) CanReauth() bool {
	return false
}

func newAgencyOptions(tenant, agency string) *AgencyAuthOptions {
	return &AgencyAuthOptions{
		DomainName: tenant,
		AgencyName: agency,
	}
}

// Create authenticates and either generates a new token, or changes the Scope
// of an existing token.
func create(c *gophercloud.ServiceClient, opts tokens.AuthOptionsBuilder) (r tokens.CreateResult) {
	scope, err := opts.ToTokenV3ScopeMap()
	if err != nil {
		r.Err = err
		return
	}

	b, err := opts.ToTokenV3CreateMap(scope)
	if err != nil {
		r.Err = err
		return
	}

	resp, err := c.Post(tokenURL(c), b, &r.Body, &gophercloud.RequestOpts{
		// TODO: use safe method to get token id
		MoreHeaders: map[string]string{"X-Auth-Token": c.TokenID},
	})
	r.Err = err
	if resp != nil {
		r.Header = resp.Header
	}
	return
}
