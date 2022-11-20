# Tutorial: Accessing a relational database

This tutorial shows how to access a relational database from a Go program. It uses the [database/sql](https://golang.org/pkg/database/sql/) package to access a [MySQL](https://www.mysql.com/) database.

## MySQL

The MySQL database is a relational database management system (RDBMS). For this tutorial, we will use the [MySQL Community Server](https://dev.mysql.com/downloads/mysql/).

### Install MySQL

Download the MySQL Community Server from [https://dev.mysql.com/downloads/mysql/](https://dev.mysql.com/downloads/mysql/).

### Create a database

Connect to the MySQL server:

```bash
mysql -u root -p
```

Create a database:

```sql
CREATE DATABASE recordings;
```

Use the database:

```sql
USE recordings;
```

For create a table, we use the script [create_table.sql](create-tables.sql).

## Execute the program

To execute the program, we use the script [run.sh](run.sh).

```bash
./run.sh
```
