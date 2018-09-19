package main

import (
	"github.com/olivere/elastic"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL("http://172.25.53.90:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}


	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://172.25.53.90:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://172.25.53.90:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Search with a term query
	//termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("twitter").   // search in index "twitter"
		//Query(termQuery).   // specify the query
		//Sort("user", true). // sort by "user" field, ascending
		From(0).Size(100).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("Query total  %d entries\n", searchResult.Hits.TotalHits)

}
