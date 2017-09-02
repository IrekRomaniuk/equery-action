package equery


import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"errors"
	"log"
	"os"
)
// Syslog struct
type Syslog struct {
	Raw,Domain,ReceiveTime,SerialNum,Type,Subtype,ConfigVersion,GenerateTime,SourceIP,DestinationIP,
	NATSourceIP,NATDestinationIP,Rule,SourceUser,DestinationUser,Application,VirtualSystem,SourceZone,DestinationZone,
	InboundInterface,OutboundInterface,LogAction,TimeLogged,SessionID,RepeatCount,SourcePort,DestinationPort,NATSourcePort,
	NATDestinationPort,Flags,Protocol,Action,URL,ThreatContentName,Category,Severity,Direction,Seqno,ActionFlags,
	SourceLocation,DestinationLocation,Cpaddingth,ContentType,Pcapid,Filedigest,Cloud,Urlidx,Useragent,Filetype,Xff,
	Referer,Sender,Subject,Recipient,Reportid string
	count uint64  
  }
// Func Agg
func Agg(url, index, field, name string) {
	// Create a client
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	timeRangeFilter := elastic.NewRangeQuery("@timestamp").Gte("now-10m").Lte("now")
	query := elastic.NewBoolQuery().
	Must(timeRangeFilter).
	Must(elastic.NewMatchAllQuery())
	
	agg := elastic.NewTermsAggregation().Field(field).Size(10)
	search := client.Search().Index(index).Query(query)
	search = search.Aggregation(name, agg)
	sr, err := search.Do(context.Background())
    if err != nil {
        log.Fatal("error in aggregation Do:", err)
    }
 
    if agg, found := sr.Aggregations.Terms(name); found {
        for _, bucket := range agg.Buckets {
      log.Println("key:", bucket.Key, ", count:", bucket.DocCount)
        }
    }
}

// Func Query
func Query(url, index, field, value string) (int64, error) {
	// See README for query example
	// Create a client
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	timeRangeFilter := elastic.NewRangeQuery("@timestamp").Gte("now-10m").Lte("now")
	query := elastic.NewBoolQuery().
	Must(timeRangeFilter).
	Must(elastic.NewMatchAllQuery())
	
	searchResult, err := client.Search().
	Index(index).   // search in index
	Query(query).   // specify the query
	Sort(field, true). // sort by field, ascending
	From(0).Size(10).   // take documents 0-9
	Pretty(true).       // pretty print request and response JSON
	Do(context.Background()) 
	return searchResult.Hits.TotalHits, err
}

// Search elastic
func Search(url, index, field, value string) (int64, error) {
	// Create a client
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	// Search with a term query	
	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		return 0, err
	}
	if exists {
		termQuery := elastic.NewTermQuery(field, value)		
		searchResult, err := client.Search().
		Index(index).   // search in index
		Query(termQuery).   // specify the query
		Sort(field, true). // sort by field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(context.Background())             // execute		
		return searchResult.Hits.TotalHits, err
	} 
	err = errors.New("Index does not exists")
	return 0, err
	
}

// Ping db
func Ping (url string) {
	// Obtain a client
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)
	client, err := elastic.NewSimpleClient(elastic.SetErrorLog(errorlog))
	if err != nil {
		log.Fatal(err)
	}
	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(url).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

// Version of db
func Version (url string) {
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	if err != nil {
		log.Fatal(err)
	}
	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Elasticsearch version %s\n", esversion)
}