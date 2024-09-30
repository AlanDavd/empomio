#/bin/bash
#
# Run chmod u+r+x internal/dev/create_migration_id.sh

name=$1
current_date_time=$(date +%s )

echo ${current_date_time}_${name}
