#!/bin/bash

# No, you don't write your server in bash unless you are a maniac. This is just an example.
# Pretend that instead of a bash script taking a command line argument this is a remote API
# which takes some kind of HTTP param or something. It doesn't matter as long as the user
# can send the UUID in with their request.

UUID=$1

SERVER=localhost
PORT=8000

result=`curl -s ${SERVER}:${PORT}/verify/${UUID}`

echo "Result = '${result}'"

if [ "$result" != "OK" ]; then
  echo "mock_server: could not verify that you passed the CAPTCHA, no API for you!"
  exit
fi

echo "mock_server: verified you passed the CAPTCHA"

