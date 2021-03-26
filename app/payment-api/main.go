package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/plutov/paypal/v4"
	"log"
	"math/rand"
	"os"
)


func main(){
	godotenv.Load(".env")

	clientId := os.Getenv("PAYPAL_CLIENT_ID")
	secretId := os.Getenv("PAYPAL_SECRET")

	c, err := paypal.NewClient(clientId, secretId, paypal.APIBaseSandBox)

	if err != nil {
		log.Fatal(err)
	}
	c.SetLog(os.Stdout)



	_, err = c.GetAccessToken(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// dummy example
	payout := paypal.Payout{
		SenderBatchHeader: &paypal.SenderBatchHeader{
			SenderBatchID: "202103241234567",
			EmailSubject:  "You have a payout!",
			EmailMessage:  "You have received a payout! Thanks for using our service!",
		},
		Items: []paypal.PayoutItem{
			{
				RecipientType:   "EMAIL",
				RecipientWallet: paypal.VenmoRecipientWallet,
				Receiver:        "receiver@example.com",
				Amount: &paypal.AmountPayout{
					Value:    "9.99",
					Currency: "USD",
				},
				Note:         "Thanks for your patronage!",
				SenderItemID: fmt.Sprintf("%d",rand.Intn(100000000)),
			},
		},
	}

	payoutResp, err := c.CreatePayout(context.Background(), payout)
	if err != nil {
		panic(err)
	}

	fmt.Println("Payout Response:",payoutResp)
}
