# GuardDutyAlertSlack

Guarddutyの結果をEventBridgeとLambdaを使ってSlackに通知する

- コンパイルとデプロイ
```bash
GOARCH=amd64 GOOS=linux go build "-ldflags=-s -w" ./main.go
sls deploy --slacktoken {$SLACK_TOKEN} --channelid  ${CHANNEL_ID} --threshold ${THRESHOLD}
```
