# Kubernetes Monorepo Layout

This repository is designed to provide a simple structure for Kubernetes projects, enabling teams to keep everything that is needed for "infrastructure as code" together in a single place. By providing a single entry point for new team members, this repository can help teams become more productive and efficient in their work.

The Kubernetes Monorepo Layout project aims to create only a structure of directories and files. This structure includes Helm charts and releases, Kustomize manifests, raw Kubernetes manifests, Terraform root modules and submodules (aka. child modules), Docker images, and more.

By using this structure, teams can keep things in the right place, making it easier to find and manage infrastructure-related files. This can help improve collaboration and ensure that everyone is on the same page when it comes to infrastructure development and deployment.

```
├── docker
│   └── images
├── helm
│   ├── charts
│   └── releases
├── kubernetes
│   ├── kustomize
│   └── manifests
└── terraform
    ├── modules
    └── submodules
```

We encourage contributions and suggestions to help improve the structure and contents of this repository. Whether you are a experienced infrastructure DevOps or a new joiner to the field, we hope this repository will provide a useful starting point for your projects.

## Infrastructure directories

### `/docker/images`

This directory is intended to contain the [Dockerfiles](https://docs.docker.com/engine/reference/builder/) and associated resources needed to build container images that are used within your Kubernetes cluster. This directory can include one or more subdirectories, each representing a different image.

In addition to the `Dockerfile`, this directory can also include any resources needed to build the image, such as scripts, configuration files, or other files that need to be copied into the image.

It can also include [container-structure-test.yaml](https://github.com/GoogleContainerTools/container-structure-test) files, which define tests to run on the container image.

### `/helm/charts`

This directory is intended to contain the [Helm charts](https://helm.sh/docs/topics/charts/). Each subdirectory should represent a different Helm chart that defines the templates and dependencies needed to install Kubernetes applications.

### `/helm/releases`

The `helm/releases` directory is where you store the configuration values required to install instances of Helm charts in your Kubernetes clusters.

Each subdirectory within `helm/releases` should contain a `values.yaml` file that defines the values needed to configure a specific instance of the chart in a specific environment.

### `/kubernetes/kustomize`

The `kubernetes/kustomize` directory contains `kustomization.yaml` files and Kubernetes manifests. [Kustomize](https://kubectl.docs.kubernetes.io/references/kustomize/) is a customization engine that allows you to adapt your Kubernetes resources across multiple environments and clusters.

### `/kubernetes/manifests`

The `kubernetes/manifests` directory contains raw Kubernetes `.yaml` files that define the desired state of your Kubernetes resources. These files can be used to deploy your Kubernetes applications directly, without using any additional template or customization tools.

### `/terraform/modules`

This directory contains Terraform root modules that are used to create resources in your infrastructure. Each subdirectory should represent a different module and contain all the necessary configuration files to define the providers and the resources for that module.

Terraform root modules can be applied to your infrastructure using the Terraform command line interface, and they also maintain state information to keep track of the resources they create. This state information is typically stored in a "backend", such as a cloud storage service, to ensure that it can be accessed and updated by all members of your team.

### `/terraform/submodules`

The `terraform/submodules` directory contains reusable Terraform child modules that can be used by other Terraform modules in your infrastructure. Each subdirectory should represent a different module.

Unlike the modules in the `terraform/modules` directory, the modules in the `terraform/submodules` directory cannot be applied on their own and do not have state. Instead, they are designed to be used as building blocks, allowing you to define common infrastructure patterns that can be reused across your entire infrastructure.

By organizing your reusable modules in the `terraform/submodules` directory, you can easily share them across your team and avoid duplicating code in different parts of your infrastructure. You can also make changes to the module in one place and have those changes automatically propagate to all other modules that use it.

## Other directories

### `/docs`

The `docs` directory is intended to contain documentation and other resources that are relevant to your project such as:

- markdown files,
- images,
- diagrams,
- architecture decision records (ADRs).

### `/scripts`

The `scripts` directory contains shell scripts that automate various procedures and tasks. These scripts can help to streamline workflows and simplify complex processes. They can be used to manage infrastructure resources, configure services, deploy applications, and perform other operations.

By storing these scripts in a central location, you can ensure that they are easily accessible to all team members, and that they can be reused across different scopes and environments.

### `/third_party`

The `third_party` directory contains manually downloaded dependencies that are required by your infrastructure but cannot be installed by any package manager.
