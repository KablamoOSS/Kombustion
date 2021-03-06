package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TaskDefinitionUlimit Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions-ulimit.html
type TaskDefinitionUlimit struct {
	HardLimit interface{} `yaml:"HardLimit"`
	Name      interface{} `yaml:"Name"`
	SoftLimit interface{} `yaml:"SoftLimit"`
}

// TaskDefinitionUlimit validation
func (resource TaskDefinitionUlimit) Validate() []error {
	errors := []error{}

	return errors
}
