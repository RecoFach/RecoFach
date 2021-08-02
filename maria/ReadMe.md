# MariaDB

MariaDB is used by the recommendation service as a database where all the described subjects are. 
This database is populated on startup by the `db_initialiser` service.

## Running MariaDB

`docker build -t masha .`

`docker run -p 3306:3306 masha`