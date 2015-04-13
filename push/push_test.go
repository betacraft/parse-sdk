package push

import (
	"testing"
)

const (
	ObjectId   = "2nbM0lyLAb"
	AppId      = "DoymS4OxgNxvSAosUSUxshEHLI9tLec0XOZuGNm3"
	RestApiKey = "m4i58Jq8PfJpu9skFDHe7PXQduACa91c1FLL8uZf"
)

func TestSendPushNotification(t *testing.T) {
	sdk := Init(AppId, RestApiKey)
	err := sdk.SendToObjectId(ObjectId, map[string]interface{}{"data": "test"})
	if err != nil {
		t.Error(err)
	}
}
