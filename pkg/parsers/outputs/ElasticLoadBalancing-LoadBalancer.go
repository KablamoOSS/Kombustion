package outputs

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ParseElasticLoadBalancingLoadBalancer Documentation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html
func ParseElasticLoadBalancingLoadBalancer(
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
	source = "kombustion-core-outputs"

	var resource, output types.TemplateObject

	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-ElasticLoadBalancingLoadBalancer-" + name,
				},
			},
		},
	}

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "CanonicalHostedZoneName"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElasticLoadBalancingLoadBalancer-" + name + "-CanonicalHostedZoneName",
			},
		},
	}

	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}

	outputs[name+"CanonicalHostedZoneName"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "CanonicalHostedZoneNameID"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElasticLoadBalancingLoadBalancer-" + name + "-CanonicalHostedZoneNameID",
			},
		},
	}

	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}

	outputs[name+"CanonicalHostedZoneNameID"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "DNSName"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElasticLoadBalancingLoadBalancer-" + name + "-DNSName",
			},
		},
	}

	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}

	outputs[name+"DNSName"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "SourceSecurityGroup.GroupName"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElasticLoadBalancingLoadBalancer-" + name + "-SourceSecurityGroupGroupName",
			},
		},
	}

	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}

	outputs[name+"SourceSecurityGroupGroupName"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "SourceSecurityGroup.OwnerAlias"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElasticLoadBalancingLoadBalancer-" + name + "-SourceSecurityGroupOwnerAlias",
			},
		},
	}

	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}

	outputs[name+"SourceSecurityGroupOwnerAlias"] = output

	return
}
