#!/usr/bin/env bash

createdb matcha
createuser root
sudo -u root psql -c "ALTER USER root PASSWORD 'root';"
