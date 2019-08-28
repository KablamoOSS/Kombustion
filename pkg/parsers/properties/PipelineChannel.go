package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PipelineChannel Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-channel.html
type PipelineChannel struct {
	ChannelName interface{} `yaml:"ChannelName,omitempty"`
	Name        interface{} `yaml:"Name,omitempty"`
	Next        interface{} `yaml:"Next,omitempty"`
}

// PipelineChannel validation
func (resource PipelineChannel) Validate() []error {
	errors := []error{}

	return errors
}