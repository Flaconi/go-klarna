package go_klarna

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	EuroAPI = "https://api.klarna.com/"
	UsAPI   = "https://api-na.klarna.com/"
)

var (
	// ErrOrderCreate error describes that Klarna API is refusing to create the order
	ErrOrderCreate = errors.New("unable to create an order, some fields constraint was violated")
	// ErrUnAuthorized error describes that you are not authorized to perform such an operation
	ErrUnAuthorized = errors.New("You were not unauthorized to execute this operation")
	// ErrReadOnlyResource error describes that the target can not be modified
	ErrReadOnlyResource = errors.New("You tried to modify a read only resource")
	// ErrOrderNotFound error describes that there are no resource found with a provided id
	ErrOrderNotFound = errors.New("No orders found with given ID")
	// ServiceUnavailable error describes any non-classified error response, e.g. 500 status codes
	ServiceUnavailable = errors.New("service temporary unavailable")
)

// Config type is the basic configurations required from the client to provide in order to function
type Config struct {
	BaseURL     *url.URL
	APIUsername string
	APIPassword string
	Timeout     time.Duration
}

// Client type abstract the functionality that the client should implement, just for more extendability
type Client interface {
	Post(path string, body interface{}) (*http.Response, error)
	Patch(path string, body interface{}) (*http.Response, error)
	Get(path string) (*http.Response, error)
	Delete(path string) (*http.Response, error)
}

type client struct {
	config Config
	client *http.Client
}

// Post method executes a Post request on the given path with the given body, if the body is empty will be omitted
func (c *client) Post(path string, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, path, body)
}

// Get method fetches the content of the given path, return http response and error interface if there is any
func (c *client) Get(path string) (*http.Response, error) {
	return c.do(http.MethodGet, path, nil)
}

// Delete method executes a Delete request on the given path and returns response pointer and error interface if there
// is any
func (c *client) Delete(path string) (*http.Response, error) {
	return c.do(http.MethodDelete, path, nil)
}

// Post method executes a Post request on the given path with the given body, if the body is empty will be omitted
func (c *client) Patch(path string, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, path, body)
}

func (c *client) do(method string, path string, body interface{}) (*http.Response, error) {
	uri := fmt.Sprintf(
		"%s://%s%s",
		c.config.BaseURL.Scheme,
		strings.TrimRight(c.config.BaseURL.Host, "/"),
		path,
	)
	var reader io.Reader
	if nil != body {
		bytesBody, err := json.Marshal(body)
		if nil != err {
			return nil, err
		}
		reader = bytes.NewReader(bytesBody)
	}

	req, err := http.NewRequest(method, uri, reader)
	if nil != err {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.config.APIUsername, c.config.APIPassword)
	res, err := c.client.Do(req)
	if nil != err {
		return nil, err
	}

	err = c.errorFromResponse(res)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// errorFromResponse method translates the sent status code into an internal error
func (c *client) errorFromResponse(res *http.Response) error {
	if res.StatusCode > 299 {
		switch res.StatusCode {
		case http.StatusBadRequest:
			return ErrOrderCreate
		case http.StatusUnauthorized:
			return ErrUnAuthorized
		case http.StatusForbidden:
			return ErrReadOnlyResource
		case http.StatusNotFound:
			return ErrOrderNotFound
		default:
			return ServiceUnavailable
		}
	}

	return nil
}

// NewClient factory method
func NewClient(c Config) Client {
	if nil == c.BaseURL {
		uri, _ := url.Parse(EuroAPI)
		c.BaseURL = uri
	}
	if 0 == c.Timeout {
		c.Timeout = time.Second * 5
	}

	return &client{
		config: c,
		client: &http.Client{
			Timeout: c.Timeout,
		},
	}
}
