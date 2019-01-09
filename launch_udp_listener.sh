#!/bin/bash

sudo docker run -it -p 1122:1122/udp --name udp_listener --rm go_udp /bin/udp_listener

