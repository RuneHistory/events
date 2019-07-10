package events

import (
	"testing"
)

func TestNewAccountEvent_IsPublishable(t *testing.T) {
	t.Parallel()
	if !eventIsPublishable(&NewAccountEvent{}) {
		t.Fail()
	}
}

func TestNewAccountEvent_IsSubscribable(t *testing.T) {
	t.Parallel()
	if !eventIsSubscribable(&NewAccountEvent{}) {
		t.Fail()
	}
}
