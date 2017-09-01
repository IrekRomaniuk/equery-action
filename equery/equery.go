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

// Search elastic
func Search(url, index string) (int64, error) {
	// Create a context
	ctx := context.Background()
	// Create a client
	client, err := elastic.NewClient(elastic.SetURL(url))
	// Search with a term query
	/*"filter": {
        "range": {
          "whatever_timestampfield": {
            "gte": "now-10m",
            "lte": "now"
          }
		}*/
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		return 0, err
	}
	if exists {
		termQuery := elastic.NewTermQuery("user", "olivere")
		searchResult, err := client.Search().
		Index(index).   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute		
		return searchResult.Hits.TotalHits, err
	} 
	
	err = errors.New("Index does not exists")
	return 0, err
	
}

// Ping db
func Ping (url string) {
	// Obtain a client
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetErrorLog(errorlog))
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
	client, err := elastic.NewClient(elastic.SetURL(url))
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