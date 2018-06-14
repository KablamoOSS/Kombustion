package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// RDSDBSubnetGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html
type RDSDBSubnetGroup struct {
	Type       string                     `yaml:"Type"`
	Properties RDSDBSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// RDSDBSubnetGroup Properties
type RDSDBSubnetGroupProperties struct {
	DBSubnetGroupDescription interface{} `yaml:"DBSubnetGroupDescription"`
	DBSubnetGroupName        interface{} `yaml:"DBSubnetGroupName,omitempty"`
	SubnetIds                interface{} `yaml:"SubnetIds"`
	Tags                     interface{} `yaml:"Tags,omitempty"`
}

// NewRDSDBSubnetGroup constructor creates a new RDSDBSubnetGroup
func NewRDSDBSubnetGroup(properties RDSDBSubnetGroupProperties, deps ...interface{}) RDSDBSubnetGroup {
	return RDSDBSubnetGroup{
		Type:       "AWS::RDS::DBSubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRDSDBSubnetGroup parses RDSDBSubnetGroup
func ParseRDSDBSubnetGroup(name string, data string) (cf types.TemplateObject, err error) {
	var resource RDSDBSubnetGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSDBSubnetGroup - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseRDSDBSubnetGroup validator
func (resource RDSDBSubnetGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRDSDBSubnetGroupProperties validator
func (resource RDSDBSubnetGroupProperties) Validate() []error {
	errs := []error{}
	if resource.DBSubnetGroupDescription == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DBSubnetGroupDescription'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
