package go_github_api

import (
	"github.com/Alitindrawan24/go-github-api/services/follower"
	"github.com/Alitindrawan24/go-github-api/services/following"
)

type GithubClient struct {
	Token string `json:"token"`
}

func New(token string) *GithubClient {
	return &GithubClient{
		Token: token,
	}
}

func (c *GithubClient) GetFollower(params follower.Params) ([]follower.Follower, error) {
	followers, err := follower.GetFollower(c.Token, params)
	if err != nil {
		return []follower.Follower{}, err
	}

	return followers, nil
}

func (c *GithubClient) GetFollowing(params following.Params) ([]following.Following, error) {
	followings, err := following.GetFollowing(c.Token, params)
	if err != nil {
		return []following.Following{}, err
	}

	return followings, nil
}

func (c *GithubClient) Unfollow(f following.Following) error {
	err := following.Unfollow(c.Token, f)
	if err != nil {
		return err
	}

	return nil
}
