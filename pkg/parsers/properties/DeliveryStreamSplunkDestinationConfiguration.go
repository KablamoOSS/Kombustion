package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DeliveryStreamSplunkDestinationConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-splunkdestinationconfiguration.html
type DeliveryStreamSplunkDestinationConfiguration struct {
	HECAcknowledgmentTimeoutInSeconds interface{} `yaml:"HECAcknowledgmentTimeoutInSeconds,omitempty"`
	HECEndpoint                       interface{} `yaml:"HECEndpoint"`
	HECEndpointType                   interface{} `yaml:"HECEndpointType"`
	HECToken                          interface{} `yaml:"HECToken"`
	S3BackupMode                      interface{} `yaml:"S3BackupMode,omitempty"`
	RetryOptions                      interface{} `yaml:"RetryOptions,omitempty"`
	S3Configuration                   interface{} `yaml:"S3Configuration"`
	ProcessingConfiguration           interface{} `yaml:"ProcessingConfiguration,omitempty"`
	CloudWatchLoggingOptions          interface{} `yaml:"CloudWatchLoggingOptions,omitempty"`
}

// DeliveryStreamSplunkDestinationConfiguration validation
func (resource DeliveryStreamSplunkDestinationConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
