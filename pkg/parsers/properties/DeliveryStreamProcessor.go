package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeliveryStreamProcessor Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processor.html
type DeliveryStreamProcessor struct {
	Type       interface{} `yaml:"Type"`
	Parameters interface{} `yaml:"Parameters"`
}

// DeliveryStreamProcessor validation
func (resource DeliveryStreamProcessor) Validate() []error {
	errors := []error{}

	return errors
}
