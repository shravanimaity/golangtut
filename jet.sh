#!/bin/bash

jet \
  -source=postgres \
  -host=dpg-d1ibu4ali9vc73fpqjt0-a.oregon-postgres.render.com \
  -port=5432 \
  -user=bootcamp_db_yj7k_user \
  -password=7uKulyPCAXk4pvvoLndAJb8mo0vExXYE \
  -dbname=bootcamp_db_yj7k \
  -schema=public \
  -sslmode=require \
  -path=./.gen


