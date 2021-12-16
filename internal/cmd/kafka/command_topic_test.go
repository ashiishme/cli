package kafka

import (
	"context"
	"fmt"
	"testing"

	"github.com/c-bata/go-prompt"
	schedv1 "github.com/confluentinc/cc-structs/kafka/scheduler/v1"
	"github.com/confluentinc/ccloud-sdk-go-v1"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	ccsdkmock "github.com/confluentinc/ccloud-sdk-go-v1/mock"

	v1 "github.com/confluentinc/cli/internal/pkg/config/v1"
	cliMock "github.com/confluentinc/cli/mock"
)

const (
	topicName  = "topic"
	isInternal = false
)

type KafkaTopicTestSuite struct {
	suite.Suite
}

func (suite *KafkaTopicTestSuite) newCmd(conf *v1.Config) *kafkaTopicCommand {
	client := &ccloud.Client{
		Kafka: &ccsdkmock.Kafka{
			ListTopicsFunc: func(_ context.Context, cluster *schedv1.KafkaCluster) ([]*schedv1.TopicDescription, error) {
				return []*schedv1.TopicDescription{
					{
						Name:     topicName,
						Internal: isInternal,
					},
				}, nil
			},
		},
	}
	prerunner := cliMock.NewPreRunnerMock(client, nil, nil, conf)
	cmd := newTopicCommand(conf, prerunner, nil, "id")
	return cmd
}

func (suite *KafkaTopicTestSuite) TestServerComplete() {
	req := suite.Require()
	type fields struct {
		Command *kafkaTopicCommand
	}
	tests := []struct {
		name   string
		fields fields
		want   []prompt.Suggest
	}{
		{
			name: "suggest for authenticated user",
			fields: fields{
				Command: suite.newCmd(v1.AuthenticatedCloudConfigMock()),
			},
			want: []prompt.Suggest{
				{
					Text:        topicName,
					Description: "",
				},
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			_ = tt.fields.Command.PersistentPreRunE(tt.fields.Command.Command, []string{})
			got := tt.fields.Command.ServerComplete()
			fmt.Println(&got)
			req.Equal(tt.want, got)
		})
	}
}

func (suite *KafkaTopicTestSuite) TestServerCompletableChildren() {
	req := require.New(suite.T())
	cmd := suite.newCmd(v1.AuthenticatedCloudConfigMock())
	completableChildren := cmd.ServerCompletableChildren()
	expectedChildren := []string{"topic describe", "topic update", "topic delete"}
	req.Len(completableChildren, len(expectedChildren))
	for i, expectedChild := range expectedChildren {
		req.Contains(completableChildren[i].CommandPath(), expectedChild)
	}
}

func TestKafkaTopicTestSuite(t *testing.T) {
	suite.Run(t, new(KafkaTopicTestSuite))
}
