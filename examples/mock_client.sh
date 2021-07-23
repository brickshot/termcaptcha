#!/bin/bash

SERVER=localhost
PORT=8000

UUID=`uuid`

curl --no-buffer --limit-rate 700 ${SERVER}:${PORT}/get/${UUID}

echo -n "What word did you see? "

read word

result=`curl -s ${SERVER}:${PORT}/check/${UUID}?word=${word}`

echo "Result = '${result}'"

if [ "$result" != "OK" ]; then
  echo "Wrong word, sorry."
  exit
fi

./mock_server.sh ${UUID}
