# resource-nexus

> `resource-nexus` is under development. A first release candidate will be released "soon".

Graphical user interface to manage Terraform resources, without dealing directly with the configuration files or the underlying infrastructure.

More details can be found in the [idea description](./IDEA.md).

## Components

`resource-nexus` consists of two components.

### resource-nexus-core

**resource-nexus-core** is the backbone of the whole stack.  
Triggered via REST API, resource-nexus-core executes Terraform commands to provision infrastructure.

Source code and detailed documentation is located under [resource-nexus-core](resource-nexus-core)

### resource-nexus-web

Not developed yet.

# Development

Please make sure the read the following guidelines before contributing to `resource-nexus`.

## Commit messages

Since this is a "mono-repository" for more than one component of `resource-nexus`, consistent commit names are important. This allows commits to be grouped and assigned to the corresponding component.  

Commit naming:
- `gh-actions - ...`: Prefix for GitHub Action changes
- `resource-nexus-core - ..`: Prefix for backend changes
- `resource-nexus-web - ...`: Prefix for frontend changes

Changes that are not part of actions, backend or frontend doenst need a consistent prefix.

It is important to separate changes to the backend and frontend into separate commits.

# Disclaimer

`resource-nexus` is an OSS project that uses and builds on Terraform. It is not affiliated with HashiCorp or Terraform.

# LICENSE

MIT © Tobias Bauriedel