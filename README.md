# Query-Lookout

// TODO: description

# Contents
<ul>
  <li><a href="#installation">Installation</a></li>
  <li><a href="#features">Features</a></li>
  <li><a href="#configuration">Configuration</a></li>
  <li><a href="#project-structure">Project structure</a></li>
  <li><a href="#connection-strings">Connection strings</a></li>
</ul>



## Installation
```sh
todo: clone and make target
```

or

```sh
todo: docker
```

## Features

Supported databases:
* SQLite
* PostgreSQL
* MySQL
* MSSQL (SQL Server)

## Configuration

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

## Project structure
The project consists of two parts - the core and the web. 
Inside the core folder is the backend for the application (written in Go)
and inside the web folder is the frontend (written in Svelte)

## Connection strings

Examples of valid connection strings:

* `postgresql://user:password@localhost:5432/yourDb`
* `mysql://root:password@tcp(127.0.0.1:3306)/yourDb`
* `sqlserver://user:password@localhost:1433?database=yourDb` 
