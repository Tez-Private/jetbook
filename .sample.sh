#!/bin/bash

ENDPOINT=http://localhost:8080/api/v1
USERTOKEN=""

##############
# User
##############
echo "###################" && echo "# User Login" && echo "###################\n"
echo "#### POST /api/v1/users ####" && curl -X POST -H "Content-Type: application/json" -d '{"name": "test","email": "test@email.com", "pass": "pass"}' $ENDPOINT/users && echo "\n"
echo "#### GETALL /api/v1//uers ####" && curl $ENDPOINT/users && echo "\n"
echo "#### GET /api/v1/users/1 ####" && curl -X GET -H "Authorization:$USERTOKEN" -H "Userid:1" $ENDPOINT/users/1 && echo "\n"
echo "#### PATCH /api/v1/users/1 ####" && curl -X PATCH -H "Content-Type: application/json" -d '{"username": "test", "pass": "pass"}' $ENDPOINT/users/1 && echo "\n"
#echo "#### DELETE /api/v1/bizaccounts/1 ####" && curl -X DELETE $ENDPOINT/users/1 && echo "\n"

##############
# Book
##############
echo "###################" && echo "# User Login" && echo "###################\n"
echo "#### POST /api/v1/books ####" && curl -X POST -H "Content-Type: application/json" -d '{"isbn": "9784295004806"}' $ENDPOINT/books && echo "\n"