# R1 - ArvanCloud CLI Tool

`r1` is a command-line interface (CLI) tool designed to help you manage your infrastructure resources on **ArvanCloud**. It allows you to define the resources you want (like virtual machines) in a configuration file and then use simple commands to create, view, or delete them.

## Table of Contents

* [Features](#features)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Configuration](#configuration)
* [Usage](#usage)
    * [Defining Resources (`r1.json`)](#defining-resources-r1json)
    * [Applying Changes](#applying-changes)
    * [Checking State](#checking-state)
    * [Destroying Resources](#destroying-resources)
    * [Command Help](#command-help)
    * [Shell Autocompletion](#shell-autocompletion)
* [Example Workflow](#example-workflow)

## Features

  * Manage ArvanCloud IaaS (Infrastructure as a Service) resources.
  * Define your desired infrastructure state using a simple JSON file.
  * Apply changes to create or update resources to match your desired state.
  * View the current state of managed resources.
  * Destroy managed resources easily.

## Prerequisites

  * **Make:** You need `make` installed to build the project.
  * **Go Toolchain:** (Assuming it's a Go project based on `make` common usage) Ensure you have a recent version of Go installed.
  * **ArvanCloud Account:** You need access credentials for your ArvanCloud account.

## Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/0xAFz/r1.git
    cd r1
    ```
2.  **Build the binary:**
    ```bash
    make
    ```
    This will create the `r1` executable file in the current directory. You might want to move this file to a directory in your system's `PATH` (like `/usr/local/bin`) for easier access:
    ```bash
    # Optional: Move r1 to your PATH
    sudo mv r1 /usr/local/bin/
    ```

## Configuration

Before using `r1`, you need to provide your ArvanCloud credentials and potentially other settings.

1.  **Copy the example environment file:**
    ```bash
    cp .env.example .env
    ```
2.  **Edit the `.env` file:**
    Open the `.env` file with your preferred text editor:
    ```bash
    vim .env
    # or nano .env, code .env, etc.
    ```
3.  **Add your details:**
    Fill in the necessary values. The `.env.example` file should indicate what's needed (e.g., API keys). For example:
    ```dotenv
    ARVANCLOUD_API_KEY="apikey xyz"
    ```
    **Important:** Keep your `.env` file secure and **do not** commit it to version control (it should be listed in your `.gitignore` file). `r1` will automatically load these variables when it runs.

## Usage

`r1` works by comparing your desired state (defined in `r1.json`) with the actual resources in ArvanCloud and making changes as needed.

### Defining Resources (`r1.json`)

1.  **Create the state file:**

    ```bash
    touch r1.json
    ```

2.  **Define your resources:**
    Open `r1.json` and define the resources you want `r1` to manage. The file should contain a list of resource objects. Here's an example structure for a virtual machine:

    ```json
        [
            {
                "region": "ir-thr-si1",
                "data": {
                    "name": "r1",
                    "backup_name": null,
                    "count": 1,
                    "enable_backup": false,
                    "flavor_id": "eco-1-1-0",
                    "security_groups": [
                        {
                            "name": "e699ac81-538d-49a4-80c4"
                        }
                    ],
                    "ssh_key": true,
                    "key_name": "r1",
                    "disk_size": 25,
                    "init_script": "",
                    "ha_enabled": true,
                    "server_volumes": [],
                    "enable_ipv4": true,
                    "enable_ipv6": false,
                    "image_id": "fb7b732b-5d1f-43d9-9377-8418d7ad303f"
                }
            }
        ]
    ```

### Applying Changes

To create or update resources in ArvanCloud to match your `r1.json` file:

```bash
r1 iaas apply
```

`r1` will figure out what needs to be created or modified and perform the necessary actions.

### Checking State

To see the resources that `r1` is currently managing (based on its understanding of the state):

```bash
r1 state
```

### Destroying Resources

To remove **all** the resources defined in your `r1.json` file from ArvanCloud:

```bash
r1 iaas destroy
```

**Warning:** This command is destructive and will permanently delete the resources managed by `r1` in ArvanCloud. Be careful when using this command, especially in production environments. Data on these resources will be lost.

### Command Help

You can get help for `r1` itself or any specific command:

  * Show general help and available commands:
    ```bash
    r1 --help
    ```
  * Show help for a specific command (e.g., `iaas`):
    ```bash
    r1 iaas --help
    ```

### Shell Autocompletion

`r1` can generate autocompletion scripts for various shells (like Bash, Zsh, Fish). This makes typing commands faster and easier.

```bash
# Example for Bash:
r1 completion bash --help
r1 completion bash > /etc/bash_completion.d/r1

# Example for Zsh:
r1 completion zsh --help
r1 completion zsh > "${fpath[1]}/_r1"
# You may need to run compinit afterwards:
# autoload -U compinit && compinit
```

Follow the instructions provided by `r1 completion [your-shell] --help`.

## Example Workflow

1.  **Install `r1`** (Build from source as shown above).
2.  **Configure `r1`** by creating and editing your `.env` file with ArvanCloud credentials.
3.  **Define resources** by creating and editing `r1.json` (e.g., add a virtual machine definition).
4.  **Create the resources:** `r1 iaas apply`
5.  **Check the managed state:** `r1 state`
6.  **(Later) Modify resources:** Edit `r1.json` (e.g., change `disk_size`).
7.  **Apply modifications:** `r1 iaas apply`
8.  **(When done) Remove resources:** `r1 iaas destroy`
