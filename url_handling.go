package main

import (
	"net/url"
)

func (cfg *apiConfig) ExtendUrl(pathSegment string) (*url.URL, error) {

	u, err := url.Parse(cfg.MemosURL)
	if err != nil {
		return nil, err
	}
	return u.ResolveReference(&url.URL{Path: pathSegment}), nil
}
