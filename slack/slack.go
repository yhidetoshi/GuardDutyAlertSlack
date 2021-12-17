package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

var (
	USERNAME    = "GuardDutyAlert"
	SLACK_TOKEN = os.Getenv("SLACK_TOKEN")
	CHANNEL_ID  = os.Getenv("CHANNEL_ID")
)

//func PostSlack(){
func PostSlack(slackColor string, title string, accountAliasName string, severity string, resource string, reason string, description string) {
	api := slack.New(SLACK_TOKEN)
	attachment := slack.Attachment{
		Title: USERNAME,
		Color: slackColor,

		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Title",
				Value: title,
				Short: false,
			}, slack.AttachmentField{
				Title: "Account",
				Value: accountAliasName,
				Short: false,
			}, slack.AttachmentField{
				Title: "Severity",
				Value: severity,
				Short: false,
			}, slack.AttachmentField{
				Title: "Affected Resource",
				Value: resource,
				Short: false,
			}, slack.AttachmentField{
				Title: "Type",
				Value: reason,
				Short: false,
			}, slack.AttachmentField{
				Title: "Description",
				Value: "```" + description + "```",
				Short: false,
			},
		},
	}

	_, _, err := api.PostMessage(
		CHANNEL_ID,
		//slack.MsgOptionText("Title", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
