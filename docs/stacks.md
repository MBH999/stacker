# Stacks

Stacks represent everything below the regions level. You can define 1 or more stacks deep.

This means you can have full control over how you layout your stacks.

## Defining Stacks

Stacks are defined in your configuration as follows:

```hcl
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
```

## Parameters

- name (***required***): This is the name of the stack to be deployed. These should be different for each stack at that level within the tree.

- tags (***required***): A list of strings to tag the stack (this is also passed through to child stacks). The region name is automatically added as a tag, as is the name of the stack (and the name of any parent stacks)

- description (***optional***): This is passed to the stack as a description in the stack.tm.hcl

- exclude_regions (***optional***): A list of strings with the names of the region(s) that this particular stack should not be created in.

## Folder Structure

Given the above configuration, and the example defined in [Regions](./regions.md), this configuration will deploy:

```
./stacks/uksouth/stack.tm.hcl
./stacks/uksouth/stackA/stack.tm.hcl
./stacks/uksouth/stackA/childStackA/stack.tm.hcl

./stacks/ukwest/stackA/stack.tm.hcl
./stacks/ukwest/stackA/childStackA/stack.tm.hcl
./stacks/ukwest/stackB/stack.tm.hcl
```
