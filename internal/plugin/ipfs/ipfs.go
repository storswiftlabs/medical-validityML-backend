package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"medical-zkml-backend/pkg/config"
	"net/http"
)

func init() {
	config.NewConfig()
}

func IPFS(data any) (string, error) {

	url := config.GetConfig().Get("ipfs.url").(string)
	byteArr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payload := bytes.NewBuffer(byteArr)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {

		return "", err
	}

	auth := "Bearer " + config.GetConfig().Get("ipfs.auth").(string)

	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		return "", err
	}
	defer resp.Body.Close()

	// Check response
	if resp.Status != "200 OK" {

		return "", err
	}

	// Extract CID (Content Identifier) from the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {

		return "", err
	}

	cid := result["value"].(map[string]interface{})["cid"].(string)
	fmt.Println("CID:", cid)
	return cid, nil
}
