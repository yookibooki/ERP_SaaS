# ERP SaaS

## Overview
This is a SaaS ERP application built using Go.

## Directory Structure
- **`/cmd`** - Main applications, with each subdirectory matching the executable name (e.g., `/cmd/myapp`). Contains minimal code, with `main` functions importing from `/internal` or `/pkg`.
- **`/internal`** - Private application and library code, enforced by the Go compiler to prevent external imports. Can include `/internal/pkg` for shared code and `/internal/app` for app-specific code.
- **`/pkg`** - Reusable library code safe for external use (e.g., `/pkg/mypubliclib`).
- **`/api`** - OpenAPI/Swagger specs, JSON schema files, and protocol definitions for APIs.
- **`/web`** - Web application components like static assets, server-side templates, and SPAs.
- **`/configs`** - Configuration file templates or defaults (e.g., `confd`, `consul-template`).
- **`/init`** - System init (systemd, upstart, sysv) and process manager configs (runit, supervisord).
- **`/scripts`** - Scripts for build, install, analysis, etc., to simplify the root-level Makefile.
- **`/build`** - Packaging and CI configs; includes `/build/package` for cloud/container/OS packages and `/build/ci` for CI scripts.
- **`/deployments`** - IaaS, PaaS, and orchestration configs (e.g., docker-compose, kubernetes/helm, terraform). May be called `/deploy` for Kubernetes apps.
- **`/test`** - External test apps and data, optionally with `/test/data` or `/test/testdata` (Go ignores `.` or `_` prefixed names).
- **`/docs`** - Design and user documentation, beyond godoc.
- **`/tools`** - Project support tools, can import from `/pkg` and `/internal`.
- **`/examples`** - Application or library usage examples.
- **`/third_party`** - External tools, forked code, or utilities (e.g., Swagger UI).
- **`/githooks`** - Git hooks for pre-commit, commit-msg, and other hooks.
- **`/assets`** - Images, logos, and other repository assets.
- **`/website`** - Project website data, if not using GitHub Pages.

## Installation

## Usage

## Contributing

## License