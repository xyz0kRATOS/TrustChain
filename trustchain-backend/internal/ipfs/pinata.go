package ipfs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	"github.com/amanp/trustchain-backend/internal/config"
)

const (
	pinataFileURL = "https://api.pinata.cloud/pinning/pinFileToIPFS"
	pinataJSONURL = "https://api.pinata.cloud/pinning/pinJSONToIPFS"
)

type Client struct {
	httpClient *http.Client
	jwt        string
}

type pinataResponse struct {
	IpfsHash string `json:"IpfsHash"`
}

type pinJSONRequest struct {
	PinataContent any `json:"pinataContent"`
}

type DonationMetadata struct {
	CampaignID   string `json:"campaignId"`
	CampaignName string `json:"campaignName"`
	DonorAddress string `json:"donorAddress"`
	AmountWei    string `json:"amountWei"`
	Timestamp    int64  `json:"timestamp"`
}

func NewClient(cfg *config.Config) (*Client, error) {
	jwt := strings.TrimSpace(cfg.PinataJWT)
	if jwt == "" {
		return nil, fmt.Errorf("pinata jwt is not configured")
	}

	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		jwt:        jwt,
	}, nil
}

func (c *Client) UploadFile(ctx context.Context, file []byte, filename, contentType string) (string, error) {
	if strings.TrimSpace(filename) == "" {
		return "", fmt.Errorf("filename is required")
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	partHeaders := make(textproto.MIMEHeader)
	partHeaders.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))
	if strings.TrimSpace(contentType) != "" {
		partHeaders.Set("Content-Type", contentType)
	} else {
		defaultType := mime.TypeByExtension(".bin")
		if defaultType == "" {
			defaultType = "application/octet-stream"
		}
		partHeaders.Set("Content-Type", defaultType)
	}

	part, err := writer.CreatePart(partHeaders)
	if err != nil {
		return "", fmt.Errorf("create file part: %w", err)
	}

	if _, err := part.Write(file); err != nil {
		return "", fmt.Errorf("write file content: %w", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("close multipart writer: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, pinataFileURL, &body)
	if err != nil {
		return "", fmt.Errorf("create pinata file request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.jwt)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	cid, err := c.doPinRequest(req)
	if err != nil {
		return "", err
	}

	return cid, nil
}

func (c *Client) UploadJSON(ctx context.Context, metadata any) (string, error) {
	payload := pinJSONRequest{PinataContent: metadata}

	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return "", fmt.Errorf("encode pinata json payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, pinataJSONURL, buf)
	if err != nil {
		return "", fmt.Errorf("create pinata json request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.jwt)
	req.Header.Set("Content-Type", "application/json")

	cid, err := c.doPinRequest(req)
	if err != nil {
		return "", err
	}

	return cid, nil
}

func (c *Client) UploadDonationMetadata(
	ctx context.Context,
	campaignID, campaignName, donorAddress, amountWei string,
	timestamp int64,
) (string, error) {
	metadata := DonationMetadata{
		CampaignID:   campaignID,
		CampaignName: campaignName,
		DonorAddress: donorAddress,
		AmountWei:    amountWei,
		Timestamp:    timestamp,
	}

	return c.UploadJSON(ctx, metadata)
}

func (c *Client) doPinRequest(req *http.Request) (string, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("pinata request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read pinata response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("pinata request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var parsed pinataResponse
	if err := json.Unmarshal(bodyBytes, &parsed); err != nil {
		return "", fmt.Errorf("decode pinata response: %w", err)
	}

	if strings.TrimSpace(parsed.IpfsHash) == "" {
		return "", fmt.Errorf("pinata response missing IpfsHash")
	}

	return parsed.IpfsHash, nil
}
