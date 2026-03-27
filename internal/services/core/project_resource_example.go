package core

import (
	"fmt"

	"github.com/microsoft/terraform-provider-azuredevops/internal/utils/rd"
)

type projectResourceExample struct {
	Data rd.RandomData
}

func NewProjectResourceExample(d *rd.RandomData) projectResourceExample {
	if d == nil {
		return projectResourceExample{Data: rd.RandomData{Int: 0, Str: "example"}}
	}
	return projectResourceExample{Data: *d}
}

func (r projectResourceExample) Basic() string {
	return fmt.Sprintf(`
resource "azuredevops_project" "test" {
  name = "%[1]s-project"
}`, r.Data.Str)
}

func (r projectResourceExample) Complete() string {
	return fmt.Sprintf(`
resource "azuredevops_project" "test" {
  name               = "%[1]s-project"
  description        = "example description"
  version_control    = "Git"
  work_item_template = "Basic"
  features = {
    boards     = false
    repositories = false
    pipelines  = false
    testplans  = false
    artifacts  = false
  }
}`, r.Data.Str)
}
