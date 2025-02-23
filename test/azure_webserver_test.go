package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var subscriptionID string = "a55a4ed4-e1da-410b-a33d-1840ef02f227"

func TestAzureLinuxVMCreation(t *testing.T) {

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"label_prefix": "jin00098",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	nicName := terraform.Output(t, terraformOptions, "nic_name")

	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))

	actualNicNames := azure.GetVirtualMachineNics(t, vmName, resourceGroupName, subscriptionID)
	assert.Equal(t, nicName, actualNicNames[0])

	vmImage := azure.GetVirtualMachineImage(t, vmName, resourceGroupName, subscriptionID)
	expectedOSPublisher := "Canonical"
	expectedOSVersion := "18.04-LTS"
	assert.Equal(t, expectedOSPublisher, vmImage.Publisher)
	assert.Equal(t, expectedOSVersion, vmImage.SKU)

	time.Sleep(60 * time.Second)
}
