package graph_test

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/graph"
	"github.com/microsoft/terraform-provider-azuredevops/internal/acceptance"
	"github.com/microsoft/terraform-provider-azuredevops/internal/acceptance/checks"
	"github.com/microsoft/terraform-provider-azuredevops/internal/acceptance/planchecks"
	"github.com/microsoft/terraform-provider-azuredevops/internal/client"
	serviceGraph "github.com/microsoft/terraform-provider-azuredevops/internal/services/graph"
	"github.com/microsoft/terraform-provider-azuredevops/internal/utils/errorutil"
)

type GroupResource struct{}

func (p GroupResource) Exists(ctx context.Context, client *client.Client, state *terraform.InstanceState) (bool, error) {
	_, err := client.GraphClient.GetGroup(ctx, graph.GetGroupArgs{GroupDescriptor: &state.ID})
	if err == nil {
		return true, nil
	}
	if errorutil.WasNotFound(err) {
		return false, nil
	}
	return false, err
}

func TestAccGroup_vsts_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuredevops_group", "test")
	r := GroupResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: serviceGraph.NewGroupResourceExample(&data.RandData).VstsBasic(),
			Check: resource.ComposeTestCheckFunc(
				checks.ExistsInAzure(t, r, data.ResourceAddr()),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "url"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "origin"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "subject_kind"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "domain"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "principal_name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "scope"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "storage_key"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroup_vsts_scopeProject(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuredevops_group", "test")
	r := GroupResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: serviceGraph.NewGroupResourceExample(&data.RandData).ScopeProject(),
			Check: resource.ComposeTestCheckFunc(
				checks.ExistsInAzure(t, r, data.ResourceAddr()),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "url"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "origin"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "subject_kind"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "domain"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "principal_name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "storage_key"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroup_vsts_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuredevops_group", "test")
	r := GroupResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: serviceGraph.NewGroupResourceExample(&data.RandData).VstsComplete(),
			Check: resource.ComposeTestCheckFunc(
				checks.ExistsInAzure(t, r, data.ResourceAddr()),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "url"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "origin"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "subject_kind"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "domain"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "principal_name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "scope"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "storage_key"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroup_vsts_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuredevops_group", "test")
	r := GroupResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: serviceGraph.NewGroupResourceExample(&data.RandData).VstsBasic(),
			Check: resource.ComposeTestCheckFunc(
				checks.ExistsInAzure(t, r, data.ResourceAddr()),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "url"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "origin"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "subject_kind"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "domain"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "principal_name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "scope"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "storage_key"),
			),
		},
		data.ImportStep(),
		{
			Config: serviceGraph.NewGroupResourceExample(&data.RandData).VstsComplete(),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					planchecks.IsNotResourceAction(data.ResourceAddr(), plancheck.ResourceActionReplace),
				},
			},
			Check: resource.ComposeTestCheckFunc(
				checks.ExistsInAzure(t, r, data.ResourceAddr()),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "url"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "origin"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "subject_kind"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "domain"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "principal_name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "scope"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "storage_key"),
			),
		},
		data.ImportStep(),
		{
			Config: serviceGraph.NewGroupResourceExample(&data.RandData).VstsBasic(),
			Check: resource.ComposeTestCheckFunc(
				checks.ExistsInAzure(t, r, data.ResourceAddr()),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "url"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "origin"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "subject_kind"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "domain"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "principal_name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "scope"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "storage_key"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroup_aad_originId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuredevops_group", "test")
	r := GroupResource{}

	if os.Getenv("ARM_TENANT_ID") == "" {
		t.Skip("AzureAD related environment variables are not specified.")
	}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: serviceGraph.NewGroupResourceExample(&data.RandData).AadOrigin(),
			Check: resource.ComposeTestCheckFunc(
				checks.ExistsInAzure(t, r, data.ResourceAddr()),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "url"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "origin"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "subject_kind"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "domain"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "principal_name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "scope"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "storage_key"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroup_aad_mail(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuredevops_group", "test")
	r := GroupResource{}

	if os.Getenv("ARM_TENANT_ID") == "" {
		t.Skip("AzureAD related environment variables are not specified.")
	}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: serviceGraph.NewGroupResourceExample(&data.RandData).AadMail(),
			Check: resource.ComposeTestCheckFunc(
				checks.ExistsInAzure(t, r, data.ResourceAddr()),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "url"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "origin"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "subject_kind"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "domain"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "principal_name"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "scope"),
				resource.TestCheckResourceAttrSet(data.ResourceAddr(), "storage_key"),
			),
		},
		data.ImportStep(),
	})
}
