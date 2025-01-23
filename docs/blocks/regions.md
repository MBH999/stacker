# Regions Block

The Regions block is used to configure the region level folder or stack.

The region level folder or stack is the second configurable layer.

For example:

```
./stacks/environments/REGION/stack/childstack
```

## Defining a region block

The region block needs to be in the stacker configuration file.

Below is an example region block;

```
region {
  config = ["ukwest", "uksouth"]
  create_region_stacks = true
}
```

- config: This is a list of strings. The items in this list will be created as stacks or folders.
- create_region_stacks: This determines if this should be a stack or a folder. If set to "true", the region folder will have a stack.tm.hcl file created for it.

The region name will also be added as a tag to the stack, and all child stacks within that region.
