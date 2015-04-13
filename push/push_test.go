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
	err := sdk.SendToObjectId(ObjectId, map[string]interface{}{"data": "test"})
	if err != nil {
		t.Error(err)
	}
}
