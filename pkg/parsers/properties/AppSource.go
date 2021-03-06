package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// AppSource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html
type AppSource struct {
	Password interface{} `yaml:"Password,omitempty"`
	Revision interface{} `yaml:"Revision,omitempty"`
	SshKey   interface{} `yaml:"SshKey,omitempty"`
	Type     interface{} `yaml:"Type,omitempty"`
	Url      interface{} `yaml:"Url,omitempty"`
	Username interface{} `yaml:"Username,omitempty"`
}

// AppSource validation
func (resource AppSource) Validate() []error {
	errors := []error{}

	return errors
}
