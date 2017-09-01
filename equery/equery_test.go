package equery

import (
	"testing"
	"fmt"
)

func TestPingorVersion(t *testing.T) {
	//Version("http://10.254.253.100:9200")
	Ping("http://10.254.253.100:9200")
}

func TestSearch(t *testing.T) {
	results, _ := Search("http://10.254.253.100:9200", "logstash-2017.09.01", "SourceIP", "172.172.172.172")
	fmt.Printf("Found a total of %d hits\n", results)
	results, _ = Query("http://10.254.253.100:9200", "logstash-2017.09.01", "SourceIP", "172.172.172.172")
	fmt.Printf("Found a total of %d hits\n", results)
	
}
