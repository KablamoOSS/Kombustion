package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// JobDefinitionVolumes Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html
type JobDefinitionVolumes struct {
	Name interface{} `yaml:"Name,omitempty"`
	Host interface{} `yaml:"Host,omitempty"`
}

// JobDefinitionVolumes validation
func (resource JobDefinitionVolumes) Validate() []error {
	errors := []error{}

	return errors
}
