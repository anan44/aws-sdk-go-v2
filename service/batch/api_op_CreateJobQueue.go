// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateJobQueueInput struct {
	_ struct{} `type:"structure"`

	// The set of compute environments mapped to a job queue and their order relative
	// to each other. The job scheduler uses this parameter to determine which compute
	// environment should execute a given job. Compute environments must be in the
	// VALID state before you can associate them with a job queue. You can associate
	// up to three compute environments with a job queue.
	//
	// ComputeEnvironmentOrder is a required field
	ComputeEnvironmentOrder []ComputeEnvironmentOrder `locationName:"computeEnvironmentOrder" type:"list" required:"true"`

	// The name of the job queue.
	//
	// JobQueueName is a required field
	JobQueueName *string `locationName:"jobQueueName" type:"string" required:"true"`

	// The priority of the job queue. Job queues with a higher priority (or a higher
	// integer value for the priority parameter) are evaluated first when associated
	// with the same compute environment. Priority is determined in descending order,
	// for example, a job queue with a priority value of 10 is given scheduling
	// preference over a job queue with a priority value of 1.
	//
	// Priority is a required field
	Priority *int64 `locationName:"priority" type:"integer" required:"true"`

	// The state of the job queue. If the job queue state is ENABLED, it is able
	// to accept jobs.
	State JQState `locationName:"state" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateJobQueueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateJobQueueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateJobQueueInput"}

	if s.ComputeEnvironmentOrder == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComputeEnvironmentOrder"))
	}

	if s.JobQueueName == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobQueueName"))
	}

	if s.Priority == nil {
		invalidParams.Add(aws.NewErrParamRequired("Priority"))
	}
	if s.ComputeEnvironmentOrder != nil {
		for i, v := range s.ComputeEnvironmentOrder {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ComputeEnvironmentOrder", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJobQueueInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ComputeEnvironmentOrder != nil {
		v := s.ComputeEnvironmentOrder

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "computeEnvironmentOrder", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.JobQueueName != nil {
		v := *s.JobQueueName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobQueueName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Priority != nil {
		v := *s.Priority

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "priority", protocol.Int64Value(v), metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type CreateJobQueueOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the job queue.
	//
	// JobQueueArn is a required field
	JobQueueArn *string `locationName:"jobQueueArn" type:"string" required:"true"`

	// The name of the job queue.
	//
	// JobQueueName is a required field
	JobQueueName *string `locationName:"jobQueueName" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateJobQueueOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJobQueueOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobQueueArn != nil {
		v := *s.JobQueueArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobQueueArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobQueueName != nil {
		v := *s.JobQueueName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobQueueName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateJobQueue = "CreateJobQueue"

// CreateJobQueueRequest returns a request value for making API operation for
// AWS Batch.
//
// Creates an AWS Batch job queue. When you create a job queue, you associate
// one or more compute environments to the queue and assign an order of preference
// for the compute environments.
//
// You also set a priority to the job queue that determines the order in which
// the AWS Batch scheduler places jobs onto its associated compute environments.
// For example, if a compute environment is associated with more than one job
// queue, the job queue with a higher priority is given preference for scheduling
// jobs to that compute environment.
//
//    // Example sending a request using CreateJobQueueRequest.
//    req := client.CreateJobQueueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/CreateJobQueue
func (c *Client) CreateJobQueueRequest(input *CreateJobQueueInput) CreateJobQueueRequest {
	op := &aws.Operation{
		Name:       opCreateJobQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/createjobqueue",
	}

	if input == nil {
		input = &CreateJobQueueInput{}
	}

	req := c.newRequest(op, input, &CreateJobQueueOutput{})
	return CreateJobQueueRequest{Request: req, Input: input, Copy: c.CreateJobQueueRequest}
}

// CreateJobQueueRequest is the request type for the
// CreateJobQueue API operation.
type CreateJobQueueRequest struct {
	*aws.Request
	Input *CreateJobQueueInput
	Copy  func(*CreateJobQueueInput) CreateJobQueueRequest
}

// Send marshals and sends the CreateJobQueue API request.
func (r CreateJobQueueRequest) Send(ctx context.Context) (*CreateJobQueueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateJobQueueResponse{
		CreateJobQueueOutput: r.Request.Data.(*CreateJobQueueOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateJobQueueResponse is the response type for the
// CreateJobQueue API operation.
type CreateJobQueueResponse struct {
	*CreateJobQueueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateJobQueue request.
func (r *CreateJobQueueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}