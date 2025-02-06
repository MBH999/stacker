# Regions

Regions are the "top" level of the stacks that are created.

This allows users to build multi region environmentments with ease using terramates code generation and stack filtering.

The regions block is defined as

```
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

- name: The name of the block, e.g. "uksouth" defines the name of the stack.
- tags: The tags parameter is a list of strings. This will apply tags to the stack in question. These tags will filter down to each child stack in the tree. The region name is also added as a tag.
- deploy_as_stack: This is a flag to determine if this region level should be deployed as a terramate stack. If false, the folder will not be deployed until there is a child stack defined for that region.
