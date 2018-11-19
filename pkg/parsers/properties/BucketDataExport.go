package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketDataExport Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html
type BucketDataExport struct {
	OutputSchemaVersion interface{} `yaml:"OutputSchemaVersion"`
	Destination         interface{} `yaml:"Destination"`
}

// BucketDataExport validation
func (resource BucketDataExport) Validate() []error {
	errors := []error{}

	return errors
}
