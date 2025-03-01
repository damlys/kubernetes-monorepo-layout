# Kubernetes Monorepo Layout

This repository is designed to provide a simple structure for Kubernetes projects, enabling teams to keep everything that is needed for "infrastructure as code" together in a single place. By providing a single entry point for new team members, this repository can help teams become more productive and efficient in their work.

The Kubernetes Monorepo Layout project aims to create only a structure of directories and files. This structure includes Helm charts and releases, Kpt packages, Kustomize manifests, raw Kubernetes manifests, Terraform root modules and submodules (aka. child modules), Docker images, and more.

By using this structure, teams can keep things in the right place, making it easier to find and manage infrastructure-related files. This can help improve collaboration and ensure that everyone is on the same page when it comes to infrastructure development and deployment.

```
projects
├── core
│   ├── docker-images
│   ├── helm-charts
│   ├── helm-releases
│   ├── kpt-instances
│   ├── kpt-packages
│   ├── kubernetes-manifests
│   ├── kustomize
│   ├── terraform-modules
│   └── terraform-submodules
└── apps
    └── ...
```

We encourage contributions and suggestions to help improve the structure and contents of this repository. Whether you are a experienced infrastructure DevOps or a new joiner to the field, we hope this repository will provide a useful starting point for your projects.

## Infrastructure directories

### `/projects/<project_scope>/docker-images/<project_name>`

`docker-images` directories are intended to contain the [Dockerfiles](https://docs.docker.com/engine/reference/builder/) and associated resources needed to build container images that are used within your Kubernetes cluster.

In addition to the `Dockerfile`, these directories can also include any resources needed to build the image, such as scripts, configuration files, or other files that need to be copied into the image.

It can also include [container-structure-test.yaml](https://github.com/GoogleContainerTools/container-structure-test) files, which define tests to run on the container image.

### `/projects/<project_scope>/helm-charts/<project_name>`

`helm-charts` directories are intended to contain the [Helm charts](https://helm.sh/docs/topics/charts/). Each subdirectory should represent a different Helm chart that defines the templates and dependencies needed to install Kubernetes applications.

### `/projects/<project_scope>/helm-releases/<project_name>`

`helm-releases` directories is where you store the configuration values required to install instances of Helm charts in your Kubernetes clusters.

Each subdirectory within `helm-releases` should contain a `values.yaml` file that defines the values needed to configure a specific instance of the chart in a specific environment.

### `/projects/<project_scope>/kpt-instances/<project_name>`

`kpt-instances` directories contains "deployable instances", also known as "variants". Each instance represents a unique configuration of your software that can be deployed to a given environment such as test, staging or production.

### `/projects/<project_scope>/kpt-packages/<project_name>`

`kpt-packages` directories contains [kpt](https://kpt.dev/) "abstract packages" that are reusable. These packages can be used in different environments and provide a common set of functionality that can be shared across projects.

### `/projects/<project_scope>/kubernetes-manifests/<project_name>`

`kubernetes-manifests` directories contains raw Kubernetes `.yaml` files that define the desired state of your Kubernetes resources. These files can be used to deploy your Kubernetes applications directly, without using any additional template or customization tools.

### `/projects/<project_scope>/kustomize/<project_name>`

`kustomize` directories contains `kustomization.yaml` files and Kubernetes manifests. [Kustomize](https://kubectl.docs.kubernetes.io/references/kustomize/) is a customization engine that allows you to adapt your Kubernetes resources across multiple environments and clusters.

### `/projects/<project_scope>/terraform-modules/<project_name>`

`terraform-modules` directories contains Terraform root modules that are used to create resources in your infrastructure. Each subdirectory should represent a different module and contain all the necessary configuration files to define the providers and the resources for that module.

Terraform root modules can be applied to your infrastructure using the Terraform command line interface, and they also maintain state information to keep track of the resources they create. This state information is typically stored in a "backend", such as a cloud storage service, to ensure that it can be accessed and updated by all members of your team.

### `/projects/<project_scope>/terraform-submodules/<project_name>`

`terraform-submodules` directories contains reusable Terraform child modules that can be used by other Terraform modules in your infrastructure. Each subdirectory should represent a different module.

Unlike the modules in the `terraform-modules` directories, the modules in the `terraform-submodules` directories cannot be applied on their own and do not have state. Instead, they are designed to be used as building blocks, allowing you to define common infrastructure patterns that can be reused across your entire infrastructure.

By organizing your reusable modules in the `terraform-submodules` directories, you can easily share them across your team and avoid duplicating code in different parts of your infrastructure. You can also make changes to the module in one place and have those changes automatically propagate to all other modules that use it.

## Other directories

### `/docs`

The `docs` directory is intended to contain documentation and other resources that are relevant to your project such as:

- markdown files,
- images,
- diagrams,
- architecture decision records (ADRs).

### `/go`

The `go` directory contains [Go](https://go.dev/) programs that can be used as an alternative to shell scripts, or when shell scripts are not sufficient.

This directory should follow the structure specified in the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) repository on GitHub.

### `/scripts`

The `scripts` directory contains shell scripts that automate various procedures and tasks. These scripts can help to streamline workflows and simplify complex processes. They can be used to manage infrastructure resources, configure services, deploy applications, and perform other operations.

By storing these scripts in a central location, you can ensure that they are easily accessible to all team members, and that they can be reused across different scopes and projects.

### `/third_party`

The `third_party` directory contains manually downloaded dependencies that are required by your infrastructure but cannot be installed by any package manager.
