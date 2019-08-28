package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PipelineDeviceRegistryEnrich Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-deviceregistryenrich.html
type PipelineDeviceRegistryEnrich struct {
	Attribute interface{} `yaml:"Attribute,omitempty"`
	Name      interface{} `yaml:"Name,omitempty"`
	Next      interface{} `yaml:"Next,omitempty"`
	RoleArn   interface{} `yaml:"RoleArn,omitempty"`
	ThingName interface{} `yaml:"ThingName,omitempty"`
}

// PipelineDeviceRegistryEnrich validation
func (resource PipelineDeviceRegistryEnrich) Validate() []error {
	errors := []error{}

	return errors
}