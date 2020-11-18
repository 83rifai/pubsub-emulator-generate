/**/
package main

// [START pubsub_create_topic]
import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

func main() {

	csvfile, err := os.Open("pubsub.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		PubSub("tri-omnichannel-prod", record[0], record[1])
	}
}

func PubSub(project string, topic string, subs string) {
	// topicID := "my-topic"
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v", err)
		return
	}

	t := client.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {

		log.Fatal(err)
	}
	if !ok {
		t, err = client.CreateTopic(ctx, topic)
		if err != nil {
			fmt.Printf("CreateTopic: %v", err)
			return
		}
	}

	fmt.Printf("Topic created: %v\n", t)

	sub, err := client.CreateSubscription(ctx, subs, pubsub.SubscriptionConfig{
		Topic:       t,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		fmt.Errorf("CreateSubscription: %v", err)
		return
	}
	fmt.Printf("Created subscription: %v\n", sub)

}
