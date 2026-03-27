package core

import (
	"fmt"

	"github.com/microsoft/terraform-provider-azuredevops/internal/utils/rd"
)

type projectDataSourceExample struct {
	Data rd.RandomData
}

func NewProjectDataSourceExample(d *rd.RandomData) projectDataSourceExample {
	if d == nil {
		return projectDataSourceExample{Data: rd.RandomData{Int: 0, Str: "example"}}
	}
	return projectDataSourceExample{Data: *d}
}

func (r projectDataSourceExample) ById() string {
	return fmt.Sprintf(`
resource "azuredevops_project" "test" {
  name               = "%[1]s-project"
  description        = "foo"
}

data "azuredevops_project" "test" {
  project_id = azuredevops_project.test.id
}`, r.Data.Str)
}

func (r projectDataSourceExample) ByName() string {
	return fmt.Sprintf(`
resource "azuredevops_project" "test" {
  name               = "%[1]s-project"
  description        = "foo"
}

data "azuredevops_project" "test" {
  name = azuredevops_project.test.name
}`, r.Data.Str)
}
