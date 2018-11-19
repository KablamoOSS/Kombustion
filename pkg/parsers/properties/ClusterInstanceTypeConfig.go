package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterInstanceTypeConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html
type ClusterInstanceTypeConfig struct {
	BidPrice                            interface{} `yaml:"BidPrice,omitempty"`
	BidPriceAsPercentageOfOnDemandPrice interface{} `yaml:"BidPriceAsPercentageOfOnDemandPrice,omitempty"`
	InstanceType                        interface{} `yaml:"InstanceType"`
	WeightedCapacity                    interface{} `yaml:"WeightedCapacity,omitempty"`
	Configurations                      interface{} `yaml:"Configurations,omitempty"`
	EbsConfiguration                    interface{} `yaml:"EbsConfiguration,omitempty"`
}

// ClusterInstanceTypeConfig validation
func (resource ClusterInstanceTypeConfig) Validate() []error {
	errors := []error{}

	return errors
}
