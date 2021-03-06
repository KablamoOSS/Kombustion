package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DeploymentGroupDeployment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html
type DeploymentGroupDeployment struct {
	Description                   interface{} `yaml:"Description,omitempty"`
	IgnoreApplicationStopFailures interface{} `yaml:"IgnoreApplicationStopFailures,omitempty"`
	Revision                      interface{} `yaml:"Revision"`
}

// DeploymentGroupDeployment validation
func (resource DeploymentGroupDeployment) Validate() []error {
	errors := []error{}

	return errors
}
