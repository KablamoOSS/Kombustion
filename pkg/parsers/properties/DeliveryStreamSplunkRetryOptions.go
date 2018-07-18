package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeliveryStreamSplunkRetryOptions Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-splunkretryoptions.html
type DeliveryStreamSplunkRetryOptions struct {
	DurationInSeconds interface{} `yaml:"DurationInSeconds"`
}

// DeliveryStreamSplunkRetryOptions validation
func (resource DeliveryStreamSplunkRetryOptions) Validate() []error {
	errs := []error{}

	if resource.DurationInSeconds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DurationInSeconds'"))
	}
	return errs
}
