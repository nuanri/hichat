#!/bin/bash

#mysqlconn="mysql -u root -pabc -S /var/lib/mysql/mysql.sock -h localhost"
mysqlconn="mysql -u root -pabc"
olddb=hichartdb
newdb=hichatdb

$mysqlconn -e "CREATE DATABASE $newdb"
params=$($mysqlconn -N -e "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE table_schema='$olddb'")

for name in $params; do
      $mysqlconn -e "RENAME TABLE $olddb.$name to $newdb.$name";
done;

#$mysqlconn -e "DROP DATABASE $olddb"
