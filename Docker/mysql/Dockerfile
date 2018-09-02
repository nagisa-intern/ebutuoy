FROM mysql:5.6

ADD ./init/jihen_intern_2018-09-01.sql /tmp/jihen_intern_2018-09-01.sql
RUN "mysqld --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci"
RUN "mysql -u root --password=root test < /tmp/jihen_intern_2018-09-01.sql"
