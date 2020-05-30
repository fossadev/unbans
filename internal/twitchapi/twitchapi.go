package twitchapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var MaxQueryErr = errors.New("cannot request more than 100 items at a time from the twitch api")

type TwitchAPI interface {
	GetUsers(req *GetUsersRequest) (users []*User, err error)
}

type Query struct {
	Key   string
	Value string
}

type twitchAPI struct {
	client   *http.Client
	clientID string
}

func New(clientID string) TwitchAPI {
	return &twitchAPI{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
		clientID: clientID,
	}
}

func (t *twitchAPI) prepareReq(method, uri string, qs []*Query, token string) (*http.Request, error) {
	if len(qs) > 0 {
		uri += "?" + buildQuery(qs)
	}

	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Client-ID", t.clientID)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (t *twitchAPI) request(method, uri string, qs []*Query, token string, out interface{}) error {
	req, err := t.prepareReq(method, uri, qs, token)
	if err != nil {
		return errors.Wrap(err, "failed to make request")
	}

	res, err := t.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= 300 {
		return fmt.Errorf("non 2xx http status code returned: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(out); err != nil {
		return errors.Wrap(err, "failed to decode body")
	}

	return nil
}

func buildQuery(queries []*Query) string {
	qs := make([]string, len(queries))
	for i, query := range queries {
		qs[i] = url.QueryEscape(query.Key) + "=" + url.QueryEscape(query.Value)
	}
	return strings.Join(qs, "&")
}
