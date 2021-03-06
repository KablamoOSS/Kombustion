package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// BudgetNotification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html
type BudgetNotification struct {
	ComparisonOperator interface{} `yaml:"ComparisonOperator"`
	NotificationType   interface{} `yaml:"NotificationType"`
	Threshold          interface{} `yaml:"Threshold"`
	ThresholdType      interface{} `yaml:"ThresholdType,omitempty"`
}

// BudgetNotification validation
func (resource BudgetNotification) Validate() []error {
	errors := []error{}

	return errors
}
