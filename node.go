package mefrpApi

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

// GetNodeFreePort retrieves a free port for a specific node and protocol
func (c *Client) GetNodeFreePort(nodeID int64, protocol string) (int32, error) {
	req := struct {
		NodeID   int64  `json:"nodeId"`
		Protocol string `json:"protocol"`
	}{NodeID: nodeID, Protocol: protocol}

	var resp Response[int32]
	err := c.request("POST", "/auth/node/freePort", req, &resp)
	if err != nil {
		return 0, err
	}

	if resp.Code != 200 {
		return 0, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
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
func (c *Client) GetNodeToken(nodeID int64) (*NodeToken, error) {
	req := struct {
		NodeID int64 `json:"nodeId"`
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
