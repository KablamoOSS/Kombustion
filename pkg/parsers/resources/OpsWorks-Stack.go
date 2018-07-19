package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// OpsWorksStack Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-stack.html
type OpsWorksStack struct {
	Type       string                  `yaml:"Type"`
	Properties OpsWorksStackProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// OpsWorksStack Properties
type OpsWorksStackProperties struct {
	AgentVersion              interface{}                                `yaml:"AgentVersion,omitempty"`
	ClonePermissions          interface{}                                `yaml:"ClonePermissions,omitempty"`
	CustomJson                interface{}                                `yaml:"CustomJson,omitempty"`
	DefaultAvailabilityZone   interface{}                                `yaml:"DefaultAvailabilityZone,omitempty"`
	DefaultInstanceProfileArn interface{}                                `yaml:"DefaultInstanceProfileArn"`
	DefaultOs                 interface{}                                `yaml:"DefaultOs,omitempty"`
	DefaultRootDeviceType     interface{}                                `yaml:"DefaultRootDeviceType,omitempty"`
	DefaultSshKeyName         interface{}                                `yaml:"DefaultSshKeyName,omitempty"`
	DefaultSubnetId           interface{}                                `yaml:"DefaultSubnetId,omitempty"`
	EcsClusterArn             interface{}                                `yaml:"EcsClusterArn,omitempty"`
	HostnameTheme             interface{}                                `yaml:"HostnameTheme,omitempty"`
	Name                      interface{}                                `yaml:"Name"`
	ServiceRoleArn            interface{}                                `yaml:"ServiceRoleArn"`
	SourceStackId             interface{}                                `yaml:"SourceStackId,omitempty"`
	UseCustomCookbooks        interface{}                                `yaml:"UseCustomCookbooks,omitempty"`
	UseOpsworksSecurityGroups interface{}                                `yaml:"UseOpsworksSecurityGroups,omitempty"`
	VpcId                     interface{}                                `yaml:"VpcId,omitempty"`
	ConfigurationManager      *properties.StackStackConfigurationManager `yaml:"ConfigurationManager,omitempty"`
	CustomCookbooksSource     *properties.StackSource                    `yaml:"CustomCookbooksSource,omitempty"`
	Attributes                interface{}                                `yaml:"Attributes,omitempty"`
	Tags                      interface{}                                `yaml:"Tags,omitempty"`
	CloneAppIds               interface{}                                `yaml:"CloneAppIds,omitempty"`
	ElasticIps                interface{}                                `yaml:"ElasticIps,omitempty"`
	RdsDbInstances            interface{}                                `yaml:"RdsDbInstances,omitempty"`
	ChefConfiguration         *properties.StackChefConfiguration         `yaml:"ChefConfiguration,omitempty"`
}

// NewOpsWorksStack constructor creates a new OpsWorksStack
func NewOpsWorksStack(properties OpsWorksStackProperties, deps ...interface{}) OpsWorksStack {
	return OpsWorksStack{
		Type:       "AWS::OpsWorks::Stack",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseOpsWorksStack parses OpsWorksStack
func ParseOpsWorksStack(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource OpsWorksStack
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseOpsWorksStack validator
func (resource OpsWorksStack) Validate() []error {
	return resource.Properties.Validate()
}

// ParseOpsWorksStackProperties validator
func (resource OpsWorksStackProperties) Validate() []error {
	errors := []error{}
	if resource.DefaultInstanceProfileArn == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DefaultInstanceProfileArn'"))
	}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.ServiceRoleArn == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ServiceRoleArn'"))
	}
	return errors
}
