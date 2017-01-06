package go_klarna

import (
	"encoding/json"
)

const (
	checkoutEndPoint = "/checkout/v3/orders"

	// Line types
	PhysicalLineType    LineType = "physical"
	DiscountLineType             = "discount"
	ShippingFeeLineType          = "shipping_fee"
	SalesTaxLineType             = "sales_tax"
)

type (
	// CheckoutSrv type represent the method that the checkout service will expose
	CheckoutSrv interface {
		CreateNewOrder(*CheckoutOrder) error
		RetrieveOrder(string) (*CheckoutOrder, error)
		UpdateOrder(string, *CheckoutOrder) error
	}

	checkoutSrv struct {
		client Client
	}

	// LineType The applicable order lines
	LineType string

	// CheckoutOrder type is the request structure to create a new order from the Checkout API
	CheckoutOrder struct {
		ID                     string                `json:"order_id,omitempty"`
		PurchaseCountry        string                `json:"purchase_country"`
		PurchaseCurrency       string                `json:"purchase_currency"`
		Locale                 string                `json:"locale"`
		Status                 string                `json:"status,omitempty"`
		BillingAddress         *Address              `json:"billing_address,omitempty"`
		ShippingAddress        *Address              `json:"shipping_address,omitempty"`
		OrderAmount            int                   `json:"order_amount"`
		OrderTaxAmount         int                   `json:"order_tax_amount"`
		OrderLines             []*Line               `json:"order_lines"`
		Customer               *CheckoutCustomer     `json:"customer,omitempty"`
		MerchantURLS           *CheckoutMerchantURLS `json:"merchant_urls"`
		HTMLSnippet            string                `json:"html_snippet,omitempty"`
		MerchantReference1     string                `json:"merchant_reference1,omitempty"`
		MerchantReference2     string                `json:"merchant_reference2,omitempty"`
		StartedAt              string                `json:"started_at,omitempty"`
		CompletedAt            string                `json:"completed_at,omitempty"`
		LastModifiedAt         string                `json:"last_modified_at,omitempty"`
		Options                *CheckoutOptions      `json:"options,omitempty"`
		Attachment             *Attachment           `json:"attachment,omitempty"`
		ExternalPaymentMethods []*PaymentProvider    `json:"external_payment_methods,omitempty"`
		ExternalCheckouts      []*PaymentProvider    `json:"external_checkouts,omitempty"`
		ShippingCountries      []string              `json:"shipping_countries,omitempty"`
		ShippingOptions        []*ShippingOption     `json:"shipping_options,omitempty"`
		MerchantData           string                `json:"merchant_data,omitempty"`
		GUI                    *GUI                  `json:"gui,omitempty"`
		MerchantRequested      *AdditionalCheckBox   `json:"merchant_requested,omitempty"`
		SelectedShippingOption *ShippingOption       `json:"selected_shipping_option,omitempty"`
	}

	// GUI type wraps the GUI options
	GUI struct {
		Options []string `json:"options,omitempty"`
	}

	// ShippingOption type is part of the CheckoutOrder structure, represent the shipping options field
	ShippingOption struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description,omitempty"`
		Promo          string `json:"promo,omitempty"`
		Price          int    `json:"price"`
		TaxAmount      int    `json:"tax_amount"`
		TaxRate        int    `json:"tax_rate"`
		Preselected    bool   `json:"preselected,omitempty"`
		ShippingMethod string `json:"shipping_method,omitempty"`
	}

	// PaymentProvider type is part of the CheckoutOrder structure, represent the ExternalPaymentMethods and
	// ExternalCheckouts field
	PaymentProvider struct {
		Name        string   `json:"name"`
		RedirectURL string   `json:"redirect_url"`
		ImageURL    string   `json:"image_url,omitempty"`
		Fee         int      `json:"fee,omitempty"`
		Description string   `json:"description,omitempty"`
		Countries   []string `json:"countries,omitempty"`
	}

	Attachment struct {
		ContentType string `json:"content_type"`
		Body        string `json:"body"`
	}

	CheckoutOptions struct {
		AcquiringChannel               string              `json:"acquiring_channel,omitempty"`
		AllowSeparateShippingAddress   bool                `json:"allow_separate_shipping_address,omitempty"`
		ColorButton                    string              `json:"color_button,omitempty"`
		ColorButtonText                string              `json:"color_button_text,omitempty"`
		ColorCheckbox                  string              `json:"color_checkbox,omitempty"`
		ColorCheckboxCheckmark         string              `json:"color_checkbox_checkmark,omitempty"`
		ColorHeader                    string              `json:"color_header,omitempty"`
		ColorLink                      string              `json:"color_link,omitempty"`
		DateOfBirthMandatory           bool                `json:"date_of_birth_mandatory,omitempty"`
		ShippingDetails                string              `json:"shipping_details,omitempty"`
		TitleMandatory                 bool                `json:"title_mandatory,omitempty"`
		AdditionalCheckbox             *AdditionalCheckBox `json:"additional_checkbox"`
		RadiusBorder                   string              `json:"radius_border,omitempty"`
		ShowSubtotalDetail             bool                `json:"show_subtotal_detail,omitempty"`
		RequireValidateCallbackSuccess bool                `json:"require_validate_callback_success,omitempty"`
		AllowGlobalBillingCountries    bool                `json:"allow_global_billing_countries,omitempty"`
	}

	AdditionalCheckBox struct {
		Text     string `json:"text"`
		Checked  bool   `json:"checked"`
		Required bool   `json:"required"`
	}

	CheckoutMerchantURLS struct {
		// URL of merchant terms and conditions. Should be different than checkout, confirmation and push URLs.
		// (max 2000 characters)
		Terms string `json:"terms"`

		// URL of merchant checkout page. Should be different than terms, confirmation and push URLs.
		// (max 2000 characters)
		Checkout string `json:"checkout"`

		// URL of merchant confirmation page. Should be different than checkout and confirmation URLs.
		// (max 2000 characters)
		Confirmation string `json:"confirmation"`

		// URL that will be requested when an order is completed. Should be different than checkout and
		// confirmation URLs. (max 2000 characters)
		Push string `json:"push"`
		// URL that will be requested for final merchant validation. (must be https, max 2000 characters)
		Validation string `json:"validation,omitempty"`

		// URL for shipping option update. (must be https, max 2000 characters)
		ShippingOptionUpdate string `json:"shipping_option_update,omitempty"`

		// URL for shipping, tax and purchase currency updates. Will be called on address changes.
		// (must be https, max 2000 characters)
		AddressUpdate string `json:"address_update,omitempty"`

		// URL for notifications on pending orders. (max 2000 characters)
		Notification string `json:"notification,omitempty"`

		// URL for shipping, tax and purchase currency updates. Will be called on purchase country changes.
		// (must be https, max 2000 characters)
		CountryChange string `json:"country_change,omitempty"`
	}

	CheckoutCustomer struct {
		// DateOfBirth in string representation 2006-01-02
		DateOfBirth string `json:"date_of_birth"`
	}

	// Address type define the address object (json serializable) being used for the API to represent billing &
	// shipping addresses
	Address struct {
		GivenName      string `json:"given_name,omitempty"`
		FamilyName     string `json:"family_name,omitempty"`
		Email          string `json:"email,omitempty"`
		Title          string `json:"title,omitempty"`
		StreetAddress  string `json:"street_address,omitempty"`
		StreetAddress2 string `json:"street_address2,omitempty"`
		PostalCode     string `json:"postal_code,omitempty"`
		City           string `json:"city,omitempty"`
		Region         string `json:"region,omitempty"`
		Phone          string `json:"phone,omitempty"`
		Country        string `json:"country,omitempty"`
	}

	Line struct {
		Type                string `json:"type,omitempty"`
		Reference           string `json:"reference,omitempty"`
		Name                string `json:"name"`
		Quantity            int    `json:"quantity"`
		QuantityUnit        string `json:"quantity_unit,omitempty"`
		UnitPrice           int    `json:"unit_price"`
		TaxRate             int    `json:"tax_rate"`
		TotalAmount         int    `json:"total_amount"`
		TotalDiscountAmount int    `json:"total_discount_amount,omitempty"`
		TotalTaxAmount      int    `json:"total_tax_amount"`
		MerchantData        string `json:"merchant_data,omitempty"`
		ProductURL          string `json:"product_url,omitempty"`
		ImageURL            string `json:"image_url,omitempty"`
	}
)

// CreateNewOrder method create a new order on the Klarna API
func (srv *checkoutSrv) CreateNewOrder(o *CheckoutOrder) error {
	res, err := srv.client.Post(checkoutEndPoint, o)
	if nil != err {
		return err
	}

	return json.NewDecoder(res.Body).Decode(o)
}

// RetrieveOrder method fetches an order by its ID
func (srv *checkoutSrv) RetrieveOrder(id string) (*CheckoutOrder, error) {
	path := checkoutEndPoint + "/" + id
	res, err := srv.client.Get(path)
	if nil != err {
		return nil, err
	}
	o := new(CheckoutOrder)
	err = json.NewDecoder(res.Body).Decode(o)

	return o, err
}

// UpdateOrder method updates an order by a given ID and CheckOrder structure, returns error if there is any
func (srv *checkoutSrv) UpdateOrder(id string, o *CheckoutOrder) error {
	path := checkoutEndPoint + "/" + id
	res, err := srv.client.Post(path, o)
	if nil != err {
		return err
	}

	return json.NewDecoder(res.Body).Decode(o)
}

// NewCheckoutSrv factory method for the checkoutSrv
func NewCheckoutSrv(c Client) CheckoutSrv {
	return &checkoutSrv{
		c,
	}
}
