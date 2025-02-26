package follower

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetFollower(token string, params Params) ([]Follower, error) {
	baseURL := "https://api.github.com/user/followers"

	var result []Follower

	u, err := url.Parse(baseURL)
	if err != nil {
		return []Follower{}, err
	}

	q := u.Query()
	q.Set("per_page", fmt.Sprintf("%d", params.PerPage))
	q.Set("page", fmt.Sprintf("%d", params.Page))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return []Follower{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []Follower{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Follower{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return []Follower{}, err
	}

	return result, nil
}

func Follow(token string, f Follower) error {
	baseURL := "https://api.github.com/user/following/" + f.Login

	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}

	q := u.Query()
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return errors.New("GitHub API returned unexpected status code: " + string(resp.StatusCode))
	}

	return nil
}
