#!/bin/bash

docker run -it --net=host --name udp_sender --rm go_udp /bin/udp_sender

