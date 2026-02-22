package mefrpApi

import (
	"fmt"
	"net/url"
)

// GetOrders retrieves the orders of the current user
func (c *Client) GetOrders(page, pageSize int, status string) (*struct {
	Orders []Order `json:"orders"`
	Total  int64   `json:"total"`
}, error) {
	params := url.Values{}
	params.Add("page", fmt.Sprintf("%d", page))
	params.Add("pageSize", fmt.Sprintf("%d", pageSize))
	if status != "" {
		params.Add("status", status)
	}

	var resp Response[struct {
		Orders []Order `json:"orders"`
		Total  int64   `json:"total"`
	}]
	err := c.request("GET", "/auth/orders?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// RepayOrder submits a request to repay an existing order
func (c *Client) RepayOrder(orderID, payMethod string, force bool) (*SubmitOrderResponse, error) {
	req := struct {
		OrderID   string `json:"orderId"`
		PayMethod string `json:"payMethod"`
		Force     bool   `json:"force"`
	}{OrderID: orderID, PayMethod: payMethod, Force: force}

	var resp Response[SubmitOrderResponse]
	err := c.request("POST", "/auth/cash/repay", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// SubmitOrder submits a new order for payment
func (c *Client) SubmitOrder(req SubmitOrderRequest) (*SubmitOrderResponse, error) {
	var resp Response[SubmitOrderResponse]
	err := c.request("POST", "/cash/submit", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// QueryOrder queries the status of an order
func (c *Client) QueryOrder(orderID string) (*QueryOrderResponse, error) {
	req := struct {
		OrderID string `json:"orderId"`
	}{OrderID: orderID}

	var resp Response[QueryOrderResponse]
	err := c.request("POST", "/cash/query", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// Proceed processes an order after payment
func (c *Client) Proceed(orderID string) error {
	req := struct {
		OrderID string `json:"orderId"`
	}{OrderID: orderID}

	var resp Response[any]
	err := c.request("POST", "/cash/proceed", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// RedeemCDK redeems a CDK code
func (c *Client) RedeemCDK(code, captchaToken string) (*RedeemCDKResponse, error) {
	req := struct {
		Code         string `json:"code"`
		CaptchaToken string `json:"captchaToken"`
	}{Code: code, CaptchaToken: captchaToken}

	var resp Response[RedeemCDKResponse]
	err := c.request("POST", "/auth/cdk/redeem", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetMyCDKUsage retrieves the CDK usage logs of the current user
func (c *Client) GetMyCDKUsage(page, pageSize int) (*CDKUsageLogList, error) {
	params := url.Values{}
	params.Add("page", fmt.Sprintf("%d", page))
	params.Add("pageSize", fmt.Sprintf("%d", pageSize))

	var resp Response[CDKUsageLogList]
	err := c.request("GET", "/auth/cdk/usage?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}
