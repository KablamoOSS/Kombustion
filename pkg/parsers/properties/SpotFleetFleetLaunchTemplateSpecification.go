package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// SpotFleetFleetLaunchTemplateSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-fleetlaunchtemplatespecification.html
type SpotFleetFleetLaunchTemplateSpecification struct {
	LaunchTemplateId   interface{} `yaml:"LaunchTemplateId,omitempty"`
	LaunchTemplateName interface{} `yaml:"LaunchTemplateName,omitempty"`
	Version            interface{} `yaml:"Version,omitempty"`
}

// SpotFleetFleetLaunchTemplateSpecification validation
func (resource SpotFleetFleetLaunchTemplateSpecification) Validate() []error {
	errs := []error{}

	return errs
}
