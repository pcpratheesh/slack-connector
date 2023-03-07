package slack_test

import (
	"testing"

	"github.com/pcpratheesh/slack-connector"
	"github.com/stretchr/testify/require"
)

func TestSlack(t *testing.T) {
	client := slack.New(
		slack.WithToken("token-id"),
		slack.WithChannelID("channel-id"),
		slack.WithSlackAPIInit(),
	)
	_ = client
	_, _, err := slack.Client().PushMessage("pushed first")
	require.Nil(t, err)
}
func TestSlackWithNewChannel(t *testing.T) {
	client := slack.New(
		slack.WithToken("token-id"),
		slack.WithChannelID("channel-id"),
		slack.WithSlackAPIInit(),
	)
	_ = client
	_, _, err := slack.Client().Channel("channel-id").PushMessage("pushed second")
	require.Nil(t, err)
}

func TestSlackForChannel(t *testing.T) {
	client := slack.New(
		slack.WithToken("token-id"),
		slack.WithChannelID("channel-id"),
		slack.WithSlackAPIInit(),
	)
	_ = client
	_, _, err := slack.Client().PushMessage("pushed third")
	require.Nil(t, err)
}

func TestSlackWithMultipleChannels(t *testing.T) {
	client := slack.New(
		slack.WithToken("token-id"),
		slack.WithChannelID("channel-id"),
		slack.WithSlackAPIInit(),
	)
	_ = client
	_, _, err := slack.Client().PushMessage("pushed A")
	require.Nil(t, err)

	_, _, err = slack.Client().Channel("channel-id").PushMessage("pushed B")
	require.Nil(t, err)

	_, _, err = slack.Client().PushMessage("pushed C")
	require.Nil(t, err)
}
