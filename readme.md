# SimpleArity

SimpleArity is a powerful command-line interface (CLI) tool designed to simplify High-Performance Computing (HPC) workflows. By integrating Singularity containers and Slurm job management, SimpleArity streamlines scientific computing tasks, making HPC environments more accessible and efficient for researchers and developers.

## Features

- **Singularity Integration**: Easily manage and deploy Singularity containers for consistent, reproducible environments.
- **HPC Environment Setup**: Quickly initialize and configure projects for HPC environments.
- **Resource Allocation**: Efficiently manage and allocate computing resources for your jobs.
- **Cross-Platform Support**: Works on both Linux and macOS systems.

## Installation

Choose the installation method that best suits your system and preferences:

 
### Manual Installation Script (macOS and Linux)

For a quick and easy installation without Homebrew, use our install script:

```bash
curl -sSL https://raw.githubusercontent.com/kemalcanbora/simplearity/main/install.sh | bash
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

Here are some basic commands to get you started with SimpleArity:

1. Initialize yaml files and environment:
   ```
   simplearity init
   ```

2. Create a convert yaml file to Dockerfile:
   ```
   simplearity create
   ```
 
3. Check job status:
   ```
   simplearity jobs
   ```

4. Available resources:
   ```
   simplearity gpu
   ```


For more detailed information on each command and its options, use the `--help` flag:

```
simplearity --help
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