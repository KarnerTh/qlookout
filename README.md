<p align="center">
  <img src="https://github.com/KarnerTh/query-lookout/assets/22556363/1057ff28-c90e-4013-b09a-9a03f86a5860" alt="Query-Lookout logo"/>
</p>

# Query-Lookout
> ⚠️ Project is under development - features and apis are subject to change

Query Lookout is a small and simple tool designed to simplify and enhance
the way you monitor and manage your database content. As data grows in
complexity and volume, keeping a vigilant eye on changes becomes crucial.
Query-Lookout allows you to define rules, alerting you to relevant
modifications or data constellations in your database.


Want to see what has changed? Take a look at
the [Changelog](https://github.com/KarnerTh/query-lookout/blob/main/CHANGELOG.md)

## How does it work 
1. Create a new **lookout**
    1. Define a display name
    1. Configure a cron expression how often it should run (details [here](https://github.com/robfig/cron#background---cron-spec-format))
    1. Define the SQL query that should be run
    1. Configure desired notification channels (in app notification is on by default)
1. Define one or more **rules** that validate the result of the lookout
    1. Column that should be checked
    1. Row from which the value should be used (counting starts with `0`)
    1. Column type (text, int, float)
    1. Rule type (exact value, should be null, greater/less than, between)
    1. Expected value


https://github.com/KarnerTh/query-lookout/assets/22556363/e9d9c451-ba07-4eaf-908a-80a1a1c2b174



# Contents
<ul>
  <li><a href="#features">Features</a></li>
  <li><a href="#getting-started">Getting started</a></li>
  <li><a href="#installation">Installation</a></li>
  <li><a href="#supported-databases">Supported databases</a></li>
  <li><a href="#configuration">Configuration</a></li>
</ul>

# Features
- **Rule-based Monitoring:**
Create customizable rules to monitor specific changes or data constellations
in your database. Define conditions that matter most to you and receive
timely alerts when those conditions are met.

- **Alert Notifications:**
Receive notifications in real-time through various channels, such as email,
in app notifications or local notifications (more coming). Stay informed
about important updates without the need to constantly check the database.

- **Easy Configuration:**
Set up rules effortlessly with an intuitive configuration interface.
No need for complex queries or scripts – Query Lookout streamlines the
process, making monitoring accessible to everyone.

- **Database Compatibility:**
Query Lookout supports a wide range of databases, ensuring flexibility and
compatibility with popular database management systems.

- **Single Binary:**
Everything you need in one single binary - the core and web ui.
Get up and running in a few steps - locally or whereever you want.

## Getting started
1. Install query-lookout (details [here](#installation))
1. Create the configuration file `~/.query-lookout` (details [here](#configuration))
    1. Important config is `data_source`, which defines the connection string
    to your database
1. Run `qlookout`
1. Open [http://localhost:63000/](http://localhost:63000/)
1. Enjoy

## Installation
```sh
go install github.com/KarnerTh/query-lookout@latest
```


## Supported databases
| Database   | Example Connection-String                                  |
| ---------- | ---------------------------------------------------------- |
| SQLite3    | `sqlite3://your-file.db`                                   |
| PostgreSQL | `postgresql://user:password@localhost:5432/yourDb`         |
| MySQL      | `mysql://root:password@tcp(127.0.0.1:3306)/yourDb`         |
| MSSQL      | `sqlserver://user:password@localhost:1433?database=yourDb` |

## Configuration
```
# required
data_source             The data source for the lookouts

# mail config
mail_from_address       From which mail address the notifcation mails are sent
mail_to_address         To which mail addresses notifcations are sent (multiple addresses separated by comma)
mail_smtp_host          SMTP host
mail_smtp_port          SMTP port
mail_username           SMTP username
mail_password           SMTP password

# optional
log_level               Log level of the core (error, warn, info, debug)
database_file           The file that is used for the internal sqlite database (defaults to `data.db` in your home directory)
base_url                Base url of the application (e.g. for the notification deeplink)
```

The configs are loaded in that order (details can be found [here](https://github.com/spf13/viper#why-viper))

- as environment variables with a `QL_` prefix in ALL_CAPS
- a yaml configuration file located in `~/.query-lookout`

### Example config file
```yaml
log_level: INFO
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

