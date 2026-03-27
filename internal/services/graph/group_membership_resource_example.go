package graph

import (
	"fmt"

	"github.com/microsoft/terraform-provider-azuredevops/internal/utils/rd"
)

type groupMembershipResourceExample struct {
	Data rd.RandomData
}

func NewGroupMembershipResourceExample(d *rd.RandomData) groupMembershipResourceExample {
	if d == nil {
		return groupMembershipResourceExample{Data: rd.RandomData{Int: 0, Str: "example"}}
	}
	return groupMembershipResourceExample{Data: *d}
}

func (r groupMembershipResourceExample) Basic() string {
	return fmt.Sprintf(`
resource "azuredevops_group" "container" {
  display_name = "%[1]s-group"
  lifecycle {
	  ignore_changes = [members]
  }
}

resource "azuredevops_group" "member" {
  display_name = "%[1]s-member"
}

resource "azuredevops_group_membership" "test" {
  group_id = azuredevops_group.container.id
  member_id = azuredevops_group.member.id
}
`, r.Data.Str)
}
