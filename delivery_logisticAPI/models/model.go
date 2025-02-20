package models

// Contains all the models used in the application

// Address struct is used from embedding in Sender and Receiver struct
type Address struct {
	Type    string `json:"type" bson:"type"`
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Pincode int    `json:"pincode" bson:"pincode"`
	Country string `json:"country" bson:"country"`
}

type Sender struct {
	ID        int     `json:"_id" bson:"_id"`
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Address   Address `json:"address" bson:"address"`
	Phone     string  `json:"phone" bson:"phone"`
}

type Receiver struct {
	ID        int     `json:"_id" bson:"_id"`
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Address   Address `json:"address" bson:"address"`
	Phone     string  `json:"phone" bson:"phone"`
}

// mongo db uses ISODate format for storing date and time
type Payment struct {
	ID             int            `json:"_id" bson:"_id"`
	InitiatedOn    string         `json:"initiated_on" bson:"initiated_on"`
	SuccessfulOn   string         `json:"successful_on" bson:"successful_on"`
	MerchantID     int            `json:"merchant_id" bson:"merchant_id"`
	ModeOfPayment  string         `json:"mode_of_payment" bson:"mode_of_payment"`
	PaymentDetails PaymentDetails `json:"payment_details" bson:"payment_details"`
}

type PaymentDetails struct {
	TransactionToken string `json:"transaction_token" bson:"transaction_token"`
}

type Carrier struct {
	ID          int    `json:"_id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	CarrierCode string `json:"carrier_code" bson:"carrier_code"`
	IsPartner   bool   `json:"is_partner" bson:"is_partner"`
}

type Shipment struct {
	ID         int    `json:"_id" bson:"_id"`
	SenderID   int    `json:"sender" bson:"sender"`
	ReceiverID int    `json:"receiver" bson:"receiver"`
	PackageID  int    `json:"package" bson:"package"`
	PaymentID  int    `json:"payment" bson:"payment"`
	CarrierID  int    `json:"carrier" bson:"carrier"`
	PromisedOn string `json:"promised_on" bson:"promised_on"`
}
