package mefrpApi

import "fmt"

// GetUserInfo retrieves the current user's information
func (c *Client) GetUserInfo() (*UserInfo, error) {
	var resp Response[UserInfo]
	err := c.request("GET", "/auth/user/info", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// Sign performs the daily sign-in
func (c *Client) Sign(captchaToken string) error {
	req := CaptchaRequest{CaptchaToken: captchaToken}
	var resp Response[any]
	err := c.request("POST", "/auth/user/sign", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetUserFrpToken retrieves the user's frp token
func (c *Client) GetUserFrpToken() (string, error) {
	var resp Response[struct {
		Token string `json:"token"`
	}]
	err := c.request("GET", "/auth/user/frpToken", nil, &resp)
	if err != nil {
		return "", err
	}

	if resp.Code != 200 {
		return "", fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data.Token, nil
}

// GetUserGroups retrieves the user groups information
func (c *Client) GetUserGroups() ([]UserGroup, error) {
	var resp Response[UserGroupsResponse]
	err := c.request("GET", "/auth/user/groups", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data.Groups, nil
}

// ResetAccessKey resets the user's access key
func (c *Client) ResetAccessKey(captchaToken string) (string, error) {
	req := CaptchaRequest{CaptchaToken: captchaToken}
	var resp Response[ResetTokenResponse]
	err := c.request("POST", "/auth/user/tokenReset", req, &resp)
	if err != nil {
		return "", err
	}

	if resp.Code != 200 {
		return "", fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data.NewToken, nil
}

// GetUserLogs retrieves the user's operation logs
func (c *Client) GetUserLogs(filter UserOperationLogFilter) (*OperationLogList, error) {
	path := fmt.Sprintf("/auth/operationLog/list?page=%d&pageSize=%d", filter.Page, filter.PageSize)
	if filter.Category != "" {
		path += "&category=" + filter.Category
	}
	if filter.Status != "" {
		path += "&status=" + filter.Status
	}
	if filter.StartTime != "" {
		path += "&startTime=" + filter.StartTime
	}
	if filter.EndTime != "" {
		path += "&endTime=" + filter.EndTime
	}

	var resp Response[OperationLogList]
	err := c.request("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetUserLogStats retrieves the user's log statistics
func (c *Client) GetUserLogStats() (*UserLogStats, error) {
	var resp Response[UserLogStats]
	err := c.request("GET", "/auth/operationLog/stats", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetRealnameInfo retrieves the user's realname status
func (c *Client) GetRealnameInfo() (*RealnameInfoResponse, error) {
	var resp Response[RealnameInfoResponse]
	err := c.request("GET", "/auth/user/info/realname", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// PerformRealnameLegacyRequest represents the realname verification request
type PerformRealnameLegacyRequest struct {
	Realname string `json:"realname"`
	IdCard   string `json:"idCard"`
}

// PerformRealnameLegacy performs realname verification
func (c *Client) PerformRealnameLegacy(req PerformRealnameLegacyRequest) error {
	var resp Response[any]
	err := c.request("POST", "/auth/user/realname/legacy", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetUserTrafficStats retrieves the user's traffic statistics for a given period
func (c *Client) GetUserTrafficStats(datePeriod int) (*UserDailyTrafficResponse, error) {
	req := struct {
		DatePeriod int `json:"datePeriod"`
	}{DatePeriod: datePeriod}

	var resp Response[UserDailyTrafficResponse]
	err := c.request("POST", "/auth/user/trafficStats", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetUserIcpDomain retrieves the user's registered ICP domains
func (c *Client) GetUserIcpDomain() ([]IcpDomain, error) {
	var resp Response[[]IcpDomain]
	err := c.request("GET", "/auth/user/icpDomain/list", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// AddIcpDomain adds a new ICP domain
func (c *Client) AddIcpDomain(domain string) error {
	req := struct {
		Domain string `json:"domain"`
	}{Domain: domain}

	var resp Response[any]
	err := c.request("POST", "/auth/user/icpDomain/add", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// DeleteIcpDomain deletes an ICP domain
func (c *Client) DeleteIcpDomain(domain string) error {
	req := struct {
		Domain string `json:"domain"`
	}{Domain: domain}

	var resp Response[any]
	err := c.request("POST", "/auth/user/icpDomain/delete", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// KickAllProxies kicks all of the user's proxies offline
func (c *Client) KickAllProxies() error {
	var resp Response[any]
	err := c.request("GET", "/auth/user/kickAllProxies", nil, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetPurchaseStatus retrieves the user's purchase status and limits
func (c *Client) GetPurchaseStatus() (*PurchaseStatus, error) {
	var resp Response[PurchaseStatus]
	err := c.request("GET", "/auth/user/purchase-status", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetOperationLogCategories retrieves the available operation log categories
func (c *Client) GetOperationLogCategories() ([]OperationLogCategory, error) {
	var resp Response[[]OperationLogCategory]
	err := c.request("GET", "/auth/operationLog/categories", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}
