package slack

import (
	"github.com/slack-go/slack"
)

type (
	slackClient struct {
		api              *slack.Client
		channelID, token string
	}

	Option func(client *slackClient) error
)

var client *slackClient

// New
func New(options ...Option) *slackClient {
	client = &slackClient{}

	for _, opt := range options {
		opt(client)
	}

	return client
}

// Client
func Client() *slackClient {
	return client
}

// WithToken
func WithToken(token string) Option {
	return func(client *slackClient) error {
		client.token = token
		return nil
	}
}

// WithChannelID
func WithChannelID(channel string) Option {
	return func(client *slackClient) error {
		client.channelID = channel
		return nil
	}
}

// WithSlackAPIInit
func WithSlackAPIInit() Option {
	return func(client *slackClient) error {
		client.api = slack.New(client.token)
		return nil
	}
}

// Channel
// This should be value receiver type
func (client slackClient) Channel(channel string) *slackClient {
	client.channelID = channel
	return &client
}

// PushMessage
func (client *slackClient) PushMessage(data string) (string, string, error) {
	channelId, timestamp, err := client.api.PostMessage(
		client.channelID,
		slack.MsgOptionText(data, false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		return "", "", err
	}

	return channelId, timestamp, err
}
