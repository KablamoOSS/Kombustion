package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ScalableTargetScheduledAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html
type ScalableTargetScheduledAction struct {
	EndTime              interface{} `yaml:"EndTime,omitempty"`
	Schedule             interface{} `yaml:"Schedule"`
	ScheduledActionName  interface{} `yaml:"ScheduledActionName"`
	StartTime            interface{} `yaml:"StartTime,omitempty"`
	ScalableTargetAction interface{} `yaml:"ScalableTargetAction,omitempty"`
}

// ScalableTargetScheduledAction validation
func (resource ScalableTargetScheduledAction) Validate() []error {
	errors := []error{}

	return errors
}
