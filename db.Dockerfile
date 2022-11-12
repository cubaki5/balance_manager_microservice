FROM mysql:8.0.23

COPY ./scripts/*.sql /docker-entrypoint-initdb.d/