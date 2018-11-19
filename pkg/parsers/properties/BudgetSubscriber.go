package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BudgetSubscriber Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html
type BudgetSubscriber struct {
	Address          interface{} `yaml:"Address"`
	SubscriptionType interface{} `yaml:"SubscriptionType"`
}

// BudgetSubscriber validation
func (resource BudgetSubscriber) Validate() []error {
	errors := []error{}

	return errors
}
