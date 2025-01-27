regions {
  config = ["uksouth","ukwest"]
  create_region_stacks = false
}

environments {
  config = ["development"]
  create_environment_stacks = true
}

stacks {
  stack "virtual_networks" {
    tags = ["connectivity", "test3"]
    stack "network_security_groups" {
      tags = ["connectivity", "test1"]
      description = "Example description for NSGs"
      stack "test" {
        tags = ["testtag"]
      }
    }
  }
}
