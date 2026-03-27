package core_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/microsoft/terraform-provider-azuredevops/internal/acceptance"
	serviceCore "github.com/microsoft/terraform-provider-azuredevops/internal/services/core"
)

func TestAccDataSourceProject_byName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuredevops_project", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: serviceCore.NewProjectDataSourceExample(&data.RandData).ByName(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "project_id"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "description"),
				resource.TestCheckResourceAttr(data.ResourceAddr(), "visibility", "private"),
				resource.TestCheckResourceAttr(data.ResourceAddr(), "version_control", "Git"),
				resource.TestCheckResourceAttr(data.ResourceAddr(), "work_item_template", "Basic"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "process_template_id"),
				resource.TestCheckResourceAttr(data.ResourceAddr(), "features.%", "5"),
			),
		},
	})
}

func TestAccDataSourceProject_byID(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuredevops_project", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: serviceCore.NewProjectDataSourceExample(&data.RandData).ById(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "project_id"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "description"),
				resource.TestCheckResourceAttr(data.ResourceAddr(), "visibility", "private"),
				resource.TestCheckResourceAttr(data.ResourceAddr(), "version_control", "Git"),
				resource.TestCheckResourceAttr(data.ResourceAddr(), "work_item_template", "Basic"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "process_template_id"),
				resource.TestCheckResourceAttr(data.ResourceAddr(), "features.%", "5"),
			),
		},
	})
}
