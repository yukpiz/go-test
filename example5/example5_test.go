package example5

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/golang/mock/gomock"

	"./mock"
)

func TestSendSQSMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sqsMock := mock.NewMockSQSAPI(ctrl)

	profile := &Profile{
		ID:   1,
		Name: "yukpiz",
	}
	jsonBytes, _ := json.Marshal(*profile)
	messageInput := &sqs.SendMessageInput{
		MessageBody:  aws.String(string(jsonBytes)),
		QueueUrl:     aws.String("https://hogehoge.com"),
		DelaySeconds: aws.Int64(1),
	}

	//SendMessageがエラーを返すモックを定義
	sqsMock.EXPECT().SendMessage(messageInput).Return(nil, errors.New("Failed Mock Error!!"))

	ret, err := SendSQSMessage(sqsMock, profile)
	fmt.Printf("Result: %+v\n", ret)
	fmt.Printf("Error: %+v\n", err)
}
