package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TopicRuleTopicRulePayload Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html
type TopicRuleTopicRulePayload struct {
	AwsIotSqlVersion interface{} `yaml:"AwsIotSqlVersion,omitempty"`
	Description      interface{} `yaml:"Description,omitempty"`
	RuleDisabled     interface{} `yaml:"RuleDisabled"`
	Sql              interface{} `yaml:"Sql"`
	Actions          interface{} `yaml:"Actions"`
	ErrorAction      interface{} `yaml:"ErrorAction,omitempty"`
}

// TopicRuleTopicRulePayload validation
func (resource TopicRuleTopicRulePayload) Validate() []error {
	errors := []error{}

	return errors
}
