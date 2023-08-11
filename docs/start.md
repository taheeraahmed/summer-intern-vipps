# Getting started

## Clone the repository

```bash
git clone git@github.com:vippsas/summerstudents-backend.git
```



## Create/start a local test database

First [install docker](https://docs.docker.com/desktop/install/mac-install/) on your machine.

Use the following makefile operations, docker needs to be running for these to work.
```bash
# create the database
make compose

# shut down the database
make down

# restart the database 
make restart
```

Note: The database is wiped each time so data will not be persistent.

## Start the server locally

You may need to install some stuff, idk?

```bash
make run
```

## Testing
We have written some tests for the repository functions, but with quite low coverage so this
needs to be improved. We also did not get around to write test for the handlers, wich also
should be done with gmock.

```bash
make test
```

This will start the server at `http://127.0.0.1:8080/<some-path>`

The health endpoint will be served at `http://127.0.0.1:8081/health`

The live endpoint can be reached at `https://ece46ec4-6f9c-489b-8fe5-146a89e11635.tech-02.net/summerstudents-backend/<some-path>` 

## Use Azure Data studio

To connect to the database (local or in azure), you can use 
azure data studio.

[Download here](https://learn.microsoft.com/en-us/sql/azure-data-studio/download-azure-data-studio?view=sql-server-ver16&tabs=redhat-install%2Credhat-uninstall)


### Connect to the local database
```
Server: localhost,1436
Authentication type: SQL Login
Username: sa
Password: SuperSecret1337
Database: summerstudents-db
```


### Connect to the azure database
```
Server: sql-rafqkaoj4kdlc.database.windows.net
Authentication type: Azure active directory
Account: <select your account>
Database: rizz-and-comp-db
```