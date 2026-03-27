package graph

import (
	"fmt"

	"github.com/microsoft/terraform-provider-azuredevops/internal/utils/rd"
)

type groupResourceExample struct {
	Data rd.RandomData
}

func NewGroupResourceExample(d *rd.RandomData) groupResourceExample {
	if d == nil {
		return groupResourceExample{Data: rd.RandomData{Int: 0, Str: "example"}}
	}
	return groupResourceExample{Data: *d}
}

func (r groupResourceExample) VstsBasic() string {
	return fmt.Sprintf(`
resource "azuredevops_group" "test" {
  display_name = "%s-group"
}
`, r.Data.Str)
}

func (r groupResourceExample) VstsComplete() string {
	return fmt.Sprintf(`
resource "azuredevops_group" "test" {
  display_name = "%[1]s-group"
  description = "description"
  members = [
  	azuredevops_group.member1.id,
  	azuredevops_group.member2.id,
  ]
}

resource "azuredevops_group" "member1" {
  display_name = "%[1]s-member1"
}
resource "azuredevops_group" "member2" {
  display_name = "%[1]s-member2"
}
`, r.Data.Str)
}

func (r groupResourceExample) ScopeProject() string {
	return fmt.Sprintf(`
resource "azuredevops_project" "test" {
  name               = "%[1]s-project"
}

resource "azuredevops_group" "test" {
  scope = azuredevops_project.test.id
  display_name = "%[1]s-group"
}
`, r.Data.Str)
}

func (r groupResourceExample) AadOrigin() string {
	return fmt.Sprintf(`
data "azuread_client_config" "current" {}

resource "azuread_group" "test" {
  display_name     = "%[1]s-group"
  owners           = [data.azuread_client_config.current.object_id]
  security_enabled = true
}

resource "azuredevops_group" "test" {
  origin_id = azuread_group.test.object_id
}
`, r.Data.Str)
}

func (r groupResourceExample) AadMail() string {
	return fmt.Sprintf(`
data "azuread_client_config" "current" {}

resource "azuread_group" "test" {
  display_name     = "%[1]s-group"
  mail_enabled     = true
  mail_nickname    = "%[1]s-mail"
  types            = ["Unified"]
  owners           = [data.azuread_client_config.current.object_id]
  security_enabled = true
}

resource "azuredevops_group" "test" {
  mail = azuread_group.test.mail
}
`, r.Data.Str)
}
