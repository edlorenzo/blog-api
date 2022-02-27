#!/bin/sh
chmod o+w ./database
docker-compose up
docker-compose down
rm -f ./database/blog.db
rm -f ./database/blog_test.db
