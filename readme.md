# Stacker

Stacker is a CLI tool written in Go to manage terramate stacks.

It uses the Hashicorp Configuration Language (HCL) to create and manage stacks.

## Install

To install stacker, clone this repository. You will need to ensure you have DevBox by Jetify installed on your local system.

### Install using DevBox

1. Install DevBox: https://www.jetify.com/devbox

2. Run ```devbox shell``` to enter into an isolated shell with the pre-requsite software installed.

3. Within DevBox shell, run ```task install```. This will build the go binary, and place it in your /usr/local/bin directory (sudo is required).

### Build Locally

To build locally, follow steps 1 & 2 above, then run ```task build```. This will build a stacker.bin binary in your current directory (sudo is not required).

## Usage

[Creating Regions](./docs/regions.md)

[Creating Stacks](./docs/stacks.md)

[Running Stacker](./docs/cli.md)

[Management](./docs/stacks_management.md)

[Examples](./config/examples/)
