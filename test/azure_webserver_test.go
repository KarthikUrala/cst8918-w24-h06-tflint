package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "fdcbdc66-325a-470d-957c-7977b0fb7718"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// Override the default terraform variables
		Vars: map[string]interface{}{
			"labelPrefix": "sama0096", // Make sure to replace this with your actual username
			"region":      "East US",  // Set the region to a valid Azure region
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of output variables
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	nicName := terraform.Output(t, terraformOptions, "nic_name")

	// Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))

	// Confirm NIC exists and is connected to VM
	actualNicNames := azure.GetVirtualMachineNics(t, vmName, resourceGroupName, subscriptionID)
	assert.Equal(t, nicName, actualNicNames[0])

	// Confirm the VM is running the correct Ubuntu version
	vmImage := azure.GetVirtualMachineImage(t, vmName, resourceGroupName, subscriptionID)
	expectedOSPublisher := "Canonical"
	expectedOSVersion := "22_04-lts-gen2"
	assert.Equal(t, expectedOSPublisher, vmImage.Publisher)
	assert.Equal(t, expectedOSVersion, vmImage.SKU)
}
