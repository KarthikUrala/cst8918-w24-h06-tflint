variable "labelPrefix" {
  description = "A prefix to use for resource names"
  type        = string
  default     = "mckennr"  # Adjust this as needed
}

variable "region" {
  description = "The Azure region to create resources in"
  type        = string
  default     = "canada central"  # Adjust this as needed
}

variable "admin_username" {
  description = "The admin username for the VM"
  type        = string
  default     = "azureuser"  # Adjust this as needed
}
