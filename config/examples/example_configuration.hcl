regions {
  region "uksouth" {
    tags = []
    deploy_as_stack = true
  }
  region "ukwest" {
    tags = []
    deploy_as_stack = false
  }
}

stacks {
  stack "stackA" {
    tags = []
      
    stack "childStackA" {
      tags = []
      description = "example description"
    }
  }

  stack "stackB" {
    tags = []
    exclude_regions = ["uksouth"]
  }
}   
