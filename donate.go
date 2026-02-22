package mefrpApi

import (
	"fmt"
)

// ApplyNodeDonate submits a node donation application
func (c *Client) ApplyNodeDonate(donate NodeDonate) error {
	var resp Response[any]
	err := c.request("POST", "/auth/node/donate", donate, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetUserNodeDonates retrieves the node donation applications of the current user
func (c *Client) GetUserNodeDonates() ([]NodeDonate, error) {
	var resp Response[[]NodeDonate]
	err := c.request("GET", "/auth/node/donate/list", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// ApplyNodeDelete submits a node deletion request
func (c *Client) ApplyNodeDelete(nodeID int64, reason string) error {
	req := struct {
		NodeID int64  `json:"nodeId"`
		Reason string `json:"reason"`
	}{NodeID: nodeID, Reason: reason}

	var resp Response[any]
	err := c.request("POST", "/auth/node/donate/delete/apply", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetUserNodeDeleteRequests retrieves the node deletion requests of the current user
func (c *Client) GetUserNodeDeleteRequests() ([]NodeDeleteRequest, error) {
	var resp Response[[]NodeDeleteRequest]
	err := c.request("GET", "/auth/node/donate/delete/list", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// ApplyNodeEdit submits a node edit request
func (c *Client) ApplyNodeEdit(req NodeEditRequest) error {
	var resp Response[any]
	err := c.request("POST", "/auth/node/donate/edit/apply", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetUserNodeEditRequests retrieves the node edit requests of the current user
func (c *Client) GetUserNodeEditRequests() ([]NodeEditRequest, error) {
	var resp Response[[]NodeEditRequest]
	err := c.request("GET", "/auth/node/donate/edit/list", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// GetInstallScript retrieves the installation script for a node
func (c *Client) GetInstallScript(nodeID int64, system, arch, nodeType string) (*GetInstallScriptResponse, error) {
	req := struct {
		NodeID   string `json:"nodeId"`
		System   string `json:"system"`
		Arch     string `json:"arch"`
		NodeType string `json:"nodeType"`
	}{
		NodeID:   fmt.Sprintf("%d", nodeID),
		System:   system,
		Arch:     arch,
		NodeType: nodeType,
	}

	var resp Response[GetInstallScriptResponse]
	err := c.request("POST", "/auth/node/donate/script", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}
