# Managing Stacks

Stacker is built to allow users to manage tags & descriptions within your stacks, without having to flick through multiple folders and files to find what you are looking for.

This is achieved by the following:

1. When running stacker, it will check if the stack exists at the generated path.
2. If the stack does exist, it will read the stack.tm.hcl file and extract the tags and description
3. If the existing tags does not match the expected tags, it will update the stack.tm.hcl file with the correct tags.
4. If the existing description does not match the expected description, it will update the stack.tm.hcl with the correct tags.

***Ensure that your stack.tm.hcl file is only housing the terramate stack block***

Stacker will also ensure that additional stack configuration (e.g. Before or After blocks) are retained. Everything defined in the stack.tm.hcl file that is outside of the stack {} block will be lost so ensure that this data is kept seperately.
