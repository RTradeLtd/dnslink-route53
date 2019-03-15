package dlink53

import (
	"errors"
	"strings"

	"github.com/mitchellh/goamz/aws"
	r "github.com/mitchellh/goamz/route53"
	route53 "github.com/segmentio/go-route53"
)

type AwsLinkManager struct {
	Auth   aws.Auth
	Region aws.Region
	Client *route53.Client
	Zone   string
}

// Deployer is used to handle deployment of DNSLink TXT records to route53 domains
type Deployer struct {
	zone   string
	region aws.Region
	client *route53.Client
}

// NewDeployer is used to instantiate our DNSLink deployer
// if authMethod is set to get, credentials variadic parameter must be given
// the first element in the array is the access key, and the second is the secret key
func NewDeployer(authMethod string, region aws.Region, credentials ...string) (*Deployer, error) {
	var (
		auth aws.Auth
		err  error
	)
	switch authMethod {
	case "env":
		auth, err = aws.EnvAuth()
	case "get":
		if len(credentials) != 2 {
			return nil, errors.New("bad credentials provided")
		}
		auth, err = aws.GetAuth(credentials[0], credentials[1])
	default:
		err = errors.New("invalid auth method provided")
	}
	if err != nil {
		return nil, err
	}
	return &Deployer{
		zone:   region.Name,
		region: region,
		client: route53.New(auth, region),
	}, nil
}

// AddEntry adds a TXT entry to a domain-name with a ttl value of 300 seconds
func (d *Deployer) AddEntry(name, value string) (*r.ChangeResourceRecordSetsResponse, error) {
	if !strings.HasSuffix(name, "_dnslink.") {
		return nil, errors.New("invalid dnslink name")
	}
	return d.client.Zone(d.zone).Add("TXT", name, value)
}
