docker run \
--name sensor_zmq_sub \
-e ZMQ_SUB_ENDPOINT="tcp://docker.host.internal:5555" \
-e TOPIC="sensor" \
-e DB_DRIVER="mysql" \
-e DS_NAME="root:Admin123@tcp(docker.host.internal:3307)/go_fleet" \
-e MYSQL_HOST="docker.host.internal" \
-e MYSQL_PASSWORD="Admin123" \
-e MYSQL_DB="go_fleet" \
-e MYSQL_USER="root" \
-e MYSQL_PORT="3307" \
ghcr.io/dragks/gofleet/sensor_zmq_pub:main
