package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// HostedZoneVPC Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html
type HostedZoneVPC struct {
	VPCId     interface{} `yaml:"VPCId"`
	VPCRegion interface{} `yaml:"VPCRegion"`
}

// HostedZoneVPC validation
func (resource HostedZoneVPC) Validate() []error {
	errors := []error{}

	return errors
}
