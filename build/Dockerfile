FROM postgres:latest

ENV POSTGRES_USER rss_aggregator
ENV POSTGRES_PASSWORD rss_aggregator
ENV POSTGRES_DB rss_aggregator

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
