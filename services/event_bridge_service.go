package services

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type IEventBridgeService interface {
	RemoveRule(ctx context.Context, input *eventbridge.DeleteRuleInput) (*eventbridge.DeleteRuleOutput, error)
	RemoveLambdaPermissions(ctx context.Context, input *lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error)
	RemoveTargetRuleLambda(ctx context.Context, input *eventbridge.RemoveTargetsInput) (*eventbridge.RemoveTargetsOutput, error)
}

type EventBridgeService struct{}

var eventBridgeClient *eventbridge.EventBridge
var lambdaClient *lambda.Lambda

func getEventBridgeClient() *eventbridge.EventBridge {
	if eventBridgeClient == nil {
		lock.Lock()
		defer lock.Unlock()

		if eventBridgeClient == nil {
			ses, _ := session.NewSession(&aws.Config{
				Region:      aws.String(os.Getenv("REGION")),
				Credentials: credentials.NewStaticCredentials("na", "na", ""),
				Endpoint:    aws.String(os.Getenv("ENDPOINT")),
			})
			eventBridgeClient = eventbridge.New(ses)
		}
	}

	return eventBridgeClient
}

func getLambdaClient() *lambda.Lambda {
	if lambdaClient == nil {
		lock.Lock()
		defer lock.Unlock()

		if lambdaClient == nil {
			ses, _ := session.NewSession(&aws.Config{
				Region:      aws.String(os.Getenv("REGION")),
				Credentials: credentials.NewStaticCredentials("na", "na", ""),
				Endpoint:    aws.String(os.Getenv("ENDPOINT")),
			})
			lambdaClient = lambda.New(ses)
		}
	}

	return lambdaClient
}

func (*EventBridgeService) RemoveRule(
	ctx context.Context,
	input *eventbridge.DeleteRuleInput,
) (*eventbridge.DeleteRuleOutput, error) {
	client := getEventBridgeClient()
	deleteRuleOutput, err := client.DeleteRuleWithContext(ctx, input)

	if err != nil {
		return nil, err
	}

	return deleteRuleOutput, nil
}

func (*EventBridgeService) RemoveLambdaPermissions(
	ctx context.Context,
	input *lambda.RemovePermissionInput,
) (*lambda.RemovePermissionOutput, error) {
	client := getLambdaClient()
	removePermissionOutput, err := client.RemovePermissionWithContext(ctx, input)

	if err != nil {
		return nil, err
	}

	return removePermissionOutput, nil
}

func (*EventBridgeService) RemoveTargetRuleLambda(
	ctx context.Context,
	input *eventbridge.RemoveTargetsInput,
) (*eventbridge.RemoveTargetsOutput, error) {
	client := getEventBridgeClient()
	removveTargetOutput, err := client.RemoveTargetsWithContext(ctx, input)

	if err != nil {
		return nil, err
	}

	return removveTargetOutput, nil
}
