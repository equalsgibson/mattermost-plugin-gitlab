package webhook

import "github.com/mattermost/mattermost-plugin-gitlab/server/subscription"

const (
	DeploymentsKey  = "deployments"
	ReleasesKey     = "releases"
	MockChannelID   = "mockChannelID"
	MockOrgFullName = "mockOrg/mockRepo"
)

func GetMockSubscriptions(feature string) []*subscription.Subscription {
	return []*subscription.Subscription{
		{ChannelID: MockChannelID, CreatorID: "1", Features: feature, Repository: MockOrgFullName},
	}
}