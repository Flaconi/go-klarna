package go_klarna

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestOrderManagementSrv_GetRefund(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/refunds/cba",
		nil,
		http.MethodGet,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.GetRefund("abc", "cba")

	assertions.Empty(err)
}

func TestOrderManagementSrv_CreateRefund(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	mockedRequest := &OrderManagementRefund{}
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/refunds",
		mockedRequest,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.CreateRefund("abc", mockedRequest)

	assertions.Empty(err)
}

func TestOrderManagementSrv_TriggerResendCustomerCommunication(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/captures/cba/trigger-send-out",
		nil,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.TriggerResendCustomerCommunication("abc", "cba")

	assertions.Empty(err)
}

func TestOrderManagementSrv_AddCaptureShippingInfo(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	mockedRequest := make([]*OrderManagementShippingInfo, 0)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/captures/cba/shipping-info",
		mockedRequest,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.AddCaptureShippingInfo("abc", "cba", mockedRequest)

	assertions.Empty(err)
}

func TestOrderManagementSrv_GetCapture(t *testing.T) {
	setupServer()
	defer tearDown()

	mockedResponse := &Capture{
		"123",
		"ref-123",
		100,
		"2015-11-T01:51:17+00:00",
		"something descriptive",
		nil,
		5,
		nil,
		nil,
		nil,
	}
	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/captures/cba",
		nil,
		http.MethodGet,
		mockedResponse,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	res, err := pSrv.GetCapture("abc", "cba")

	assertions.Empty(err)
	assertions.Equal(mockedResponse, res)
}

func TestOrderManagementSrv_CreateCapture(t *testing.T) {
	setupServer()
	defer tearDown()

	mockedRequest := &CreateCapture{
		100,
		"something descriptive",
		nil,
		nil,
		5,
	}
	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/captures",
		mockedRequest,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.CreateCapture("abc", mockedRequest)

	assertions.Empty(err)
}

func TestOrderManagementSrv_GetAllCaptures(t *testing.T) {
	setupServer()
	defer tearDown()

	response := &Capture{
		"123",
		"ref-123",
		100,
		"2015-11-T01:51:17+00:00",
		"something descriptive",
		nil,
		5,
		nil,
		nil,
		nil,
	}

	mockedResponse := make([]*Capture, 1)
	mockedResponse = append(mockedResponse, response)

	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/captures",
		nil,
		http.MethodGet,
		mockedResponse,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	res, err := pSrv.GetAllCaptures("abc")

	assertions.Empty(err)
	assertions.Equal(mockedResponse, res)
}

func TestOrderManagementSrv_GetOrder(t *testing.T) {
	setupServer()
	defer tearDown()

	mockedResponse := &OrderManagementOrder{
		ID: "10000",
	}

	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc",
		nil,
		http.MethodGet,
		mockedResponse,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	res, err := pSrv.GetOrder("abc")

	assertions.Empty(err)
	assertions.Equal(mockedResponse, res)
}

func TestOrderManagementSrv_AcknowledgeOrder(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/acknowledge",
		nil,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.AcknowledgeOrder("abc")

	assertions.Empty(err)
}

func TestOrderManagementSrv_SetOrderAmountLines(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	mockedRequest := &OrderAmountLines{}
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/authorization",
		mockedRequest,
		http.MethodPatch,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.SetOrderAmountLines("abc", mockedRequest)

	assertions.Empty(err)
}

func TestOrderManagementSrv_AdjustOrderAmountLines(t *testing.T) {
	setupServer()
	defer tearDown()

	mockedRequest := &AdjustAmountLines{}
	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/authorization-adjustments",
		mockedRequest,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.AdjustOrderAmountLines("abc", mockedRequest)

	assertions.Empty(err)
}

func TestOrderManagementSrv_CancelOrder(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/cancel",
		nil,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.CancelOrder("abc")

	assertions.Empty(err)
}

func TestOrderManagementSrv_UpdateCustomerAddress(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	mockedRequest := &CustomerAddress{}
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/customer-details",
		mockedRequest,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.UpdateCustomerAddress("abc", mockedRequest)

	assertions.Empty(err)
}

func TestOrderManagementSrv_ExtendAuthorizationTime(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/extend-authorization-time",
		nil,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.ExtendAuthorizationTime("abc")

	assertions.Empty(err)
}

func TestOrderManagementSrv_UpdateMerchantReferences(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	mockedRequest := &MerchantReferences{
		"ref1", "ref2",
	}
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/merchant-references",
		mockedRequest,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.UpdateMerchantReferences("abc", mockedRequest)

	assertions.Empty(err)
}

func TestOrderManagementSrv_ReleaseRemainingAuthorization(t *testing.T) {
	setupServer()
	defer tearDown()

	// initialization
	assertions := assert.New(t)
	setupMux(
		assertions,
		"/ordermanagement/v1/orders/abc/release-remaining-authorization",
		nil,
		http.MethodPost,
		nil,
	)

	c := testingClient()
	pSrv := NewOrderManagement(c)
	err := pSrv.ReleaseRemainingAuthorization("abc")

	assertions.Empty(err)
}
