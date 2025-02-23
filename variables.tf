variable "label_prefix" {
  type        = string
  default     = "jin00098"  # 使用你的学院用户名
  description = "Your college username. This will form the beginning of various resource names."
}

variable "region" {
  type        = string
  default     = "westus3"
  description = "The Azure region where resources will be deployed."
}

variable "admin_username" {
  type        = string
  default     = "azureadmin"
  description = "The username for the local user account on the VM."
}