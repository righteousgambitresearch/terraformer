// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListAvailableSolutionStacksInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListAvailableSolutionStacksInput) String() string {
	return awsutil.Prettify(s)
}

// A list of available AWS Elastic Beanstalk solution stacks.
type ListAvailableSolutionStacksOutput struct {
	_ struct{} `type:"structure"`

	// A list of available solution stacks and their SolutionStackDescription.
	SolutionStackDetails []SolutionStackDescription `type:"list"`

	// A list of available solution stacks.
	SolutionStacks []string `type:"list"`
}

// String returns the string representation
func (s ListAvailableSolutionStacksOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAvailableSolutionStacks = "ListAvailableSolutionStacks"

// ListAvailableSolutionStacksRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Returns a list of the available solution stack names, with the public version
// first and then in reverse chronological order.
//
//    // Example sending a request using ListAvailableSolutionStacksRequest.
//    req := client.ListAvailableSolutionStacksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ListAvailableSolutionStacks
func (c *Client) ListAvailableSolutionStacksRequest(input *ListAvailableSolutionStacksInput) ListAvailableSolutionStacksRequest {
	op := &aws.Operation{
		Name:       opListAvailableSolutionStacks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAvailableSolutionStacksInput{}
	}

	req := c.newRequest(op, input, &ListAvailableSolutionStacksOutput{})
	return ListAvailableSolutionStacksRequest{Request: req, Input: input, Copy: c.ListAvailableSolutionStacksRequest}
}

// ListAvailableSolutionStacksRequest is the request type for the
// ListAvailableSolutionStacks API operation.
type ListAvailableSolutionStacksRequest struct {
	*aws.Request
	Input *ListAvailableSolutionStacksInput
	Copy  func(*ListAvailableSolutionStacksInput) ListAvailableSolutionStacksRequest
}

// Send marshals and sends the ListAvailableSolutionStacks API request.
func (r ListAvailableSolutionStacksRequest) Send(ctx context.Context) (*ListAvailableSolutionStacksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAvailableSolutionStacksResponse{
		ListAvailableSolutionStacksOutput: r.Request.Data.(*ListAvailableSolutionStacksOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAvailableSolutionStacksResponse is the response type for the
// ListAvailableSolutionStacks API operation.
type ListAvailableSolutionStacksResponse struct {
	*ListAvailableSolutionStacksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAvailableSolutionStacks request.
func (r *ListAvailableSolutionStacksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}