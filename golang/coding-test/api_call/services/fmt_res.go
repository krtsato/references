package services

import (
	"encoding/json"
	"fmt"
	"time"
)

// by JSON-to-GO (https://mholt.github.io/json-to-go/)
type FmtBodyType []struct {
	RenderedBody   string      `json:"rendered_body"`
	Body           string      `json:"body"`
	Coediting      bool        `json:"coediting"`
	CommentsCount  int         `json:"comments_count"`
	CreatedAt      time.Time   `json:"created_at"`
	Group          interface{} `json:"group"`
	ID             string      `json:"id"`
	LikesCount     int         `json:"likes_count"`
	Private        bool        `json:"private"`
	ReactionsCount int         `json:"reactions_count"`
	Tags           []struct {
		Name     string        `json:"name"`
		Versions []interface{} `json:"versions"`
	} `json:"tags"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	User      struct {
		Description       interface{} `json:"description"`
		FacebookID        interface{} `json:"facebook_id"`
		FolloweesCount    int         `json:"followees_count"`
		FollowersCount    int         `json:"followers_count"`
		GithubLoginName   interface{} `json:"github_login_name"`
		ID                string      `json:"id"`
		ItemsCount        int         `json:"items_count"`
		LinkedinID        interface{} `json:"linkedin_id"`
		Location          interface{} `json:"location"`
		Name              string      `json:"name"`
		Organization      interface{} `json:"organization"`
		PermanentID       int         `json:"permanent_id"`
		ProfileImageURL   string      `json:"profile_image_url"`
		TeamOnly          bool        `json:"team_only"`
		TwitterScreenName interface{} `json:"twitter_screen_name"`
		WebsiteURL        interface{} `json:"website_url"`
	} `json:"user"`
	PageViewsCount interface{} `json:"page_views_count"`
}

func FormatResponse(resBody []byte) (FmtBodyType, error) {
	fmtBody := FmtBodyType{}
	if err := json.Unmarshal(resBody, &fmtBody); err != nil {
		fmt.Errorf("Error: while unmarshalling response json")
		return nil, err
	}
	return fmtBody, nil
}
