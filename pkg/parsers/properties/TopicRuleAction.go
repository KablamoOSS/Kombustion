package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TopicRuleAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html
type TopicRuleAction struct {
	StepFunctions    interface{} `yaml:"StepFunctions,omitempty"`
	Sqs              interface{} `yaml:"Sqs,omitempty"`
	Sns              interface{} `yaml:"Sns,omitempty"`
	S3               interface{} `yaml:"S3,omitempty"`
	Republish        interface{} `yaml:"Republish,omitempty"`
	Lambda           interface{} `yaml:"Lambda,omitempty"`
	Kinesis          interface{} `yaml:"Kinesis,omitempty"`
	IotAnalytics     interface{} `yaml:"IotAnalytics,omitempty"`
	Firehose         interface{} `yaml:"Firehose,omitempty"`
	Elasticsearch    interface{} `yaml:"Elasticsearch,omitempty"`
	DynamoDBv2       interface{} `yaml:"DynamoDBv2,omitempty"`
	DynamoDB         interface{} `yaml:"DynamoDB,omitempty"`
	CloudwatchMetric interface{} `yaml:"CloudwatchMetric,omitempty"`
	CloudwatchAlarm  interface{} `yaml:"CloudwatchAlarm,omitempty"`
}

// TopicRuleAction validation
func (resource TopicRuleAction) Validate() []error {
	errors := []error{}

	return errors
}
