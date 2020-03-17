// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListFlowDefinitionsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only flow definitions with a creation time greater
	// than or equal to the specified timestamp.
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only flow definitions that were created before the
	// specified timestamp.
	CreationTimeBefore *time.Time `type:"timestamp"`

	// The total number of items to return. If the total number of available items
	// is more than the value specified in MaxResults, then a NextToken will be
	// provided in the output that you can use to resume pagination.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to resume pagination.
	NextToken *string `type:"string"`

	// An optional value that specifies whether you want the results sorted in Ascending
	// or Descending order.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListFlowDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFlowDefinitionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFlowDefinitionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListFlowDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects describing the flow definitions.
	//
	// FlowDefinitionSummaries is a required field
	FlowDefinitionSummaries []FlowDefinitionSummary `type:"list" required:"true"`

	// A token to resume pagination.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListFlowDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListFlowDefinitions = "ListFlowDefinitions"

// ListFlowDefinitionsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns information about the flow definitions in your account.
//
//    // Example sending a request using ListFlowDefinitionsRequest.
//    req := client.ListFlowDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListFlowDefinitions
func (c *Client) ListFlowDefinitionsRequest(input *ListFlowDefinitionsInput) ListFlowDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListFlowDefinitions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFlowDefinitionsInput{}
	}

	req := c.newRequest(op, input, &ListFlowDefinitionsOutput{})
	return ListFlowDefinitionsRequest{Request: req, Input: input, Copy: c.ListFlowDefinitionsRequest}
}

// ListFlowDefinitionsRequest is the request type for the
// ListFlowDefinitions API operation.
type ListFlowDefinitionsRequest struct {
	*aws.Request
	Input *ListFlowDefinitionsInput
	Copy  func(*ListFlowDefinitionsInput) ListFlowDefinitionsRequest
}

// Send marshals and sends the ListFlowDefinitions API request.
func (r ListFlowDefinitionsRequest) Send(ctx context.Context) (*ListFlowDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFlowDefinitionsResponse{
		ListFlowDefinitionsOutput: r.Request.Data.(*ListFlowDefinitionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFlowDefinitionsRequestPaginator returns a paginator for ListFlowDefinitions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFlowDefinitionsRequest(input)
//   p := sagemaker.NewListFlowDefinitionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFlowDefinitionsPaginator(req ListFlowDefinitionsRequest) ListFlowDefinitionsPaginator {
	return ListFlowDefinitionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFlowDefinitionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListFlowDefinitionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFlowDefinitionsPaginator struct {
	aws.Pager
}

func (p *ListFlowDefinitionsPaginator) CurrentPage() *ListFlowDefinitionsOutput {
	return p.Pager.CurrentPage().(*ListFlowDefinitionsOutput)
}

// ListFlowDefinitionsResponse is the response type for the
// ListFlowDefinitions API operation.
type ListFlowDefinitionsResponse struct {
	*ListFlowDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFlowDefinitions request.
func (r *ListFlowDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}