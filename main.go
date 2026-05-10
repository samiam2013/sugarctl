package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const userEnv = "LIBRE_LINKUP_USER"
const userPass = "LIBRE_LINKUP_PASS"
const libreProduct = "llu.android"
const libreVersion = "4.16.0"
const baseURL = "https://api-us.libreview.io"

func main() {
	slog.Info("Starting sugarctl")

	if err := godotenv.Load(); err != nil {
		slog.Error("Failed to load environment", "err", err)
		return
	}
	slog.Info("Environment loaded")

	userInput, passInput := os.Getenv(userEnv), os.Getenv(userPass)
	if strings.TrimSpace(userInput) == "" || strings.TrimSpace(passInput) == "" {
		slog.Error("User or password not provided")
		return
	}
	slog.Info("User and password gathered",
		"user_len", len(userInput),
		"pass_len", len(passInput))

	acc := NewAccount(userInput, passInput)
	err := acc.login()
	if err != nil {
		slog.Error("Login failed", "error", err)
		return
	}
	slog.Info("Login successful", "account", acc)

	for {
		connData, err := acc.getConnections()
		if err != nil {
			slog.Error("Failed getting connections", "error", err)
			return
		}

		fmt.Printf("sugar level: %d mg/dL\n",
			connData.Data[0].GlucoseMeasurement.ValueInMgPerDl)
		time.Sleep(time.Second * 60)
	}

}

type account struct {
	user      string
	pass      string
	accountID string
	token     string
}

func NewAccount(user, pass string) account {
	acc := account{
		user: user,
		pass: pass,
	}
	return acc
}

// login takes the user and pass and returns the token and account id
func (acc *account) login() error {
	u, err := url.Parse(baseURL + "/llu/auth/login")
	if err != nil {
		return fmt.Errorf("Failed to parse login URL: %w", err)
	}
	bodyMap := map[string]string{"email": acc.user, "password": acc.pass}
	b, err := json.Marshal(bodyMap)
	if err != nil {
		return fmt.Errorf("Failed to marshal body: %w", err)
	}
	r, err := http.NewRequest(
		http.MethodPost,
		u.String(),
		bytes.NewBuffer(b))
	//r.Header.Set("accept-encoding", "gzip")
	r.Header.Set("Cache-Control", "No-Cache")
	r.Header.Set("Connection", "Keep-Alive")
	r.Header.Set("Product", libreProduct)
	r.Header.Set("Version", libreVersion)
	r.Header.Set("Content-Type", "application/json")

	// make the request
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("Failed to make request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body: %w", err)
	}

	lr := loginResponse{}
	if err := json.Unmarshal(body, &lr); err != nil {
		return fmt.Errorf("Failed to unmarshal response: %w", err)
	}

	acc.token = lr.Data.AuthTicket.Token
	acc.accountID = lr.Data.User.ID
	return nil
}

func (acc *account) getConnections() (connectionsResponse, error) {
	connData := connectionsResponse{}
	u, err := url.Parse(baseURL + "/llu/connections")
	if err != nil {
		return connData, fmt.Errorf("Failed to parse URL for request: %w", err)
	}

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return connData, fmt.Errorf("Failed to create request for request: %w", err)
	}
	sha256id := sha256.New()
	sha256id.Write([]byte(acc.accountID))
	hexIDhash := hex.EncodeToString(sha256id.Sum(nil))
	r.Header.Set("Authorization", "Bearer "+acc.token)
	r.Header.Set("Account-ID", string(hexIDhash))
	r.Header.Set("Cache-Control", "No-Cache")
	r.Header.Set("Connection", "Keep-Alive")
	r.Header.Set("Product", libreProduct)
	r.Header.Set("Version", libreVersion)
	r.Header.Set("Content-Type", "application/json")

	c := http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		return connData, fmt.Errorf("Failed to do request for request: %w", err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return connData, fmt.Errorf("Failed to read response body for request: %w", err)
	}

	if err := json.Unmarshal(b, &connData); err != nil {
		return connData, fmt.Errorf("Failed to unmarshal response data for request: %w", err)
	}

	return connData, nil
}
