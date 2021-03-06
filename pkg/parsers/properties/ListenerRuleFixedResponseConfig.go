package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ListenerRuleFixedResponseConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-fixedresponseconfig.html
type ListenerRuleFixedResponseConfig struct {
	ContentType interface{} `yaml:"ContentType,omitempty"`
	MessageBody interface{} `yaml:"MessageBody,omitempty"`
	StatusCode  interface{} `yaml:"StatusCode"`
}

// ListenerRuleFixedResponseConfig validation
func (resource ListenerRuleFixedResponseConfig) Validate() []error {
	errors := []error{}

	return errors
}
