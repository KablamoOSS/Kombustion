package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// WebACLRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-rule.html
type WebACLRule struct {
	Priority interface{} `yaml:"Priority"`
	RuleId   interface{} `yaml:"RuleId"`
	Action   interface{} `yaml:"Action"`
}

// WebACLRule validation
func (resource WebACLRule) Validate() []error {
	errors := []error{}

	return errors
}
