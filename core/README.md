# Query-Lookout Core

The core for the query-lookout

## Config

```
log_level               Log level of the core (panic, fatal, error, warn, info, debug, trace)
database_file           The file that is used for the internal sqlite database (defaults to `data.db` in your home directory)
data_source             The data source for the lookouts
base_url                Base url of the application (e.g. for the notification deeplink)
mail_from_address       From which mail address the notifcation mails are sent
mail_to_address         To which mail addresses notifcations are sent (multiple addresses separated by comma)
mail_smtp_host          SMTP host
mail_smtp_port          SMTP port
mail_username           SMTP username
mail_password           SMTP password
```

The configs are loaded in that order (details can be found [here](https://github.com/spf13/viper#why-viper))

- as environment variables with a `QL_` prefix in ALL_CAPS
- a yaml configuration file located in `~/.query-lookout`

### Example config file

```yaml
log_level: TRACE
database_file: "custom_file_name.db"
data_source: "sqlite3://data.db"
base_url: "https://karnerth.github.io/portfolio/"
mail_from_address: "test@works.at"
mail_to_address: "first@works.at,second@works.at"
mail_smtp_host: "localhost"
mail_smtp_port: "1025"
mail_username: "your@mail.net"
mail_password: "superSecret1"
```

### Example environment variables

```bash
export QL_LOG_LEVEL="INFO"
export QL_DATABASE_FILE="custom_file_name.db"
export QL_DATA_SOURCE="sqlite3://data.db"
export QL_BASE_URL="https://karnerth.github.io/portfolio/"
export QL_MAIL_FROM_ADDRESS="test@works.at"
export QL_MAIL_TO_ADDRESS="first@works.at,second@works.at"
export QL_MAIL_SMTP_HOST="localhost"
export QL_MAIL_SMTP_PORT="1025"
export QL_MAIL_USERNAME="your@mail.net"
export QL_MAIL_PASSWORD="superSecret1"
```

### GraphQl schema

The schema is split up into its different usecase subfolders.
If a change was made to them you need to call `make merge-graphql-files` to merge them
