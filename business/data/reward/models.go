package reward

import "time"

// Define a reward
type Reward struct {
	Id int64 // Unique identifier
	Offers []Offer // Special offer for few selected customers. Must have at least one offer at any point in time. Don't enforce it initially.
	Discounts []Discount // Special discounts for few selected customers. Must have at least one discount at any point in time. Don't enforce it initially.
	Points []Point // Points system for every customer. Must have at least one point system at any point in time. Don't enforce it initially.
	Services []Service // Special services for few customers that have met certain conditions.
}


// Define an offer
type Offer struct {
	Id int64
	Type OfferType
	Constraints []OfferConstraint
	Conditions []OfferCondition
}


// Define an offer type
type OfferType struct {
	Type string
}


// Define an offer constraint
type OfferConstraint struct {
	Time time.Time
	Duration time.Duration
}


// Define an offer condition
type OfferCondition struct {

}


// Define a discount
type Discount struct {
	Id int64
	Constraints []DiscountConstraint // Potential list of constraints like time duration, specific dates like black friday, remaining quantity constraints
}


// Define a discount constraint
type DiscountConstraint struct {

}


// Define a rewards Point system
type Point struct {
	IncrementConditions []PointIncrementCondition
	DecrementConditions []PointDecrementCondition

	DecrementPoints bool
	RedeemForDiscount bool
	RedeemForCash bool
	RedeemForService bool
}


type PointIncrementCondition struct {

}


type PointDecrementCondition struct {

}


// Define a Service
type Service struct {

}





