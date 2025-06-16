output "resource_group_name" {
  value       = azurerm_resource_group.rg.name
  description = "The name of the Azure Resource Group"
}

output "vm_name" {
  value       = azurerm_linux_virtual_machine.webserver.name
  description = "The name of the Virtual Machine"
}

output "nic_name" {
  value       = azurerm_network_interface.webserver.name
  description = "The name of the Network Interface"
}

output "public_ip" {
  value       = azurerm_public_ip.webserver.ip_address
  description = "The Public IP address assigned to the Virtual Machine"
}
