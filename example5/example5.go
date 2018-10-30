package example5

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type Profile struct {
	ID   int32
	Name string
}

// QueueURL: https://hogehoge.com
// 上記キューに対して引数で渡したinterface{}をMessageBodyとして書き込む
func SendSQSMessage(client sqsiface.SQSAPI, p *Profile) (*sqs.SendMessageOutput, error) {
	jsonBytes, err := json.Marshal(*p)
	if err != nil {
		return nil, err
	}

	messageInput := &sqs.SendMessageInput{
		MessageBody:  aws.String(string(jsonBytes)),
		QueueUrl:     aws.String("https://hogehoge.com"),
		DelaySeconds: aws.Int64(1),
	}

	_ = messageInput
	return nil, nil
	//return client.SendMessage(messageInput)
}
