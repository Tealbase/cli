## tealbase-init

Initialize configurations for Tealbase local development.

A `tealbase/config.toml` file is created in your current working directory. This configuration is specific to each local project.

> You may override the directory path by specifying the `TEALBASE_WORKDIR` environment variable or `--workdir` flag.

In addition to `config.toml`, the `tealbase` directory may also contain other Tealbase objects, such as `migrations`, `functions`, `tests`, etc.
