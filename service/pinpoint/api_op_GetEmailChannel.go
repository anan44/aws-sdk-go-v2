// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetEmailChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetEmailChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEmailChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEmailChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEmailChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetEmailChannelOutput struct {
	_ struct{} `type:"structure" payload:"EmailChannelResponse"`

	// Provides information about the status and settings of the email channel for
	// an application.
	//
	// EmailChannelResponse is a required field
	EmailChannelResponse *EmailChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetEmailChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEmailChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EmailChannelResponse != nil {
		v := s.EmailChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "EmailChannelResponse", v, metadata)
	}
	return nil
}

const opGetEmailChannel = "GetEmailChannel"

// GetEmailChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of the email channel
// for an application.
//
//    // Example sending a request using GetEmailChannelRequest.
//    req := client.GetEmailChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetEmailChannel
func (c *Client) GetEmailChannelRequest(input *GetEmailChannelInput) GetEmailChannelRequest {
	op := &aws.Operation{
		Name:       opGetEmailChannel,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/channels/email",
	}

	if input == nil {
		input = &GetEmailChannelInput{}
	}

	req := c.newRequest(op, input, &GetEmailChannelOutput{})
	return GetEmailChannelRequest{Request: req, Input: input, Copy: c.GetEmailChannelRequest}
}

// GetEmailChannelRequest is the request type for the
// GetEmailChannel API operation.
type GetEmailChannelRequest struct {
	*aws.Request
	Input *GetEmailChannelInput
	Copy  func(*GetEmailChannelInput) GetEmailChannelRequest
}

// Send marshals and sends the GetEmailChannel API request.
func (r GetEmailChannelRequest) Send(ctx context.Context) (*GetEmailChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEmailChannelResponse{
		GetEmailChannelOutput: r.Request.Data.(*GetEmailChannelOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEmailChannelResponse is the response type for the
// GetEmailChannel API operation.
type GetEmailChannelResponse struct {
	*GetEmailChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEmailChannel request.
func (r *GetEmailChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}