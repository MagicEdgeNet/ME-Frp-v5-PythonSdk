package mefrp

import "fmt"

// GetNodeList retrieves the list of nodes
func (c *Client) GetNodeList() ([]Node, error) {
	var resp Response[[]Node]
	err := c.request("GET", "/auth/node/list", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// GetNodeStatus retrieves the status of nodes
func (c *Client) GetNodeStatus() ([]NodeStatus, error) {
	var resp Response[[]NodeStatus]
	err := c.request("GET", "/auth/node/status", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// GetNodeToken retrieves the token for a specific node
func (c *Client) GetNodeToken(nodeID int) (*NodeToken, error) {
	req := struct {
		NodeID int `json:"nodeId"`
	}{NodeID: nodeID}

	var resp Response[NodeToken]
	err := c.request("POST", "/auth/node/secret", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetNodeConnectionList retrieves the connection addresses for nodes (only for created proxies)
func (c *Client) GetNodeConnectionList() ([]NodeConnection, error) {
	var resp Response[[]NodeConnection]
	err := c.request("GET", "/auth/node/nameList", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}
