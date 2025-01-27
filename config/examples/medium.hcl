regions {
  config = ["uksouth"]
  create_region_stacks = false
}

environments {
  config = ["production"]
  create_environment_stacks = false
}

stacks {
  stack "1_networking" {
    tags = ["core"]
  }
  stack "2_management" {
    tags = ["core", "virtualisation"]
  }
  stack "3_kubernetes" {
    tags = ["containers"]
    stack "masters" {
      tags = ["kubernetes"]
      stack "nodes" {
        tags = ["kubernetes"]
      }
    }
  }
  stack "4_management" {
    tags = ["management"]
    stack "monitoring" {
      tags = []
    }
  }
}
