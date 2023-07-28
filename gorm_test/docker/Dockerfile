# 使用官方的MySQL映像作為基本映像
FROM mysql:latest

# 複製初始化SQL腳本到容器內的指定目錄
COPY ./docker/init.sql /docker-entrypoint-initdb.d/

ENV MYSQL_ROOT_PASSWORD=123456

ENV MYSQL_USER=test
ENV MYSQL_PASSWORD=123456

# 對外暴露MySQL服務的默認端口
EXPOSE 3306
