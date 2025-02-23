output "resource_group_name" {
  description = "The name of the resource group."
  value       = azurerm_resource_group.example.name
}

output "vm_name" {
  description = "The name of the virtual machine."
  value       = azurerm_virtual_machine.example.name
}

output "nic_name" {
  description = "The name of the network interface."
  value       = azurerm_network_interface.example.name
}

output "public_ip" {
  description = "The public IP address of the virtual machine."
  value       = azurerm_public_ip.example.ip_address
}