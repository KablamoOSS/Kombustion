package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ApplicationInputProcessingConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html
type ApplicationInputProcessingConfiguration struct {
	InputLambdaProcessor interface{} `yaml:"InputLambdaProcessor,omitempty"`
}

// ApplicationInputProcessingConfiguration validation
func (resource ApplicationInputProcessingConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
