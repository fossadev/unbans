package twitchapi

import (
	"net/http"
)

type User struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	ProfileImageURL string `json:"profile_image_url"`
}

type GetUsersRequest struct {
	IDs    []string
	Logins []string
	Token  string
}

type usersResponse struct {
	Data []*User `json:"data"`
}

func (t *twitchAPI) GetUsers(req *GetUsersRequest) ([]*User, error) {
	if len(req.IDs)+len(req.Logins) > 100 {
		return nil, MaxQueryErr
	}

	qs := make([]*Query, 0)

	for _, id := range req.IDs {
		qs = append(qs, &Query{Key: "id", Value: id})
	}

	for _, login := range req.Logins {
		qs = append(qs, &Query{Key: "login", Value: login})
	}

	var body usersResponse
	if err := t.request(http.MethodGet, "https://api.twitch.tv/helix/users", qs, req.Token, &body); err != nil {
		return nil, err
	}

	return body.Data, nil
}
