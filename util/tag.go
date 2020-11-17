package util

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"strings"
)

func CompareTags(t1, t2 secretsmanager.Tag) bool {
	return strings.Compare(*t1.Key, *t2.Key) == 0 && strings.Compare(*t1.Value, *t2.Value) == 0
}