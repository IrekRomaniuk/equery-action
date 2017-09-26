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
	results, _ := Search("http://10.254.253.100:9200", "logstash-2017.09.26", "SourceIP", "4.4.4.4")
	fmt.Printf("Found a total of %d hits\n", results)
	results, _ = Query("http://10.254.253.100:9200", "logstash-2017.09.26", "SourceIP", "4.4.4.4")
	fmt.Printf("Found a total of %d hits in last 10m\n", results)
	Agg("http://10.254.253.100:9200", "logstash-2017.09.26", "SourceIP")
	
}
