# checkmate

[![CI](https://github.com/soupdevsolutions/healthchecker/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/soupdevsolutions/healthchecker/actions/workflows/ci.yml)

## Overview

Checkmate is a web application that lets you set up a list of URIs to check periodically, reporting their status based on the HTTP response status code. The status of each URI, as well as the last few checks, can be viewed in the web dashboard. The data is also saved, so it can be used for further reporting and analysis.

## Running and deploying the app

### Locally 

To run the application locally, you will need **Docker** and **PostgreSQL** installed. Having [just](https://github.com/casey/just) installed is not required, but it does make thing easier.  

First, pull the repository and navigate to the new directory:
```
git pull https://github.com/soupdevsolutions/healthchecker.git
cd ./healthchecker
```

- With **just**:
```
just start-local
```

This will initialize the database, apply the migrations and start the web server. If you already have a database set up and/or don't want to run the migrations, do not use this command.

- Without **just** / with an already running database:
```
./scripts/init_db.sh # only if you do not have a running database on port 5432
./scripts/migrate_db.sh # only if you have not applied the migrations to the running database
cd ./src && go run .
```
