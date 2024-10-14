# SimpleArity

SimpleArity is a powerful command-line interface (CLI) tool designed to simplify High-Performance Computing (HPC) workflows. By integrating Singularity containers and Slurm job management, SimpleArity streamlines scientific computing tasks, making HPC environments more accessible and efficient for researchers and developers.

## Features

- **Singularity Integration**: Easily manage and deploy Singularity containers for consistent, reproducible environments.
- **HPC Environment Setup**: Quickly initialize and configure projects for HPC environments.
- **Resource Allocation**: Efficiently manage and allocate computing resources for your jobs.
- **Cross-Platform Support**: Works on both Linux and macOS systems.
- **Environment Configuration**: Uses a `simplearity.env`  and `simplearity.yaml` file for easy environment setup and management.

## Installation

Choose the installation method that best suits your system and preferences:

### 1) Script
0. Copy `install.sh` to your local directory this files under [released](released) directory.
1. chmod +x install.sh
2. ./install.sh

### 2) Manual Installation Script (macOS and Linux)

For a quick and easy installation without Homebrew, use our install script:

```bash
curl -sSL https://raw.githubusercontent.com/kemalcanbora/simplearity/refs/heads/main/released/install.sh | bash
```

This script will download the appropriate version for your system and set up SimpleArity in your home directory.

### Manual Download and Installation

If you prefer to manually download and install:

1. Go to the [Releases](https://github.com/kemalcanbora/simplearity/releases) page.
2. Download the appropriate version for your operating system and architecture.
3. Extract the archive:
   ```bash
   tar -xzf simplearity_<OS>_<ARCH>.tar.gz
   ```
4. Move the binary to a directory in your PATH:
   ```bash
   mv simplearity /usr/local/bin/
   ```

## Usage

Here are the basic commands to get you started with SimpleArity:

1. Initialize the environment and configuration:
   ```
   simplearity init
   ```
   This command creates a `simplearity.env`  and `simplearity.yaml` file with necessary configuration.

2. Deploy your job:
   ```
   simplearity deploy
   ```
   This command uses the configuration from `simplearity.env` and `simplearity.yaml` to deploy your job to the HPC environment.

For more detailed information on each command and its options, use the `--help` flag:

```
simplearity --help
```

## Configuration

SimpleArity uses a `simplearity.env` file for configuration. This file should contain the following information:

- HPC Username
- Docker Hub Username
- Image Name
- Job Name
- Memory Allocation
- Partition
- CPUs Per Task

You can create and edit this file manually, or use the `simplearity init` command to set it up interactively.

## Yaml Configuration

SimpleArity uses a `simplearity.yaml` file for configuration. This file should contain the following information:

## Example Yaml Configuration
Example 1:
```
image:
  base: python:3.9-slim
  packages:
    - numpy
  environment:
    - PYTHONUNBUFFERED=1
    - DEBUG=False

install:
  - pip install --upgrade pip

data:
  - path: /path/to/local/dataset
    mount: /data/dataset

code:
  - path: /path/to/local/script.py
    dest: /app/script.py

run:
  command: python /app/script.py
  args:
    - --input /data/dataset
    - --output /data/results
 ```  
Example 2:
```
image:
  base: python:3.9-slim
  packages:
    - numpy
  environment:
    - PYTHONUNBUFFERED=1
    - DEBUG=False

install:
  - pip install --upgrade pip

run:
  command: ls -l
```

## Contributing

We welcome contributions to SimpleArity! If you have suggestions for improvements or encounter any issues, please feel free to:

- Open an issue
- Submit a pull request
- Contact the maintainers

Please read our [Contributing Guide](CONTRIBUTING.md) for more details on how to contribute to this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any problems or have questions about using SimpleArity in your HPC environment, please:

1. Check the [documentation](https://github.com/kemalcanbora/simplearity/wiki) (if available)
2. Look through [existing issues](https://github.com/kemalcanbora/simplearity/issues) on GitHub
3. Open a new issue if your problem or question isn't already addressed

## Acknowledgements

We'd like to thank all contributors and users of SimpleArity. Your support and feedback drive the continuous improvement of this tool.

---

Developed with 🐭 by [Kemalcan Bora](https://github.com/kemalcanbora)