FROM golang:1.12-alpine3.9 as builder
LABEL maintainer="Irakli Nadareishvili"

ENV PORT=3737
# Commented-out because these are defaults anyway
# ENV GOPATH=/go
# ENV PATH=${GOPATH}/bin:${PATH}
ENV APP_USER=appuser
ENV SRC_PATH=/app
ENV GIN_MODE=release
ENV GO111MODULE=on
# ENV APP_ENV=production

COPY . ${SRC_PATH}
WORKDIR ${SRC_PATH}

USER root

RUN adduser -s /bin/false -D ${APP_USER} \
 && echo "Installing git and bash support" \
 && apk update && apk upgrade \
 && apk add --no-cache bash git \
 && echo "Installing code hot reloader" \
 && go get -u github.com/cespare/reflex \
 && echo "Installing go dependencies…" \
 && go mod verify \
 && echo "Fixing permissions..." \
 && chown -R ${APP_USER}:${APP_USER} ${GOPATH} \
 && chown -R ${APP_USER}:${APP_USER} ${SRC_PATH} \
 && echo "Cleaning up installation caches to reduce image size" \
 && rm -rf /root/src /tmp/* /usr/share/man /var/cache/apk/*

USER ${APP_USER}

EXPOSE ${PORT}


FROM scratch as release
ENV APP_ENV=production
ENV PORT=3737

WORKDIR /
COPY --from=builder /go/src/app/main .
CMD ["/main"]