package go_klarna

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPaymentSrv_CancelExistingAuthorization(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/credit/v1/authorizations/abc",
		nil,
		http.MethodDelete,
		nil,
	)

	c := testingClient()
	pSrv := NewPaymentSrv(c)
	err := pSrv.CancelExistingAuthorization("abc")

	assertions.Empty(err)
}

func TestPaymentSrv_CreateNewOrder(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	request := &PaymentOrder{}
	assertions := assert.New(t)
	mockedResponse := &PaymentOrderInfo{
		OrderID:     "123",
		RedirectURL: "redirectUri",
		FraudStatus: "no",
	}
	setupMux(
		assertions,
		"/credit/v1/authorizations/abc/order",
		request,
		http.MethodPost,
		mockedResponse,
	)

	c := testingClient()
	pSrv := NewPaymentSrv(c)
	actualResponse, err := pSrv.CreateNewOrder("abc", request)

	assertions.Empty(err)
	assertions.Equal(mockedResponse, actualResponse)
}

func TestPaymentSrv_UpdateExistingSession(t *testing.T) {
	setupServer()
	defer tearDown()
	// initialization
	request := &PaymentOrder{}
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/credit/v1/sessions/1a2b",
		request,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewPaymentSrv(c)
	err := pSrv.UpdateExistingSession("1a2b", request)

	assertions.Empty(err)
}

func TestPaymentSrv_CreateNewSession(t *testing.T) {
	setupServer()
	defer tearDown()
	// initialization
	assertions := assert.New(t)
	mockedRes := &PaymentSession{
		SessionID:   "101",
		ClientToken: "abc",
	}
	request := &PaymentOrder{
		PurchaseCountry:  "DE",
		PurchaseCurrency: "EUR",
		Locale:           "de-DE",
		BillingAddress: &Address{
			GivenName: "Floating tester",
		},
		OrderAmount:    4,
		OrderTaxAmount: 6,
		OrderLines: []*Line{
			{
				Name:        "line 1",
				Quantity:    3,
				UnitPrice:   2,
				TotalAmount: 6,
			},
		},
	}
	// mock the server response
	setupMux(
		assertions,
		"/credit/v1/sessions",
		request,
		http.MethodPost,
		mockedRes,
	)

	c := testingClient()
	pSrv := NewPaymentSrv(c)
	ps, err := pSrv.CreateNewSession(request)

	assertions.Empty(err)
	assertions.Equal(mockedRes, ps)
}
