// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a get commit operation.
type GetCommitInput struct {
	_ struct{} `type:"structure"`

	// The commit ID. Commit IDs are the full SHA of the commit.
	//
	// CommitId is a required field
	CommitId *string `locationName:"commitId" type:"string" required:"true"`

	// The name of the repository to which the commit was made.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCommitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCommitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCommitInput"}

	if s.CommitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommitId"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a get commit operation.
type GetCommitOutput struct {
	_ struct{} `type:"structure"`

	// A commit data type object that contains information about the specified commit.
	//
	// Commit is a required field
	Commit *Commit `locationName:"commit" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetCommitOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCommit = "GetCommit"

// GetCommitRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about a commit, including commit message and committer
// information.
//
//    // Example sending a request using GetCommitRequest.
//    req := client.GetCommitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetCommit
func (c *Client) GetCommitRequest(input *GetCommitInput) GetCommitRequest {
	op := &aws.Operation{
		Name:       opGetCommit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCommitInput{}
	}

	req := c.newRequest(op, input, &GetCommitOutput{})
	return GetCommitRequest{Request: req, Input: input, Copy: c.GetCommitRequest}
}

// GetCommitRequest is the request type for the
// GetCommit API operation.
type GetCommitRequest struct {
	*aws.Request
	Input *GetCommitInput
	Copy  func(*GetCommitInput) GetCommitRequest
}

// Send marshals and sends the GetCommit API request.
func (r GetCommitRequest) Send(ctx context.Context) (*GetCommitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCommitResponse{
		GetCommitOutput: r.Request.Data.(*GetCommitOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCommitResponse is the response type for the
// GetCommit API operation.
type GetCommitResponse struct {
	*GetCommitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCommit request.
func (r *GetCommitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}