#!/bin/sh
JSON_STRING='window.configs = { \
    "VUE_APP_BACKEND_PORT":"'"${VUE_APP_BACKEND_PORT}"'", \
    "VUE_APP_BACKEND_HOST":"'"${VUE_APP_BACKEND_HOST}"'" \
  }'
sed -i "s@// CONFIGURATION_PLACEHOLDER@${JSON_STRING}@" /usr/share/nginx/html/index.html
exec "$@"