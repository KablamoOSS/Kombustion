package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ListenerRuleAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-actions.html
type ListenerRuleAction struct {
	TargetGroupArn interface{} `yaml:"TargetGroupArn"`
	Type           interface{} `yaml:"Type"`
}

// ListenerRuleAction validation
func (resource ListenerRuleAction) Validate() []error {
	errors := []error{}

	return errors
}
