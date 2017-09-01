package equery

import (
	"testing"
)

func TestPingandVersion(t *testing.T) {
	Version("http://10.254.253.100:9200")
	Ping("http://10.254.253.100:9200")
}