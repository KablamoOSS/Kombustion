package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EC2FleetSpotOptionsRequest Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-spotoptionsrequest.html
type EC2FleetSpotOptionsRequest struct {
	AllocationStrategy           interface{} `yaml:"AllocationStrategy,omitempty"`
	InstanceInterruptionBehavior interface{} `yaml:"InstanceInterruptionBehavior,omitempty"`
	InstancePoolsToUseCount      interface{} `yaml:"InstancePoolsToUseCount,omitempty"`
}

// EC2FleetSpotOptionsRequest validation
func (resource EC2FleetSpotOptionsRequest) Validate() []error {
	errors := []error{}

	return errors
}
