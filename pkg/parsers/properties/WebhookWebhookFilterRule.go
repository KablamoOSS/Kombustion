package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// WebhookWebhookFilterRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookfilterrule.html
type WebhookWebhookFilterRule struct {
	JsonPath    interface{} `yaml:"JsonPath"`
	MatchEquals interface{} `yaml:"MatchEquals,omitempty"`
}

// WebhookWebhookFilterRule validation
func (resource WebhookWebhookFilterRule) Validate() []error {
	errors := []error{}

	return errors
}
