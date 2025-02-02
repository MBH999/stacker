environments {
  environment "prd" {
    tags = ["production"]
    deploy_as_stack = true
  }
  environment "dev" {
    tags = []
    deploy_as_stack = false
  }
  environment "tst" {
    tags = []
    deploy_as_stack = false
  }
}

regions {
  region "ukwest" {
    tags = []
    deploy_as_stack = false
  }
  region "uksouth" {
    tags = []
    deploy_as_stack = true
  }
}

stacks {

  stack "test_uksouth_dev" {
    exclude_environments = ["prd"]
    exclude_regions = ["ukwest"]
    tags = []
  
    stack "nested_stack_uksouth_dev" {
      tags = ["test"]
    }

  }

  stack "test_ukwest_prod" {
    exclude_environments = ["dev"]
    exclude_regions = ["uksouth"]
    tags = []
  
    stack "nested_stack_ukwest_prod" {
      tags = []
    }

  }
  

}

