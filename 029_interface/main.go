package main

import "fmt"

type payment struct {
	gatway stripe
}

func (p payment) makePayment(amount  float64){
	// razorPaymentGateWaye := razorpay{}

	// stripePaymentGateway := stripe{}
	// stripePaymentGateway.pay(amount)
	// razorPaymentGateWaye.pay(amount)

	// p.gatway.pay()
}

type razorpay struct{}

func (r razorpay) pay(amount float64){
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float64){
	fmt.Println("Making payment using stripe", amount)
}

func main(){

	stripePaymentGateway := stripe{}
	newPayment := payment{
		// getway: stripePaymentGateway,
	}
	newPayment.makePayment(100)
}