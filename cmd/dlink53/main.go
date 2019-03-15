package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	dlink53 "github.com/RTradeLtd/dnslink-route53"
	"github.com/mitchellh/goamz/aws"
)

var (
	authMethod  = flag.String("name", "env", "set aws authentication method, valid values are env or get")
	accessKey   = flag.String("access.key", "", "aws access key to use if not using env auth")
	secretKey   = flag.String("secret.key", "", "aws secret key use if not using env auth method")
	region      = flag.String("region", "us-east-1", "the aws region your domain is hosted in")
	recordName  = flag.String("record.name", "", "the name of the dnslink record, ie _dnslink.foo.bar")
	recordValue = flag.String("record.value", "", "the value of the dnslink record, ie \"dnslink=/ipns/foo/bar\"")
)

func init() {
	flag.Parse()
	if *recordName == "" {
		log.Fatal("record.name cant be empty")
	}
	if *recordValue == "" {
		log.Fatal("record.value cant be empty")
	}
}

func returnAuthCredentials() ([]string, error) {
	if *accessKey == "" {
		return nil, errors.New("access.key flag is empty")
	} else if *secretKey == "" {
		return nil, errors.New("secret.key flag is empty")
	}
	return []string{*accessKey, *secretKey}, nil
}

func main() {
	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	var (
		authCredentials []string
		awsRegion       aws.Region
		err             error
	)
	if *authMethod == "get" {
		authCredentials, err = returnAuthCredentials()
		if err != nil {
			log.Fatal(err)
		}
	}
	awsRegion, ok := aws.Regions[*region]
	if !ok {
		log.Fatalf("%s is not a valid region", *region)
	}
	deployer, err := dlink53.NewDeployer(*authMethod, awsRegion, authCredentials...)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := deployer.AddEntry(*recordName, *recordValue); err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully deployed dnslink entry to route53")
}
