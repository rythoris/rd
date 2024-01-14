package rd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const ApiBaseURL = "https://api.raindrop.io/rest/v1"
const HTTPClientTimeout = time.Second * 5

var APIError = errors.New("unexpected api error (status_code != 200)")

type Raindrop struct {
	ID int `json:"_id"`

	Title  string   `json:"title"`
	Desc   string   `json:"excerpt"`
	Note   string   `json:"note"`
	Link   string   `json:"link"`
	Domain string   `json:"domain"`
	Tags   []string `json:"tags"`
	Type   string   `json:"type"`

	Cover string  `json:"cover"`
	Media []Media `json:"media"`

	Created  time.Time `json:"created"`
	Modified time.Time `json:"lastUpdate"`
}

type Media struct {
	Link string `json:"link"`
}

type Tag struct {
	Name  string `json:"_id"`
	Count int    `json:"count"`
}

type Backup struct {
	ID      string    `json:"_id"`
	Created time.Time `json:"created"`
}

func GetRaindrops(token string) ([]Raindrop, error) {
	type ResJSON struct {
		Items []Raindrop `json:"items"`
	}

	body, err := request(token, "GET", "/raindrops/0", http.NoBody)
	if err != nil {
		return nil, err
	}

	var resJSON ResJSON
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %w", err)
	}

	return resJSON.Items, nil
}

func GetTags(token string) ([]Tag, error) {
	type ResJSON struct {
		Items []Tag `json:"items"`
	}

	body, err := request(token, "GET", "/tags", http.NoBody)
	if err != nil {
		return nil, err
	}

	var resJSON ResJSON
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %w", err)
	}

	return resJSON.Items, nil
}

func CreateRaindrop(token, link string, tags []string) error {
	type ReqJSON struct {
		Link string   `json:"link"`
		Tags []string `json:"tags"`
	}

	b, err := json.Marshal(ReqJSON{Link: link, Tags: tags})
	if err != nil {
		return fmt.Errorf("json marshal error: %s", err.Error())
	}

	_, err = request(token, "POST", "/raindrop", bytes.NewReader(b))
	return err
}

func RemoveRaindrop(token string, id int) error {
	_, err := request(token, "DELETE", fmt.Sprintf("/raindrop/%d", id), http.NoBody)
	return err
}

func UpdateRaindrop(token string, id int, link string, tags []string) error {
	type ReqJSON struct {
		Link string   `json:"link"`
		Tags []string `json:"tags"`
	}

	b, err := json.Marshal(ReqJSON{Link: link, Tags: tags})
	if err != nil {
		return fmt.Errorf("json marshal error: %s", err.Error())
	}

	_, err = request(token, "PUT", fmt.Sprintf("/raindrop/%d", id), bytes.NewReader(b))
	return err
}

func RenameTag(token, tag, newTag string) error {
	type ReqJSON struct {
		Replace string   `json:"replace"`
		Tags    []string `json:"tags"`
	}

	b, err := json.Marshal(ReqJSON{Tags: []string{tag}, Replace: newTag})
	if err != nil {
		return fmt.Errorf("json marshal error: %s", err.Error())
	}

	_, err = request(token, "PUT", "/tags", bytes.NewReader(b))
	return err
}

func GetBackups(token string) ([]Backup, error) {
	type ResJSON struct {
		Items []Backup `json:"items"`
	}

	body, err := request(token, "GET", "/backups", http.NoBody)
	if err != nil {
		return nil, err
	}

	var resJSON ResJSON
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %w", err)
	}

	return resJSON.Items, nil
}

func CreateBackup(token string) error {
	_, err := request(token, "GET", "/backup", http.NoBody)
	return err
}

func request(token, method, path string, body io.Reader) ([]byte, error) {
	var client = http.Client{Timeout: HTTPClientTimeout}

	req, err := http.NewRequest("GET", ApiBaseURL+path, body)
	if err != nil {
		return nil, fmt.Errorf("http request error: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client error: %w", err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("http response read error: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %s", APIError, string(b))
	}

	return b, nil
}
