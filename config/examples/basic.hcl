regions {
  config = ["uksouth", "ukwest"]
  create_region_stacks = false
}

environments {
  config = ["production", "testing", "development"]
  create_environment_stacks = true
}

stacks {
  stack "virtual_machines" {
    tags = []
  }
  stack "virtual_networks" {
    tags = []
    stack "network_security_groups" {
      tags = []
    }
  }
  stack "firewall" {
    tags = []
    exclude_regions = ["uksouth"]
    exclude_environments = ["testing", "development"]
  }
}
