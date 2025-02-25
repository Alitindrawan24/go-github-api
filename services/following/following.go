package following

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetFollowing(token string, params Params) ([]Following, error) {
	baseURL := "https://api.github.com/user/following"

	var result []Following

	u, err := url.Parse(baseURL)
	if err != nil {
		return []Following{}, err
	}

	q := u.Query()
	q.Set("per_page", fmt.Sprintf("%d", params.PerPage))
	q.Set("page", fmt.Sprintf("%d", params.Page))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return []Following{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []Following{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Following{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return []Following{}, err
	}

	return result, nil
}
