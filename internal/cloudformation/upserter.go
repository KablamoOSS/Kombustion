package cloudformation

import (
	"fmt"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
)

type StackUpserter interface {
	Open(string, string) string
	DescribeStackEvents(*awsCF.DescribeStackEventsInput) (*awsCF.DescribeStackEventsOutput, error)
	DescribeStackEventsPages(*awsCF.DescribeStackEventsInput, func(*awsCF.DescribeStackEventsOutput, bool) bool) error
	DescribeStacks(*awsCF.DescribeStacksInput) (*awsCF.DescribeStacksOutput, error)
	CreateChangeSet(*awsCF.CreateChangeSetInput) (*awsCF.CreateChangeSetOutput, error)
	DescribeChangeSet(*awsCF.DescribeChangeSetInput) (*awsCF.DescribeChangeSetOutput, error)
	WaitUntilChangeSetCreateComplete(*awsCF.DescribeChangeSetInput) error
	DeleteChangeSet(*awsCF.DeleteChangeSetInput) (*awsCF.DeleteChangeSetOutput, error)
	DeleteStack(*awsCF.DeleteStackInput) (*awsCF.DeleteStackOutput, error)
	ExecuteChangeSet(*awsCF.ExecuteChangeSetInput) (*awsCF.ExecuteChangeSetOutput, error)
}

func (wr *Wrapper) DescribeStacks(in *awsCF.DescribeStacksInput) (*awsCF.DescribeStacksOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.DescribeStacks(in)
}

func (wr *Wrapper) DeleteStack(in *awsCF.DeleteStackInput) (*awsCF.DeleteStackOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.DeleteStack(in)
}

func (wr *Wrapper) CreateChangeSet(in *awsCF.CreateChangeSetInput) (*awsCF.CreateChangeSetOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.CreateChangeSet(in)
}

func (wr *Wrapper) DescribeChangeSet(in *awsCF.DescribeChangeSetInput) (*awsCF.DescribeChangeSetOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.DescribeChangeSet(in)
}

func (wr *Wrapper) WaitUntilChangeSetCreateComplete(in *awsCF.DescribeChangeSetInput) error {
	if wr.client == nil {
		return fmt.Errorf("connection not opened")
	}
	return wr.client.WaitUntilChangeSetCreateComplete(in)
}

func (wr *Wrapper) ExecuteChangeSet(in *awsCF.ExecuteChangeSetInput) (*awsCF.ExecuteChangeSetOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.ExecuteChangeSet(in)
}

func (wr *Wrapper) DeleteChangeSet(in *awsCF.DeleteChangeSetInput) (*awsCF.DeleteChangeSetOutput, error) {
	if wr.client == nil {
		return nil, fmt.Errorf("connection not opened")
	}
	return wr.client.DeleteChangeSet(in)
}
