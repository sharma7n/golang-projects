psql -h localhost -c 'CREATE DATABASE donutshop;'

for migration in migrations/*.sql
do
    psql -h localhost -d donutshop -f "$migration"
done