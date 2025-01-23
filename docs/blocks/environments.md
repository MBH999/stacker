# Environments Block

The Environments block is used to configure the environment level folder or stack.

The environment level folder or stack is the first configurable layer.

For example:

```
./stacks/ENVIRONMENTS/region/stack/childstack
```

## Defining an Environment block

The environment block needs to be in the stacker configuration file.

Below is an example environment block;

```
environments {
  config = ["development", "production"]
  create_environment_stacks = true
}
```

- config: This is a list of strings. The items in this list will be created as stacks or folders.
- create_environment_stacks: This determines if this should be a stack or a folder. If set to "true", the environment folder will have a stack.tm.hcl file created for it.

The environment name will also be added as a tag to the stack, and all child stacks within that environment.
