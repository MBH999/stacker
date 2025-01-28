regions {
  config = [
    "uksouth", 
    "ukwest"
  ]
  create_region_stacks = false
}

environments {
  config = [
    "development", 
    "prod"
  ]
  create_environment_stacks = false
}

stacks {
  stack "test_uksouth_dev" {
    exclude_environments = ["prod"]
    exclude_regions = ["ukwest"]
    tags = []
  
    stack "nested_stack_uksouth_dev" {
      tags = []
    }

  }

  stack "test_ukwest_prod" {
    exclude_environments = ["development"]
    exclude_regions = ["uksouth"]
    tags = []
  
    stack "nested_stack_ukwest_prod" {
      tags = []
    }

  }
  

}

