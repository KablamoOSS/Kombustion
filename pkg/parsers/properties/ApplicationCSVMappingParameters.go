package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ApplicationCSVMappingParameters Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-csvmappingparameters.html
type ApplicationCSVMappingParameters struct {
	RecordColumnDelimiter interface{} `yaml:"RecordColumnDelimiter"`
	RecordRowDelimiter    interface{} `yaml:"RecordRowDelimiter"`
}

// ApplicationCSVMappingParameters validation
func (resource ApplicationCSVMappingParameters) Validate() []error {
	errors := []error{}

	return errors
}
