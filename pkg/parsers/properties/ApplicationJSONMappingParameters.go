package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ApplicationJSONMappingParameters Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-jsonmappingparameters.html
type ApplicationJSONMappingParameters struct {
	RecordRowPath interface{} `yaml:"RecordRowPath"`
}

// ApplicationJSONMappingParameters validation
func (resource ApplicationJSONMappingParameters) Validate() []error {
	errors := []error{}

	return errors
}
