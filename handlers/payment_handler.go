package handlers

import (
	"context"
	"escort-book-payment-release/models"
	"escort-book-payment-release/repositories"
	"escort-book-payment-release/services"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/lambda"
	_ "github.com/joho/godotenv/autoload"
)

type PaymentHandler struct {
	ServiceRepository  *repositories.ServiceRepository
	UserRepository     *repositories.UserRepository
	PaymentRepository  *repositories.PaymentRepository
	FirebaseService    *services.FirebaseService
	EventBridgeService *services.EventBridgeService
}

func (h *PaymentHandler) HandleRequest(ctx context.Context, event events.CloudWatchEvent) {
	jobName := strings.Split(event.Resources[0], "/")[1]
	id := strings.Split(jobName, "-")
	userType, serviceId := id[0], id[1]

	service := h.ServiceRepository.GetById(ctx, serviceId)

	if service.Status != "started" {
		return
	}

	if userType == "Customer" {
		user := h.UserRepository.GetById(ctx, service.CustomerId)
		log.Println("WE SEND A NOTIFICATION TO CUSTOMER: ", user.Email)
	} else {
		newPayment := models.Payment{
			EscortId:    service.EscortId,
			CustomerId:  service.CustomerId,
			ServiceId:   service.Id,
			LogRequest:  `{"request": "dummy request"}`,
			LogResponse: `{"response": "dummy response"}`,
		}
		newPayment.SetDefaultValues()
		h.PaymentRepository.Create(ctx, &newPayment)

		newService := map[string]string{"status": "completed"}
		h.ServiceRepository.UpdateById(ctx, serviceId, &newService)

		user := h.UserRepository.GetById(ctx, service.EscortId)
		log.Println("WE SEND A NOTIFICATION TO ESCORT: ", user.Email)
	}

	removeTargetInput := eventbridge.RemoveTargetsInput{
		Ids: []*string{
			aws.String(os.Getenv("LAMBDA")),
		},
		Rule:  aws.String(jobName),
		Force: aws.Bool(true),
	}

	_, err := h.EventBridgeService.RemoveTargetRuleLambda(ctx, &removeTargetInput)

	if err != nil {
		log.Println("ERROR WHEN REMOVE TARGET: ", err.Error())
		return
	}

	deleteRuleInput := eventbridge.DeleteRuleInput{
		Name:  aws.String(jobName),
		Force: aws.Bool(true),
	}

	_, err = h.EventBridgeService.RemoveRule(ctx, &deleteRuleInput)

	if err != nil {
		log.Println("ERROR WHEN DELETING A RULE: ", err.Error())
		return
	}

	removePermisionInput := lambda.RemovePermissionInput{
		FunctionName: aws.String(os.Getenv("LAMBDA")),
		StatementId:  aws.String(jobName),
	}

	_, err = h.EventBridgeService.RemoveLambdaPermissions(ctx, &removePermisionInput)

	if err != nil {
		log.Println("ERROR WHEN REMOVE PERMISSION FOR LAMBDA: ", err.Error())
	}
}
