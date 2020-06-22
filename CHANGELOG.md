# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Contributing

How to contribute to the `CHANGELOG`:

- When opening a new Pull Request, describe your changes (if necessary) under
  the `[Unreleased]` section.
- Your changes should fit one of these "categories": `Added`, `Fixed`,
  `Removed`.
- If you find the appropiate category under `[Unreleased]`, append your changes
  on it as a list item. If you don't find it, create it and append your changes.
- Create configuration directory.
- Added file configuration and functions for get configuration.
- Add GORM for database connection.
- Add database file configuration.
- Add a configuration file for Deepsource.

Please check https://keepachangelog.com/en/1.0.0/#how

## [`Unreleased`]

- Start a Go Module using `gorilla/mux` to handle the http server.
- Add `air` to add hot reload to the app.
- Use `gorilla/websocket` to connect users.
- Create an initial directory structure that handles routes in `main.go`, but
  delegates all implementations of handlers to `api` and `socket` subpackages.

[`unreleased`]: https://github.com/daque-dev/sv-racegex/tree/develop
