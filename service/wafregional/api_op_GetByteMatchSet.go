// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

type GetByteMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The ByteMatchSetId of the ByteMatchSet that you want to get. ByteMatchSetId
	// is returned by CreateByteMatchSet and by ListByteMatchSets.
	//
	// ByteMatchSetId is a required field
	ByteMatchSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetByteMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetByteMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetByteMatchSetInput"}

	if s.ByteMatchSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ByteMatchSetId"))
	}
	if s.ByteMatchSetId != nil && len(*s.ByteMatchSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ByteMatchSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetByteMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the ByteMatchSet that you specified in the GetByteMatchSet
	// request. For more information, see the following topics:
	//
	//    * ByteMatchSet: Contains ByteMatchSetId, ByteMatchTuples, and Name
	//
	//    * ByteMatchTuples: Contains an array of ByteMatchTuple objects. Each ByteMatchTuple
	//    object contains FieldToMatch, PositionalConstraint, TargetString, and
	//    TextTransformation
	//
	//    * FieldToMatch: Contains Data and Type
	ByteMatchSet *waf.ByteMatchSet `type:"structure"`
}

// String returns the string representation
func (s GetByteMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetByteMatchSet = "GetByteMatchSet"

// GetByteMatchSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the ByteMatchSet specified by ByteMatchSetId.
//
//    // Example sending a request using GetByteMatchSetRequest.
//    req := client.GetByteMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetByteMatchSet
func (c *Client) GetByteMatchSetRequest(input *GetByteMatchSetInput) GetByteMatchSetRequest {
	op := &aws.Operation{
		Name:       opGetByteMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetByteMatchSetInput{}
	}

	req := c.newRequest(op, input, &GetByteMatchSetOutput{})
	return GetByteMatchSetRequest{Request: req, Input: input, Copy: c.GetByteMatchSetRequest}
}

// GetByteMatchSetRequest is the request type for the
// GetByteMatchSet API operation.
type GetByteMatchSetRequest struct {
	*aws.Request
	Input *GetByteMatchSetInput
	Copy  func(*GetByteMatchSetInput) GetByteMatchSetRequest
}

// Send marshals and sends the GetByteMatchSet API request.
func (r GetByteMatchSetRequest) Send(ctx context.Context) (*GetByteMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetByteMatchSetResponse{
		GetByteMatchSetOutput: r.Request.Data.(*GetByteMatchSetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetByteMatchSetResponse is the response type for the
// GetByteMatchSet API operation.
type GetByteMatchSetResponse struct {
	*GetByteMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetByteMatchSet request.
func (r *GetByteMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}