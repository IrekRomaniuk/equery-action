POST logstash-2017.09.02/_search 
 {
   "query": {
    "bool": {
      "must": { "match_all": {} },
      "filter": {
        "range": {
          "@timestamp": {
            "gte": "now-5m",
            "lte": "now"
          }
        }
      }
    }
  },
   "aggs": {
     "src": {
       "terms": {
         "field": "SourceIP"
       }
     }
   },
   "size": 0
 }