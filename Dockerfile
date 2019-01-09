FROM yuenwenl/my_go_dev:0.1

EXPOSE 1122:1122/udp

RUN apt-get update
RUN apt-get install -y curl

# build the go program
WORKDIR $GOPATH/src/my_udp
ADD src .
RUN go build -o udp_sender   udp_sender.go
RUN go build -o udp_listener udp_listener.go

RUN cp udp_sender   /bin
RUN cp udp_listener /bin

CMD ["/bin/bash"]
