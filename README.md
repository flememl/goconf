goconf
======

A simple JSON-based configuration reader.

Disclaimer
----------

At the moment, the configuration only supports Database related variables. You will eventually be able to define your own `configuration` structure in your own packages.

Read the `config_test.json` file to know what structure the package is expecting.

Installation
------------

```$> go get github.com/flememl/goconf```

Usage
-----

Set the `GO_CONFIG_PATH` environment variable to your JSON configuration file path.

Simply import the `goconf` package wherever you need it, your configuration variables will be set automatically.

You can then start using it:

```
import (
    "database/sql"
    "fmt"
    "github.com/flememl/goconf"
    _ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
    db, err := sql.Open("mysql",
        fmt.Sprintf("%s:%s@tcp(%s:%s)/",
            goconf.Config.Database.Username,
            goconf.Config.Database.Password,
            goconf.Config.Database.Host,
            goconf.Config.Database.Port,
        ))
    return db
}
```