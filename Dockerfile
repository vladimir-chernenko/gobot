# Build Stage
FROM lacion/docker-alpine:gobuildimage AS build-stage

LABEL app="build-gobot"
LABEL REPO="https://github.com/vladimir-chernenko/gobot"

ENV GOROOT=/usr/lib/go \
    GOPATH=/gopath \
    GOBIN=/gopath/bin \
    PROJPATH=/gopath/src/github.com/vladimir-chernenko/gobot

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /gopath/src/github.com/vladimir-chernenko/gobot
WORKDIR /gopath/src/github.com/vladimir-chernenko/gobot

RUN make build-alpine

# Final Stage
FROM lacion/docker-alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/vladimir-chernenko/gobot"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/gobot/bin

WORKDIR /opt/gobot/bin

COPY --from=build-stage /gopath/src/github.com/vladimir-chernenko/gobot/bin/gobot /opt/gobot/bin/
RUN chmod +x /opt/gobot/bin/gobot

CMD /opt/gobot/bin/gobot