package mefrpApi

import (
	"fmt"
	"net/url"
)

// GetUserAds retrieves the advertisements owned by the current user
func (c *Client) GetUserAds() ([]Ads, error) {
	var resp Response[[]Ads]
	err := c.request("GET", "/auth/ads/manage", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// GetAdsByPlacement retrieves advertisements for a specific placement
func (c *Client) GetAdsByPlacement(placement string, slotID int64) ([]Ads, error) {
	params := url.Values{}
	if placement != "" {
		params.Add("placement", placement)
	}
	if slotID != 0 {
		params.Add("slotId", fmt.Sprintf("%d", slotID))
	}

	var resp Response[[]Ads]
	err := c.request("GET", "/auth/ads/query?"+params.Encode(), nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// AddAd adds a new advertisement
func (c *Client) AddAd(ad Ads) error {
	var resp Response[any]
	err := c.request("POST", "/auth/ads/add", ad, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// UpdateAd updates an existing advertisement
func (c *Client) UpdateAd(ad Ads) error {
	var resp Response[any]
	err := c.request("POST", "/auth/ads/update", ad, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// DeleteAd deletes an advertisement
func (c *Client) DeleteAd(adsID int64) error {
	req := struct {
		AdsID int64 `json:"adsId"`
	}{AdsID: adsID}

	var resp Response[any]
	err := c.request("POST", "/auth/ads/delete", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// RenewAd renews an advertisement
func (c *Client) RenewAd(adsID int64) error {
	req := struct {
		AdsID int64 `json:"adsId"`
	}{AdsID: adsID}

	var resp Response[any]
	err := c.request("POST", "/auth/ads/renew", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// TrackAdClick tracks a click on an advertisement
func (c *Client) TrackAdClick(adsID int64) error {
	var resp Response[any]
	err := c.request("GET", fmt.Sprintf("/auth/ads/track?adId=%d", adsID), nil, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetAvailableAdSlots retrieves available advertisement slots
func (c *Client) GetAvailableAdSlots() ([]AdSlotWithUsage, error) {
	var resp Response[[]AdSlotWithUsage]
	err := c.request("GET", "/auth/ads/slots", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// GetUserAdCredits retrieves the advertisement credits for the current user
func (c *Client) GetUserAdCredits() ([]AdCredit, error) {
	var resp Response[struct {
		Credits []AdCredit `json:"credits"`
	}]
	err := c.request("GET", "/auth/ads/credits", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data.Credits, nil
}

// ApplyAd submits an advertisement application
func (c *Client) ApplyAd(ad Ads) error {
	var resp Response[any]
	err := c.request("POST", "/auth/ads/apply", ad, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetUserAdsStats retrieves advertisement statistics for the current user
func (c *Client) GetUserAdsStats() (*AdsStats, error) {
	var resp Response[AdsStats]
	err := c.request("GET", "/auth/ads/stats", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetAdSlotByPlacement retrieves an advertisement slot by its placement
func (c *Client) GetAdSlotByPlacement(placement string) (*AdSlot, error) {
	var resp Response[AdSlot]
	err := c.request("GET", "/auth/ads/slot?placement="+placement, nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// ValidateCoupon validates a discount coupon
func (c *Client) ValidateCoupon(code, productType string, orderAmount float64) (*ValidateCouponResponse, error) {
	req := struct {
		Code        string  `json:"code"`
		ProductType string  `json:"productType"`
		OrderAmount float64 `json:"orderAmount"`
	}{
		Code:        code,
		ProductType: productType,
		OrderAmount: orderAmount,
	}

	var resp Response[ValidateCouponResponse]
	err := c.request("POST", "/auth/ads/coupon/validate", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// PurchaseAdCredits purchases advertisement credits for a specific slot
func (c *Client) PurchaseAdCredits(slotID int64, amount int) error {
	req := struct {
		SlotID int64 `json:"slot_id"`
		Amount int   `json:"amount"`
	}{SlotID: slotID, Amount: amount}

	var resp Response[struct {
		Credits    int     `json:"credits"`
		Purchased  int     `json:"purchased"`
		TotalPrice float64 `json:"totalPrice"`
	}]
	err := c.request("POST", "/auth/ads/credits/purchase", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetPublicAdsByPlacement retrieves advertisements for a specific placement (Public API)
func (c *Client) GetPublicAdsByPlacement(placement string) ([]Ads, error) {
	var resp Response[[]Ads]
	err := c.request("GET", "/public/ads/query?placement="+placement, nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}
