package cards

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/paymentmethod"
)

type Card struct {
	Secret, Key, Currency string
}

type Transaction struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFour            string
	BankReturnCode      string
}

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	// create payment intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	//params.AddMetadata()
	pi, err := paymentintent.New(params)

	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code)
		}
		return nil, msg, err
	}
	return pi, "", err
}

// GetPaymentMethod get payment method by payment intend id
func(c *Card) GetPaymentMethod(s string) (*stripe.PaymentMethod, error) {
	stripe.Key = c.Secret

	pm, err := paymentmethod.Get(s, nil)
	if err != nil {
		return nil ,err
	}
	return pm, nil
}

// RetrievePaymentIntent gets an existing payment intent by id
func (c *Card) RetrievePaymentIntent(id string) (*stripe.PaymentIntent, error){
	stripe.Key = c.Secret

	pi, err := paymentintent.Get(id, nil)
	if err != nil {
		return nil, err
	}
	return pi, err
}

// cardErrorMessage returns verbose error for card messages
func cardErrorMessage(code stripe.ErrorCode) string {
	var msg = ""
	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Card is declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "Card is expired"
	case stripe.ErrorCodeIncorrectCVC:
		msg = "Incorrect CVC code"
	case stripe.ErrorCodeIncorrectZip:
		msg = "Incorrect ZIP/Postal code"
	case stripe.ErrorCodeAmountTooLarge:
		msg = "This amount too large to charge to your card"
	case stripe.ErrorCodeAmountTooSmall:
		msg = "This amount too small to charge to your card"
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Insufficient valance"
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "Postal code is invalid"
	default:
		msg = "Invalid Card"
	}
	return msg
}
