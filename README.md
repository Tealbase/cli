# Tealbase CLI (v1)

[![Coverage Status](https://coveralls.io/repos/github/tealbase/cli/badge.svg?branch=main)](https://coveralls.io/github/tealbase/cli?branch=main)

[Tealbase](https://tealbase.io) is an open source Firebase alternative. We're building the features of Firebase using enterprise-grade open source tools.

This repository contains all the functionality for Tealbase CLI.

- [x] Running Tealbase locally
- [x] Managing database migrations
- [x] Pushing your local changes to production
- [x] Create and Deploy Tealbase Functions
- [ ] Manage your Tealbase Account
- [x] Manage your Tealbase Projects
- [x] Generating types directly from your database schema
- [ ] Generating API and validation schemas from your database

## Getting started

### Install the CLI

Available via [NPM](https://www.npmjs.com) as dev dependency. To install:

```bash
npm i tealbase --save-dev
```

To install the beta release channel:

```bash
npm i tealbase@beta --save-dev
```

> **Note**
For Bun versions below v1.0.17, you must add `tealbase` as a [trusted dependency](https://bun.sh/guides/install/trusted) before running `bun add -D tealbase`.

<details>
  <summary><b>macOS</b></summary>

  Available via [Homebrew](https://brew.sh). To install:

  ```sh
  brew install tealbase/tap/tealbase
  ```

  To install the beta release channel:
  
  ```sh
  brew install tealbase/tap/tealbase-beta
  brew link --overwrite tealbase-beta
  ```
  
  To upgrade:

  ```sh
  brew upgrade tealbase
  ```
</details>

<details>
  <summary><b>Windows</b></summary>

  Available via [Scoop](https://scoop.sh). To install:

  ```powershell
  scoop bucket add tealbase https://github.com/tealbase/scoop-bucket.git
  scoop install tealbase
  ```

  To upgrade:

  ```powershell
  scoop update tealbase
  ```
</details>

<details>
  <summary><b>Linux</b></summary>

  Available via [Homebrew](https://brew.sh) and Linux packages.

  #### via Homebrew

  To install:

  ```sh
  brew install tealbase/tap/tealbase
  ```

  To upgrade:

  ```sh
  brew upgrade tealbase
  ```

  #### via Linux packages

  Linux packages are provided in [Releases](https://github.com/tealbase/cli/releases). To install, download the `.apk`/`.deb`/`.rpm`/`.pkg.tar.zst` file depending on your package manager and run the respective commands.

  ```sh
  sudo apk add --allow-untrusted <...>.apk
  ```

  ```sh
  sudo dpkg -i <...>.deb
  ```

  ```sh
  sudo rpm -i <...>.rpm
  ```

  ```sh
  sudo pacman -U <...>.pkg.tar.zst
  ```
</details>

<details>
  <summary><b>Other Platforms</b></summary>

  You can also install the CLI via [go modules](https://go.dev/ref/mod#go-install) without the help of package managers.

  ```sh
  go install github.com/tealbase/cli@latest
  ```

  Add a symlink to the binary in `$PATH` for easier access:

  ```sh
  ln -s "$(go env GOPATH)/cli" /usr/bin/tealbase
  ```

  This works on other non-standard Linux distros.
</details>

<details>
  <summary><b>Community Maintained Packages</b></summary>

  Available via [pkgx](https://pkgx.sh/). Package script [here](https://github.com/pkgxdev/pantry/blob/main/projects/tealbase.com/cli/package.yml).
  To install in your working directory:

  ```bash
  pkgx install tealbase
  ```

  Available via [Nixpkgs](https://nixos.org/). Package script [here](https://github.com/NixOS/nixpkgs/blob/master/pkgs/development/tools/tealbase-cli/default.nix).
</details>

### Run the CLI

```bash
tealbase help
```

Or using npx:

```bash
npx tealbase help
```

## Docs

Command & config reference can be found [here](https://tealbase.com/docs/reference/cli/about).

## Breaking changes

The CLI is a WIP and we're still exploring the design, so expect a lot of breaking changes. We try to document migration steps in [Releases](https://github.com/tealbase/cli/releases). Please file an issue if these steps don't work!

## Developing

To run from source:

```sh
# Go >= 1.20
go run . help
```

