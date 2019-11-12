// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateFunctionDefinitionInput struct {
	_ struct{} `type:"structure"`

	// FunctionDefinitionId is a required field
	FunctionDefinitionId *string `location:"uri" locationName:"FunctionDefinitionId" type:"string" required:"true"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateFunctionDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFunctionDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFunctionDefinitionInput"}

	if s.FunctionDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFunctionDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionDefinitionId != nil {
		v := *s.FunctionDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateFunctionDefinitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateFunctionDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFunctionDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateFunctionDefinition = "UpdateFunctionDefinition"

// UpdateFunctionDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Updates a Lambda function definition.
//
//    // Example sending a request using UpdateFunctionDefinitionRequest.
//    req := client.UpdateFunctionDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateFunctionDefinition
func (c *Client) UpdateFunctionDefinitionRequest(input *UpdateFunctionDefinitionInput) UpdateFunctionDefinitionRequest {
	op := &aws.Operation{
		Name:       opUpdateFunctionDefinition,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/definition/functions/{FunctionDefinitionId}",
	}

	if input == nil {
		input = &UpdateFunctionDefinitionInput{}
	}

	req := c.newRequest(op, input, &UpdateFunctionDefinitionOutput{})
	return UpdateFunctionDefinitionRequest{Request: req, Input: input, Copy: c.UpdateFunctionDefinitionRequest}
}

// UpdateFunctionDefinitionRequest is the request type for the
// UpdateFunctionDefinition API operation.
type UpdateFunctionDefinitionRequest struct {
	*aws.Request
	Input *UpdateFunctionDefinitionInput
	Copy  func(*UpdateFunctionDefinitionInput) UpdateFunctionDefinitionRequest
}

// Send marshals and sends the UpdateFunctionDefinition API request.
func (r UpdateFunctionDefinitionRequest) Send(ctx context.Context) (*UpdateFunctionDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFunctionDefinitionResponse{
		UpdateFunctionDefinitionOutput: r.Request.Data.(*UpdateFunctionDefinitionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFunctionDefinitionResponse is the response type for the
// UpdateFunctionDefinition API operation.
type UpdateFunctionDefinitionResponse struct {
	*UpdateFunctionDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFunctionDefinition request.
func (r *UpdateFunctionDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}