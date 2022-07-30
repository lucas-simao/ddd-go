#!/bin/bash
set -e

echo "Running migrations..."

for file in `ls ./migrations/*.up.sql | sort -V`; do
  echo "Applying migration $file"
  psql -v "ON_ERROR_STOP=1" --file=$file
  if [ $? -ne 0 ]; then
    down=$(echo $file | sed "s/.up.sql/.down.sql/g")
    echo "Failed. Running down migration $down"
    psql -v "ON_ERROR_STOP=1" --file=$down
    if [ $? -ne 0 ]; then
      echo "Down migrations has failed. The database may be in an incosistent state"
    else
      echo "Rollback complete $down"
    fi
    exit
  fi
  echo "Applied successfully $file"
done

echo "Migration script is done"

