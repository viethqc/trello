// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"fmt"
)

// Organization represents a Trello organization or team, i.e. a collection of members and boards.
// https://developers.trello.com/reference/#organizations
type Organization struct {
	client      *Client
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Desc        string   `json:"desc"`
	URL         string   `json:"url"`
	Website     string   `json:"website"`
	Products    []string `json:"products"`
	PowerUps    []string `json:"powerUps"`
}

// GetOrganization takes an organization id and Arguments and either
// GETs returns an Organization, or an error.
func (c *Client) GetOrganization(orgID string, extraArgs ...Arguments) (organization *Organization, err error) {
	args := flattenArguments(extraArgs)
	path := fmt.Sprintf("organizations/%s", orgID)
	err = c.Get(path, args, &organization)
	if organization != nil {
		organization.SetClient(c)
	}
	return
}

// GetOrganization takes an organization id and Arguments and either
// GETs returns an Organization, or an error.
func (c *Client) CreateOrganization(org *Organization, extraArgs ...Arguments) error {
	path := "organizations"
	args := Arguments{
		"name":      org.Name,
		"displayName":      org.DisplayName,
		"desc":       org.Desc,
		"url":    org.URL,
		"website": org.Website,
	}

	args.flatten(extraArgs)
	err := c.Post(path, args, &org)
	if err == nil {
		org.SetClient(c)
	}
	return err
}

// SetClient can be used to override this Organization's internal connection
// to the Trello API. Normally, this is set automatically after API calls.
func (o *Organization) SetClient(newClient *Client) {
	o.client = newClient
}
