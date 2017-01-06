package go_klarna

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCheckoutSrv_CreateNewOrder(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	request := new(CheckoutOrder)
	assertions := assert.New(t)

	setupMux(
		assertions,
		"/checkout/v3/orders",
		request,
		http.MethodPost,
		mockedResponse,
	)

	c := testingClient()
	pSrv := NewCheckoutSrv(c)
	err := pSrv.CreateNewOrder(request)

	assertions.Empty(err)
}

func TestCheckoutSrv_RetrieveOrder(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)

	setupMux(
		assertions,
		"/checkout/v3/orders/abc",
		nil,
		http.MethodGet,
		mockedResponse,
	)

	c := testingClient()
	pSrv := NewCheckoutSrv(c)
	ord, err := pSrv.RetrieveOrder("abc")

	assertions.Empty(err)
	assertions.Equal(mockedResponse, ord)
}

func TestCheckoutSrv_UpdateOrder(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)

	setupMux(
		assertions,
		"/checkout/v3/orders/abc",
		nil,
		http.MethodPost,
		mockedResponse,
	)

	c := testingClient()
	pSrv := NewCheckoutSrv(c)
	err := pSrv.UpdateOrder("abc", mockedResponse)

	assertions.Empty(err)
}

var mockedResponse = &CheckoutOrder{
	PurchaseCountry:  "US",
	PurchaseCurrency: "EUR",
	Locale:           "en-US",
	OrderAmount:      0,
	OrderTaxAmount:   0,
	OrderLines: []*Line{
		{
			Name:           "line 1",
			Quantity:       3,
			UnitPrice:      2,
			TaxRate:        12,
			TotalTaxAmount: 12,
			TotalAmount:    6,
		},
	},
	MerchantURLS: &CheckoutMerchantURLS{
		Terms:        "",
		Checkout:     "",
		Confirmation: "",
		Push:         "",
	},
}
