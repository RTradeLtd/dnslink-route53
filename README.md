# dnslink-route53

Used to deploy dnslink txt records to AWS Route53

Can be used as a library, or as a cli

## Usage - CLI

To install run `make install`

By default it uses ENV authentication which expects the following environment variables to be set:

* `AWS_ACCESS_KEY_ID`
* `AWS_SECRET_ACCESS_KEY`
* Optinonally `AWS_SECURITY_TOKEN`

If your AWS account needs a security token, you have to use the env auth method.

You can provide a region if you want, but it looks like AWS Route53 doesn't require a specific region, so it should be fine to leave the default at `us-east-1.

You *must* provide the zone id of your hosted zone.

Example invocation using env auth:

```shell
$ dlink53 -record.name _dnslink.foo.bar -record.value dnslink=/ipns/foo/bar -zone.id <zone-ID>

successfully deployed dnslink entry to route53
```