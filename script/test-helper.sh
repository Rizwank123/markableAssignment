#!/usr/bin/env bash
DB_DATABASE_NAME=$(grep DB_DATABASE_NAME test.env | cut -d '=' -f2)
psql -d "${DB_DATABASE_NAME}" -c "DROP SCHEMA public CASCADE;"
psql -d "${DB_DATABASE_NAME}" -c "CREATE SCHEMA public;"