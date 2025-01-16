# regions = ["uksouth", "ukwest"]
# environments = ["production", "testing", "development"]

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
