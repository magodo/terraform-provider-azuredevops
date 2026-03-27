package acceptance

import (
	"testing"

	"github.com/microsoft/terraform-provider-azuredevops/internal/utils/rd"
)

type TestData struct {
	// Either the resource type (e.g. azuredevops_project) or the data source type (e.g. data.azuredevops_project)
	ResourceType  string
	ResourceLabel string
	RandData      rd.RandomData
}

func (d TestData) ResourceAddr() string {
	return d.ResourceType + "." + d.ResourceLabel
}

// BuildTestData generates some test data for the given resource
func BuildTestData(t *testing.T, resourceType string, resourceLabel string) TestData {
	testData := TestData{
		ResourceType:  resourceType,
		ResourceLabel: resourceLabel,
		RandData:      rd.NewRandomData(5),
	}

	return testData
}
