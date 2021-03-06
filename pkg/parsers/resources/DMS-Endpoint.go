package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// DMSEndpoint Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html
type DMSEndpoint struct {
	Type       string                `yaml:"Type"`
	Properties DMSEndpointProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

// DMSEndpoint Properties
type DMSEndpointProperties struct {
	CertificateArn            interface{} `yaml:"CertificateArn,omitempty"`
	DatabaseName              interface{} `yaml:"DatabaseName,omitempty"`
	EndpointIdentifier        interface{} `yaml:"EndpointIdentifier,omitempty"`
	EndpointType              interface{} `yaml:"EndpointType"`
	EngineName                interface{} `yaml:"EngineName"`
	ExtraConnectionAttributes interface{} `yaml:"ExtraConnectionAttributes,omitempty"`
	KmsKeyId                  interface{} `yaml:"KmsKeyId,omitempty"`
	Password                  interface{} `yaml:"Password,omitempty"`
	Port                      interface{} `yaml:"Port,omitempty"`
	ServerName                interface{} `yaml:"ServerName,omitempty"`
	SslMode                   interface{} `yaml:"SslMode,omitempty"`
	Username                  interface{} `yaml:"Username,omitempty"`
	S3Settings                interface{} `yaml:"S3Settings,omitempty"`
	MongoDbSettings           interface{} `yaml:"MongoDbSettings,omitempty"`
	Tags                      interface{} `yaml:"Tags,omitempty"`
	KinesisSettings           interface{} `yaml:"KinesisSettings,omitempty"`
	ElasticsearchSettings     interface{} `yaml:"ElasticsearchSettings,omitempty"`
	DynamoDbSettings          interface{} `yaml:"DynamoDbSettings,omitempty"`
}

// NewDMSEndpoint constructor creates a new DMSEndpoint
func NewDMSEndpoint(properties DMSEndpointProperties, deps ...interface{}) DMSEndpoint {
	return DMSEndpoint{
		Type:       "AWS::DMS::Endpoint",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDMSEndpoint parses DMSEndpoint
func ParseDMSEndpoint(
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
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource DMSEndpoint
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

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-DMSEndpoint-" + name,
				},
			},
		},
	}

	return
}

// ParseDMSEndpoint validator
func (resource DMSEndpoint) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDMSEndpointProperties validator
func (resource DMSEndpointProperties) Validate() []error {
	errors := []error{}
	return errors
}
