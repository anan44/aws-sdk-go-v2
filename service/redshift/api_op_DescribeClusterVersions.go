// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeClusterVersionsInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific cluster parameter group family to return details for.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 alphanumeric characters
	//
	//    * First character must be a letter
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens
	ClusterParameterGroupFamily *string `type:"string"`

	// The specific cluster version to return.
	//
	// Example: 1.0
	ClusterVersion *string `type:"string"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterVersions request exceed
	// the value specified in MaxRecords, AWS returns a value in the Marker field
	// of the response. You can retrieve the next set of response records by providing
	// the returned marker value in the Marker parameter and retrying the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeClusterVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output from the DescribeClusterVersions action.
type DescribeClusterVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of Version elements.
	ClusterVersions []ClusterVersion `locationNameList:"ClusterVersion" type:"list"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeClusterVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClusterVersions = "DescribeClusterVersions"

// DescribeClusterVersionsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns descriptions of the available Amazon Redshift cluster versions. You
// can call this operation even before creating any clusters to learn more about
// the Amazon Redshift versions. For more information about managing clusters,
// go to Amazon Redshift Clusters (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using DescribeClusterVersionsRequest.
//    req := client.DescribeClusterVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeClusterVersions
func (c *Client) DescribeClusterVersionsRequest(input *DescribeClusterVersionsInput) DescribeClusterVersionsRequest {
	op := &aws.Operation{
		Name:       opDescribeClusterVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeClusterVersionsInput{}
	}

	req := c.newRequest(op, input, &DescribeClusterVersionsOutput{})
	return DescribeClusterVersionsRequest{Request: req, Input: input, Copy: c.DescribeClusterVersionsRequest}
}

// DescribeClusterVersionsRequest is the request type for the
// DescribeClusterVersions API operation.
type DescribeClusterVersionsRequest struct {
	*aws.Request
	Input *DescribeClusterVersionsInput
	Copy  func(*DescribeClusterVersionsInput) DescribeClusterVersionsRequest
}

// Send marshals and sends the DescribeClusterVersions API request.
func (r DescribeClusterVersionsRequest) Send(ctx context.Context) (*DescribeClusterVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClusterVersionsResponse{
		DescribeClusterVersionsOutput: r.Request.Data.(*DescribeClusterVersionsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeClusterVersionsRequestPaginator returns a paginator for DescribeClusterVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeClusterVersionsRequest(input)
//   p := redshift.NewDescribeClusterVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeClusterVersionsPaginator(req DescribeClusterVersionsRequest) DescribeClusterVersionsPaginator {
	return DescribeClusterVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeClusterVersionsInput
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

// DescribeClusterVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeClusterVersionsPaginator struct {
	aws.Pager
}

func (p *DescribeClusterVersionsPaginator) CurrentPage() *DescribeClusterVersionsOutput {
	return p.Pager.CurrentPage().(*DescribeClusterVersionsOutput)
}

// DescribeClusterVersionsResponse is the response type for the
// DescribeClusterVersions API operation.
type DescribeClusterVersionsResponse struct {
	*DescribeClusterVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClusterVersions request.
func (r *DescribeClusterVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}