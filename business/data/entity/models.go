package entity

import (
	"github.com/nekruzvatanshoev/Microservices/business/data/reward"
)
/*
Our goal is to connect local business owners with their frequent customers. Customers that subscribe to the business owner
will receive notifications everytime the local business that they are subscribed to has (1) discounts, (2) special offers,
It's a platform that connect frequent customers with their local businesses.


Stage 1

Create a platform that connects local businesses with their frequent customers. The customers (subscribers) will receive
notifications whenever a business they are subscribed to has any (1) discounts, (2) special offers, special events,
or services for special customers, either immediately or at a custom frequency rate.


Local businesses! We will deliver you customers! You only have to pay 0.01% from the transactions coming from users on
Naboft Rewards. Online ads for all. A platform specifically for that! No ads! For every $1000 you receive from your
customers coming from our platform you will pay us $1! This platform will not have ads outside our business model!


Stage 2
Create new features to reward systems for customers.


Stage 3
A social media platform that makes it easier for customers and business owners to create relationships and connect with each other!

*/

// Define an Entity
type Entity struct {
	Id int64
	Name string
	Address string
	Category EntityCategory
	Rewards []reward.Reward
	Subscribers []Subscriber
}


type EntityCategory struct {
	Id int64
	Type string
}


type Subscriber struct {

}
