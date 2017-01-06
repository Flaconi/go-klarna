package go_klarna

import (
	"encoding/json"
	"fmt"
)

const (
	OrderManagementEndpoint = "/ordermanagement/v1/orders"

	Authorized   = "AUTHORIZED"
	PartCaptured = "PART_CAPTURED"
	Captured     = "CAPTURED"
	Cancelled    = "CANCELLED"
	Expired      = "EXPIRED"
	Closed       = "LOSED"

	Accepted = "ACCEPTED"
	Pending  = "PENDING"
	Rejected = "REJECTED"
)

type (
	OrderManagementSrv interface {
		// Order Management - order end-points
		GetOrder(string) (*OrderManagementOrder, error)
		AcknowledgeOrder(string) error
		SetOrderAmountLines(string, *OrderAmountLines) error
		AdjustOrderAmountLines(string, *AdjustAmountLines) error
		CancelOrder(string) error
		UpdateCustomerAddress(string, *CustomerAddress) error
		ExtendAuthorizationTime(string) error
		UpdateMerchantReferences(string, *MerchantReferences) error
		ReleaseRemainingAuthorization(string) error

		// Order Management - capture end-points
		GetRefund(string, string) error
		CreateRefund(string, *OrderManagementRefund) error
		GetAllCaptures(string) ([]*Capture, error)
		TriggerResendCustomerCommunication(string, string) error
		AddCaptureShippingInfo(string, string, []*OrderManagementShippingInfo) error
		GetCapture(string, string) (*Capture, error)
		CreateCapture(string, *CreateCapture) error
	}

	orderManagementSrv struct {
		client Client
	}

	OrderManagementOrder struct {
		ID                        string                   `json:"order_id,omitempty"`
		Status                    string                   `json:"status,omitempty"`
		FraudStatus               string                   `json:"fraud_status,omitempty"`
		OrderAmount               int                      `json:"order_amount,omitempty"`
		OriginalOrderAmount       int                      `json:"original_order_amount,omitempty"`
		CapturedAmount            int                      `json:"captured_amount,omitmepty"`
		RefundedAmount            int                      `json:",omitempty"`
		RemainingAuthorizedAmount int                      `json:"remaining_authorized_amount,omitempty"`
		PurchaseCurrency          string                   `json:"purchase_currency,omitempty"`
		Locale                    string                   `json:",omitempty"`
		OrderLines                []*Line                  `json:"order_lines,omitempty"`
		MerchantReference1        string                   `json:"merchant_reference1,omitempty"`
		MerchantReference2        string                   `json:"merchant_reference2,omitempty"`
		KlarnaReference           string                   `json:"klarna_reference"`
		Customer                  *OrderManagementCustomer `json:"customer,omitempty"`
		BillingAddress            *Address                 `json:"billing_address,omitempty"`
		ShippingAddress           *Address                 `json:"shipping_address,omitempty"`
		CreatedAt                 string                   `json:"created_at,omitempty"` // DateTime string of ISO 8601
		PurchaseCountry           string                   `json:"purchase_country,omitempty"`
		ExpiresAt                 string                   `json:"expires_at,omitempty"` // DateTime string of ISO 8601
		Captures                  *[]Capture               `json:"captures,omitempty"`
		Refunds                   *OrderManagementRefund   `json:"refunds,omitempty"`
		MerchantData              string                   `json:"merchant_data,omitempty"`
	}

	OrderManagementCustomer struct {
		DateOfBirth                  string `json:"date_of_birth,omitempty"`
		NationalIdentificationNumber string `json:"national_identification_number,omitempty"`
	}

	Capture struct {
		ID              string                       `json:"capture_id,omitempty"`
		KlarnaReference string                       `json:"klarna_reference,omitempty"`
		CaptureAmount   int                          `json:"capture_amount,omitempty"`
		CapturedAt      string                       `json:"captured_at,omitempty"` // DateTime string of ISO 8601
		Description     string                       `json:"description,omitempty"`
		OrderLines      []*Line                      `json:"order_lines,omitempty"`
		RefundedAmount  int                          `json:"refunded_amount,omitempty"`
		BillingAddress  *Address                     `json:"billing_address,omitempty"`
		ShippingAddress *Address                     `json:"shipping_address,omitempty"`
		ShippingInfo    *OrderManagementShippingInfo `json:"shipping_info,omitempty"`
	}

	OrderManagementShippingInfo struct {
		ShippingCompany       string `json:"shipping_company,omitempty"`
		ShippingMethod        string `json:"shipping_method,omitempty"`
		TrackingNumber        string `json:"tracking_number,omitempty"`
		TrackingUri           string `json:"tracking_uri,omitempty"`
		ReturnShippingCompany string `json:"return_shipping_company,omitempty"`
		ReturnTrackingNumber  string `json:"return_tracking_number,omitempty"`
		ReturnTrackingUri     string `json:"return_tracking_uri,omitempty"`
	}

	OrderManagementRefund struct {
		RefundAmount int     `json:"refund_amount,omitempty"`
		RefundedAt   string  `json:"refunded_at,omitempty"` // DateTime string of ISO 8601
		Description  string  `json:"description,omitempty"`
		OrderLines   []*Line `json:"order_lines,omitempty"`
	}

	OrderAmountLines struct {
		OrderAmount int     `json:"order_amount"`
		Description string  `json:"description,omitempty"`
		OrderLines  []*Line `json:"order_lines,omitempty"`
	}

	AdjustAmountLines struct {
		AdjustAmount int     `json:"adjust_amount"`
		Description  string  `json:"description,omitempty"`
		OrderLines   []*Line `json:"order_lines,omitempty"`
	}

	CustomerAddress struct {
		ShippingAddress *Address `json:"shipping_address,omitempty"`
		BillingAddress  *Address `json:"billing_address,omitempty"`
	}

	MerchantReferences struct {
		MerchantReference1 string `json:"merchant_reference1,omitmepty"`
		MerchantReference2 string `json:"merchant_reference2,omitmepty"`
	}

	CreateCapture struct {
		CapturedAmount int                            `json:"captured_amount"`
		Description    string                         `json:"description,omitempty"`
		OrderLines     []*Line                        `json:"order_lines,omitempty"`
		ShippingInfo   []*OrderManagementShippingInfo `json:"shipping_info,omitempty"`
		ShippingDelay  int                            `json:"shipping_delay,omitempty"`
	}
)

func (srv *orderManagementSrv) GetRefund(oid, rid string) error {
	path := fmt.Sprintf("%s/%s/refunds/%s", OrderManagementEndpoint, oid, rid)
	_, err := srv.client.Get(path)

	return err
}

func (srv *orderManagementSrv) CreateRefund(oid string, rf *OrderManagementRefund) error {
	path := fmt.Sprintf("%s/%s/refunds", OrderManagementEndpoint, oid)
	_, err := srv.client.Post(path, rf)

	return err
}

func (srv *orderManagementSrv) TriggerResendCustomerCommunication(oid, cid string) error {
	path := fmt.Sprintf("%s/%s/captures/%s/trigger-send-out", OrderManagementEndpoint, oid, cid)
	_, err := srv.client.Post(path, nil)

	return err
}

func (srv *orderManagementSrv) AddCaptureShippingInfo(oid, cid string, si []*OrderManagementShippingInfo) error {
	path := fmt.Sprintf("%s/%s/captures/%s/shipping-info", OrderManagementEndpoint, oid, cid)
	_, err := srv.client.Post(path, si)

	return err
}

func (srv *orderManagementSrv) GetCapture(oid, cid string) (*Capture, error) {
	path := fmt.Sprintf("%s/%s/captures/%s", OrderManagementEndpoint, oid, cid)
	res, err := srv.client.Get(path)
	if nil != err {
		return nil, err
	}

	var capture *Capture
	err = json.NewDecoder(res.Body).Decode(&capture)

	return capture, err
}

func (srv *orderManagementSrv) CreateCapture(oid string, c *CreateCapture) error {
	path := fmt.Sprintf("%s/%s/captures", OrderManagementEndpoint, oid)
	_, err := srv.client.Post(path, c)

	return err
}

func (srv *orderManagementSrv) GetAllCaptures(oid string) ([]*Capture, error) {
	path := fmt.Sprintf("%s/%s/captures", OrderManagementEndpoint, oid)
	res, err := srv.client.Get(path)
	if nil != err {
		return nil, err
	}
	var captures []*Capture
	err = json.NewDecoder(res.Body).Decode(&captures)

	return captures, err
}

func (srv *orderManagementSrv) GetOrder(id string) (*OrderManagementOrder, error) {
	path := OrderManagementEndpoint + "/" + id

	res, err := srv.client.Get(path)
	if nil != err {
		return nil, err
	}

	o := new(OrderManagementOrder)
	err = json.NewDecoder(res.Body).Decode(o)

	return o, err
}

func (srv *orderManagementSrv) AcknowledgeOrder(oid string) error {
	path := fmt.Sprintf("%s/%s/acknowledge", OrderManagementEndpoint, oid)
	_, err := srv.client.Post(path, nil)

	return err
}

func (srv *orderManagementSrv) SetOrderAmountLines(oid string, oal *OrderAmountLines) error {
	path := fmt.Sprintf("%s/%s/authorization", OrderManagementEndpoint, oid)
	_, err := srv.client.Patch(path, oal)

	return err
}

func (srv *orderManagementSrv) AdjustOrderAmountLines(oid string, adjust *AdjustAmountLines) error {
	path := fmt.Sprintf("%s/%s/authorization-adjustments", OrderManagementEndpoint, oid)
	_, err := srv.client.Post(path, adjust)

	return err
}

func (srv *orderManagementSrv) CancelOrder(oid string) error {
	path := fmt.Sprintf("%s/%s/cancel", OrderManagementEndpoint, oid)
	_, err := srv.client.Post(path, nil)

	return err
}

func (srv *orderManagementSrv) UpdateCustomerAddress(oid string, ca *CustomerAddress) error {
	path := fmt.Sprintf("%s/%s/customer-details", OrderManagementEndpoint, oid)
	_, err := srv.client.Post(path, ca)

	return err
}

func (srv *orderManagementSrv) ExtendAuthorizationTime(oid string) error {
	path := fmt.Sprintf("%s/%s/extend-authorization-time", OrderManagementEndpoint, oid)
	_, err := srv.client.Post(path, nil)

	return err
}

func (srv *orderManagementSrv) UpdateMerchantReferences(oid string, mr *MerchantReferences) error {
	path := fmt.Sprintf("%s/%s/merchant-references", OrderManagementEndpoint, oid)
	_, err := srv.client.Post(path, mr)

	return err
}

func (srv *orderManagementSrv) ReleaseRemainingAuthorization(oid string) error {
	path := fmt.Sprintf("%s/%s/release-remaining-authorization", OrderManagementEndpoint, oid)
	_, err := srv.client.Post(path, nil)

	return err
}

func NewOrderManagement(c Client) OrderManagementSrv {
	return &orderManagementSrv{c}
}
