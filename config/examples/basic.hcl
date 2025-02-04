regions {
  region "ukwest" {
    tags = []
    deploy_as_stack = false
  }
  region "uksouth" {
    tags = ["test"]
    deploy_as_stack = true
  }
}

stacks {

  stack "StackA" {
    tags = ["testB"]
    stack "StackB" {
      tags = []
    }
  }
}

