// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT license.
// Details in the LICENSE file.

package trello

import (
	"testing"
)

func TestGetOrganization(t *testing.T) {
	organization := testOrganization(t)
	if organization.DisplayName != "Culture Foundry" {
		t.Errorf("Expected name 'Culture Foundry'. Got '%s'.", organization.DisplayName)
	}
}

func TestCreateOrganization(t *testing.T) {
	client := testClient()
	// client.BaseURL = mockResponse("organizations", "culturefoundry.json").URL
	err := client.CreateOrganization(&Organization{DisplayName: "1", Name: "2"}, Defaults())
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrganizationSetClient(t *testing.T) {
	o := Organization{}
	client := testClient()
	o.SetClient(client)
	if o.client == nil {
		t.Error("Expected non-nil Organization.client")
	}
}

func testOrganization(t *testing.T) *Organization {
	client := testClient()
	// client.BaseURL = mockResponse("organizations", "culturefoundry.json").URL
	organization, err := client.GetOrganization("5f19949edb414362a6ebd013", Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return organization
}


