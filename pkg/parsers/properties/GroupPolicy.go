package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// GroupPolicy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type GroupPolicy struct {
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	PolicyName     interface{} `yaml:"PolicyName"`
}

// GroupPolicy validation
func (resource GroupPolicy) Validate() []error {
	errors := []error{}

	return errors
}
