package binance

import (
	"time"
)

func (c *Client) Ping() error {
	req, err := c.newRequest("GET", "/api/v1/ping", nil)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	return err
}

type TimeResponse struct {
	ServerTime int64
}

func (c *Client) Time() (*time.Time, error) {
	req, err := c.newRequest("GET", "/api/v1/time", nil)
	if err != nil {
		return nil, err
	}

	var timeResponse TimeResponse

	_, err = c.do(req, &timeResponse)
	if err != nil {
		return nil, err
	}

	seconds := timeResponse.ServerTime / 1000
	milliseconds := timeResponse.ServerTime - (seconds * 1000)

	serverTime := time.Unix(seconds, milliseconds)

	return &serverTime, err
}
