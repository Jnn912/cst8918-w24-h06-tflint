plugin "terraform" {
  enabled = true
<<<<<<< HEAD
  preset  = "all"
=======
  preset  = "all"  
>>>>>>> 9fa82c5f52896152a90816c502ce9befca08040b
}

plugin "azurerm" {
  enabled = true
  version = "0.25.1"
  source  = "github.com/terraform-linters/tflint-ruleset-azurerm"
}