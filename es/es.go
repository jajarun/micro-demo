package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	config := elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200",
		},
		//Username: "foo",
		//Password: "bar",
	}
	es, err := elasticsearch.NewClient(config)

	//es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(elasticsearch.Version)
	fmt.Printf("info:%s \n", es.Info)

	//{
	//	res, err := es.Index(
	//		"my-index-000001",
	//		strings.NewReader(`{
	//	  "date": "2015-10-01T01:30:00Z"
	//	}`),
	//		es.Index.WithDocumentID("2"),
	//		es.Index.WithRefresh("true"),
	//		es.Index.WithPretty(),
	//	)
	//	fmt.Println(res, err)
	//	if err != nil { // SKIP
	//		fmt.Printf("Error getting the response: %s \n", err) // SKIP
	//	} // SKIP
	//	//defer res.Body.Close() // SKIP
	//}

	//{
	//	res, err := es.Search(
	//		es.Search.WithIndex("my-index-000001"),
	//		es.Search.WithBody(strings.NewReader(`{
	//	  "aggs": {
	//	    "by_day": {
	//	      "date_histogram": {
	//	        "field": "date",
	//	        "calendar_interval": "day"
	//	      }
	//	    }
	//	  }
	//	}`)),
	//		es.Search.WithSize(0),
	//		es.Search.WithPretty(),
	//	)
	//	fmt.Println(res, err)
	//	if err != nil { // SKIP
	//		fmt.Printf("Error getting the response: %s \n", err) // SKIP
	//	} // SKIP
	//	//defer res.Body.Close() // SKIP
	//}

	//res, err := es.Bulk(
	//	strings.NewReader(`{
	//	  "aggs": {
	//	    "by_day": {
	//	      "date_histogram": {
	//	        "field": "date",
	//	        "calendar_interval": "day"
	//	      }
	//	    }
	//	  }
	//	}`),
	//	es.Bulk.WithIndex("my-index-000001"),
	//)
	//
	//req := esapi.BulkRequest{}
	//req.Do(context.Background(), es)
}
