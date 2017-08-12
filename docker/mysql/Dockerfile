# Dcokerfile for MySQL

FROM mysql:5.7

COPY ./conf/custom.cnf /etc/mysql/conf.d/custom.cnf
COPY ./init.d/user.sql /docker-entrypoint-initdb.d/user.sql
COPY ./init.d/dump.sql /docker-entrypoint-initdb.d/dump.sql

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=questionnaire
ENV MYSQL_USER=hiromaily
ENV MYSQL_PASSWORD=12345678

EXPOSE 3306
