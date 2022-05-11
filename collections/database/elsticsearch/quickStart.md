



- Username

  `elastic`

- Password

  `ZFlM758vLCp5drTXVxKD9vFC`



![image-20220424151631473](/Users/user/playground/share/nrookie.github.io/collections/database/elsticsearch/image-20220424151631473.png)



ucloudcn123naqing







``` shell
GET logs-my_app-default/_search
{
  "runtime_mappings": {
    "source.ip": {
      "type": "ip",
      "script": """
        String sourceip=grok('%{IPORHOST:sourceip} .*').extract(doc[ "event.original" ].value)?.sourceip;
        if (sourceip != null) emit(sourceip);
      """
    }
  },
  "query": {
    "range": {
      "@timestamp": {
        "gte": "2099-05-05",
        "lt": "2099-05-08"
      }
    }
  },
  "fields": [
    "@timestamp",
    "source.ip"
  ],
  "_source": false,
  "sort": [
    {
      "@timestamp": "desc"
    }
  ]
}
```





``` shell
GET logs-my_app-default/_search
{
  "query": {
    "match_all": { }
  },
  "sort": [
    {
      "@timestamp": "desc"
    }
  ]
}
```





https://www.elastic.co/guide/en/elasticsearch/reference/7.17/grok.html







what is _score in elastic-search



The `_score` in Elasticsearch is a way of determining how relevant a match is to the query. The default scoring function used by Elasticsearch is actually the default built in to Lucene which is what Elasticsearch runs under the hood. Here's an article that describes scoring fairly well.





The 'took' attribute in the response object is the execution time in milliseconds







### get the nodes from kibana



``` shell
GET /_nodes
```





### query with curl



``` shell
curl -X GET "http://100.127.9.84:9200"   
```

