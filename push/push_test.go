package push

import (
	"testing"
)

const (
	ObjectId   = "-"
	AppId      = "-"
	RestApiKey = "-"
)

func TestSendPushNotification(t *testing.T) {
	sdk := Init(AppId, RestApiKey)
	err := sdk.SendAlertTo(ObjectId, "test message")
	if err != nil {
		t.Error(err)
	}
}
