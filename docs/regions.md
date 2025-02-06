# Regions

Regions represent the highest level of the Terramate stack hierarchy. They allow you to build multi-region environments effortlessly using Terramate’s code generation and stack filtering features.

## Defining Regions

Regions are defined in your configuration as follows:

```hcl
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
```

## Parameters

- name (***required***): The unique identifier for the region. In this example, "uksouth" and "ukwest" serve as the names for their respective stacks.

- tags (***required***): A list of strings used to tag the region. These tags are applied to the region's stack and propagate down to each child stack. Additionally, the region name is automatically added as a tag.

- deploy_as_stack (***required***): A Boolean flag indicating whether the region should be deployed as an independent Terramate stack. If set to false, the region’s folder will not be deployed unless it contains at least one child stack.

## Folder Structure

Given the above configuration—and assuming you have at least one Stack configured—Terramate will deploy the stacks with the following folder structure:

```
./stacks/uksouth/stack.tm.hcl
./stacks/uksouth/exampleStack/stack.tm.hcl

./stacks/ukwest/exampleStack/stack.tm.hcl
```

- uksouth: Because deploy_as_stack is true, the uksouth region is deployed as its own stack, along with any child stacks (e.g., exampleStack).

- ukwest: With deploy_as_stack set to false, only the child stacks within ukwest (e.g., exampleStack) are deployed.
