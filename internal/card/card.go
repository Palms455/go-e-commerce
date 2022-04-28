package card

import (
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/v72"
)

type Card struct {
	Secret, Key, Currency string
}

type Transaction struct {
	TransactionStatusID int
	Amount int
	Currency string
	LastFour string
	BankReturnCode string
}

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	// create payment intent
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(amount)),
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