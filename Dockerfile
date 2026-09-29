FROM golang:1.26-trixie


RUN apt-get update && apt-get install -qy build-essential git curl ca-certificates ssh jq

ENV PATH=$PATH:/usr/local/go/bin/
ENV GOPATH=/go/
RUN go version

RUN curl -svL \
  -H "Accept: application/vnd.github+json" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/meta | jq -r '.ssh_keys | map("github.com "+.) | .[]' >> /etc/ssh/known_hosts

RUN git clone https://github.com/WiringPi/WiringPi.git && \
  cd WiringPi && \
  WIRINGPI_SUDO= ./build

RUN mkdir -p /go/src/github.com/icco/lights
WORKDIR /go/src/github.com/icco/lights

COPY .	.
RUN go mod download

RUN go build -o /go/bin/lights ./lights
RUN go build -o /go/bin/cube ./cube

CMD ["/go/bin/lights"]
