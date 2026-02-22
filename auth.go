package mefrpApi

import "fmt"

// GetRegisterEmailCode requests an email verification code for registration
// Requires captcha token for human verification
func (c *Client) GetRegisterEmailCode(email, captchaToken string) error {
	req := struct {
		Email        string `json:"email"`
		CaptchaToken string `json:"captchaToken"`
	}{Email: email, CaptchaToken: captchaToken}

	var resp Response[any]
	err := c.request("POST", "/public/register/emailCode", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// RegisterRequest represents the user registration request
type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	EmailCode string `json:"emailCode"`
	Password  string `json:"password"`
}

// Register creates a new user account
func (c *Client) Register(req RegisterRequest) error {
	var resp Response[any]
	err := c.request("POST", "/public/register", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// LoginRequest represents the login request
type LoginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaToken string `json:"captchaToken"`
}

// LoginResponse represents the login response with token
type LoginResponse struct {
	Token string `json:"token"`
}

// Login authenticates a user and returns a token
// Requires captcha token for human verification
func (c *Client) Login(req LoginRequest) (string, error) {
	var resp Response[struct {
		Token string `json:"token"`
	}]
	err := c.request("POST", "/public/login", req, &resp)
	if err != nil {
		return "", err
	}

	if resp.Code != 200 {
		return "", fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	// Update client token for subsequent requests
	c.token = resp.Data.Token
	return resp.Data.Token, nil
}

// GenerateMagicLinkRequest represents the magic link generation request
type GenerateMagicLinkRequest struct {
	User         string `json:"user"`
	Callback     string `json:"callback"`
	CaptchaToken string `json:"captchaToken"`
}

// GenerateMagicLink requests a magic login link to be sent to the user's email
func (c *Client) GenerateMagicLink(req GenerateMagicLinkRequest) error {
	var resp Response[any]
	err := c.request("POST", "/public/mlogin/link", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// VerifyMagicLinkResponse represents the magic link verification response
type VerifyMagicLinkResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Group    string `json:"group"`
}

// VerifyMagicLink verifies a magic link ID and returns user information and token
func (c *Client) VerifyMagicLink(mid string) (*VerifyMagicLinkResponse, error) {
	var resp Response[VerifyMagicLinkResponse]
	err := c.request("GET", "/public/mlogin/verify?mid="+mid, nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	// Update client token for subsequent requests
	c.token = resp.Data.Token
	return &resp.Data, nil
}

// RequestIForgotEmailCodeRequest represents the request for a password recovery email code
type RequestIForgotEmailCodeRequest struct {
	Email        string `json:"email"`
	CaptchaToken string `json:"captchaToken"`
}

// RequestIForgotEmailCode requests a password recovery email code
func (c *Client) RequestIForgotEmailCode(req RequestIForgotEmailCodeRequest) error {
	var resp Response[any]
	err := c.request("POST", "/public/iforgot/emailCode", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// IForgotRequest represents the password recovery request
type IForgotRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	EmailCode string `json:"emailCode"`
}

// IForgot resets the user's password using an email code
func (c *Client) IForgot(req IForgotRequest) error {
	var resp Response[any]
	err := c.request("POST", "/public/iforgot", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// ChangePasswordRequest represents the password change request
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// ChangePassword changes the user's password
// Warning: This will reset the frp token and access key
func (c *Client) ChangePassword(req ChangePasswordRequest) error {
	var resp Response[any]
	err := c.request("POST", "/auth/user/passwordReset", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}
