# go-template

<!-- put some description here -->

## Usage

<!-- instructions about how to use -->

### Command line flags

| Flag            | Environment Variable    | Description                                              | Default Value |
| --------------- | ----------------------- | -------------------------------------------------------- | ------------- |
| `-log.level`    | `APP_LOG_LEVEL`         | Application minimum log level (DEBUG, INFO, WARN, ERROR) | "INFO"        |
| `-log.format`   | `APP_LOG_FORMAT`        | Application log level format (text, json)                | "text"        |
| `-http.listen-addr` | `APP_HTTP_LISTEN_ADDR` | Address of the HTTP server                         | ":8080"       |
| `-debug.listen-addr` | `APP_DEBUG_LISTEN_ADDR` | Address of the debug HTTP server                         | ":8081"       |
| `-health.path`  | `APP_HEALTH_PATH`       | Path of the health endpoint                              | "/health"     |
| `-metrics.path` | `APP_METRICS_PATH`      | Path of the metrics endpoint                             | "/metrics"    |
| `-v`            |                         | Print version and exit                                   |               |
| `-V`            |                         | Print build information and exit                         |               |
| `-h`            |                         | Print help and exit                                      |               |

## Development

<!-- instructions about how to contribute -->
