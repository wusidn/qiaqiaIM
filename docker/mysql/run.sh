#!/bin/bash

docker run -it \
-p 3306:3306 \
-v $(pwd)/docker/mysql/conf.d:/etc/mysql/conf.d \
-v $(pwd)/docker/mysql/db:/usr/local/mysql/data \
-e MYSQL_ROOT_PASSWORD=wusi930208 \
--network qiaqia-net \
-name mysql \
-d mysql:8.0.16

docker run -it -p 3306:3306 -v $(pwd)/docker/mysql/conf.d:/etc/mysql/conf.d -v $(pwd)/docker/mysql/db:/usr/local/mysql/data -e MYSQL_ROOT_PASSWORD=wusi930208 --network qiaqia-net -name mysql -d mysql:8.0.16