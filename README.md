# go-dbgen

> go-dbgen is generate batabase CURD models and markdown file.

### Supported Databases 
1. `PostgresSQL`

### Quick Start
1. see help   `go-dbgen help`
2. connection database `./go-dbgen -h 127.0.0.1 -P 5432 -u username -p password -d dbname -t tablename`

### Install
```shell
git clone https://github.com/zhouxuwen/go-dbgen
```

### Parameter Description
```shell
   -h value  Database address (example: "127.0.0.1")
   -P value  port number (example: 5432)
   -u value  database username (example: "root")
   -p value  database password
   -d value  database name (example: qipaidb)
   -t value  table name (example: user) (all is all tables)
```

### Operation Command Description
```shell
NO:0 Set build directory
NO:1 Generate the table markdown document
NO:2 Generate table structure entities
NO:3 Generate CURD insert, delete, update and select
NO:4 Sets the struct mapping name
NO:5 Find or set the table name
NO:7, c, clear Clear the screen
NO:8, h, help Show help list
NO:9, q, quit Quit
```


### Other
1. [go-mygen](https://github.com/yezihack/go-mygen) generate MySQL CURD .etc File.
2. [go-bindata](https://github.com/go-bindata/go-bindata) This package converts any file into manageable Go source code.