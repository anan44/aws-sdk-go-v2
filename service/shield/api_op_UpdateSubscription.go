// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// When you initally create a subscription, AutoRenew is set to ENABLED. If
	// ENABLED, the subscription will be automatically renewed at the end of the
	// existing subscription period. You can change this by submitting an UpdateSubscription
	// request. If the UpdateSubscription request does not included a value for
	// AutoRenew, the existing value for AutoRenew remains unchanged.
	AutoRenew AutoRenew `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

type UpdateSubscriptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSubscription = "UpdateSubscription"

// UpdateSubscriptionRequest returns a request value for making API operation for
// AWS Shield.
//
// Updates the details of an existing subscription. Only enter values for parameters
// you want to change. Empty parameters are not updated.
//
//    // Example sending a request using UpdateSubscriptionRequest.
//    req := client.UpdateSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/UpdateSubscription
func (c *Client) UpdateSubscriptionRequest(input *UpdateSubscriptionInput) UpdateSubscriptionRequest {
	op := &aws.Operation{
		Name:       opUpdateSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSubscriptionInput{}
	}

	req := c.newRequest(op, input, &UpdateSubscriptionOutput{})
	return UpdateSubscriptionRequest{Request: req, Input: input, Copy: c.UpdateSubscriptionRequest}
}

// UpdateSubscriptionRequest is the request type for the
// UpdateSubscription API operation.
type UpdateSubscriptionRequest struct {
	*aws.Request
	Input *UpdateSubscriptionInput
	Copy  func(*UpdateSubscriptionInput) UpdateSubscriptionRequest
}

// Send marshals and sends the UpdateSubscription API request.
func (r UpdateSubscriptionRequest) Send(ctx context.Context) (*UpdateSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSubscriptionResponse{
		UpdateSubscriptionOutput: r.Request.Data.(*UpdateSubscriptionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSubscriptionResponse is the response type for the
// UpdateSubscription API operation.
type UpdateSubscriptionResponse struct {
	*UpdateSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSubscription request.
func (r *UpdateSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}