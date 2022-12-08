docker run \
--name sensor_zmq_pub \
-p 5555:5555 \
-e SENSOR_ID="s0" \
-e TOPIC="sensor" \
-e ZMQ_PUB_ENDPOINT="tcp://*:5555" \
ghcr.io/dragks/gofleet/sensor_zmq_pub:main
