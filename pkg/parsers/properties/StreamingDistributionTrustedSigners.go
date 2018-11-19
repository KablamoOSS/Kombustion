package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// StreamingDistributionTrustedSigners Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-trustedsigners.html
type StreamingDistributionTrustedSigners struct {
	Enabled           interface{} `yaml:"Enabled"`
	AwsAccountNumbers interface{} `yaml:"AwsAccountNumbers,omitempty"`
}

// StreamingDistributionTrustedSigners validation
func (resource StreamingDistributionTrustedSigners) Validate() []error {
	errors := []error{}

	return errors
}
