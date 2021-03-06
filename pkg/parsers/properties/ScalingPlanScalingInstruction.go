package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ScalingPlanScalingInstruction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html
type ScalingPlanScalingInstruction struct {
	DisableDynamicScaling                interface{} `yaml:"DisableDynamicScaling,omitempty"`
	MaxCapacity                          interface{} `yaml:"MaxCapacity"`
	MinCapacity                          interface{} `yaml:"MinCapacity"`
	PredictiveScalingMaxCapacityBehavior interface{} `yaml:"PredictiveScalingMaxCapacityBehavior,omitempty"`
	PredictiveScalingMaxCapacityBuffer   interface{} `yaml:"PredictiveScalingMaxCapacityBuffer,omitempty"`
	PredictiveScalingMode                interface{} `yaml:"PredictiveScalingMode,omitempty"`
	ResourceId                           interface{} `yaml:"ResourceId"`
	ScalableDimension                    interface{} `yaml:"ScalableDimension"`
	ScalingPolicyUpdateBehavior          interface{} `yaml:"ScalingPolicyUpdateBehavior,omitempty"`
	ScheduledActionBufferTime            interface{} `yaml:"ScheduledActionBufferTime,omitempty"`
	ServiceNamespace                     interface{} `yaml:"ServiceNamespace"`
	PredefinedLoadMetricSpecification    interface{} `yaml:"PredefinedLoadMetricSpecification,omitempty"`
	TargetTrackingConfigurations         interface{} `yaml:"TargetTrackingConfigurations"`
	CustomizedLoadMetricSpecification    interface{} `yaml:"CustomizedLoadMetricSpecification,omitempty"`
}

// ScalingPlanScalingInstruction validation
func (resource ScalingPlanScalingInstruction) Validate() []error {
	errors := []error{}

	return errors
}
