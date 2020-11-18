package main

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "ctcd-omnichannel-dev")
	if err != nil {
		fmt.Errorf("pubsub.NewClient: %v", err)
		return
	}

	sub, err := client.CreateSubscription(ctx, "omni-pim-product", pubsub.SubscriptionConfig{
		Topic:       "omni.pim.product",
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("CreateSubscription: %v", err)
	}
	fmt.Fprintf(w, "Created subscription: %v\n", sub)
	return nil
}
