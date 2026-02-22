package mefrpApi

import "fmt"

// GetStatistics retrieves public statistics
func (c *Client) GetStatistics() (*Statistics, error) {
	var resp Response[Statistics]
	err := c.request("GET", "/public/statistics", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetStoreItems retrieves items from the store
func (c *Client) GetStoreItems() ([]StoreItem, error) {
	var resp Response[[]StoreItem]
	err := c.request("GET", "/public/store/products", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// GetHolidayData retrieves holiday data for a given year
func (c *Client) GetHolidayData(year int) ([]string, error) {
	var resp Response[[]string]
	err := c.request("GET", fmt.Sprintf("/public/holiday?year=%d", year), nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// CheckUpdateRequest represents the update check request
type CheckUpdateRequest struct {
	ProductID      string `json:"productId"`
	CurrentVersion string `json:"currentVersion"`
	System         string `json:"system"`
	Arch           string `json:"arch"`
}

// CheckUpdate checks for software updates
func (c *Client) CheckUpdate(req CheckUpdateRequest) (*CheckUpdateResponse, error) {
	var resp Response[CheckUpdateResponse]
	err := c.request("POST", "/public/checkUpdate", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetDownloadSources retrieves the list of download sources
func (c *Client) GetDownloadSources() ([]DownloadSource, error) {
	var resp Response[[]DownloadSource]
	err := c.request("GET", "/auth/downloadSources", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// GetProducts retrieves the list of products
func (c *Client) GetProducts() ([]Product, error) {
	var resp Response[[]Product]
	err := c.request("GET", "/auth/products", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}
