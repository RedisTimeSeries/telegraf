package kube_lite

import (
	"crypto/tls"
	"testing"
	"time"
)

type mockHandler struct {
	responseMap map[string]interface{}
}

func toStrPtr(s string) *string {
	return &s
}

func toInt32Ptr(i int32) *int32 {
	return &i
}

func toInt64Ptr(i int64) *int64 {
	return &i
}

func toBoolPtr(b bool) *bool {
	return &b
}

func TestNewClient(t *testing.T) {
	_, err := newClient("https://localhost:443/", "default", "abc123", time.Second, &tls.Config{})
	if err != nil {
		t.Errorf("Failed to create new client - %s", err.Error())
	}
}