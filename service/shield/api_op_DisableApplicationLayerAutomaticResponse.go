// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disable the Shield Advanced automatic application layer DDoS mitigation feature
// for the resource. This stops Shield Advanced from creating, verifying, and
// applying WAF rules for attacks that it detects for the resource.
func (c *Client) DisableApplicationLayerAutomaticResponse(ctx context.Context, params *DisableApplicationLayerAutomaticResponseInput, optFns ...func(*Options)) (*DisableApplicationLayerAutomaticResponseOutput, error) {
	if params == nil {
		params = &DisableApplicationLayerAutomaticResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableApplicationLayerAutomaticResponse", params, optFns, c.addOperationDisableApplicationLayerAutomaticResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableApplicationLayerAutomaticResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableApplicationLayerAutomaticResponseInput struct {

	// The ARN (Amazon Resource Name) of the resource.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type DisableApplicationLayerAutomaticResponseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableApplicationLayerAutomaticResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisableApplicationLayerAutomaticResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableApplicationLayerAutomaticResponse{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDisableApplicationLayerAutomaticResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableApplicationLayerAutomaticResponse(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDisableApplicationLayerAutomaticResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "DisableApplicationLayerAutomaticResponse",
	}
}