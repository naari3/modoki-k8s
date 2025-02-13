FROM golang:1.13 as tools
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/${DOCKERIZE_VERSION}/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

ENV SQLDEF_VERSION v0.4.14
RUN wget https://github.com/k0kubun/sqldef/releases/download/${SQLDEF_VERSION}/mysqldef_linux_amd64.tar.gz \
    && tar -C /usr/local/bin -xzvf mysqldef_linux_amd64.tar.gz \
    && rm mysqldef_linux_amd64.tar.gz

FROM golang:1.13 as builder

COPY . /work
WORKDIR /work
ENV GO111MODULE=on
RUN go build -tags netgo -ldflags '-extldflags "-static"' -o /modokid /work/daemon

FROM alpine:3.9

COPY --from=tools /usr/local/bin/mysqldef /bin
COPY --from=tools /usr/local/bin/dockerize /bin
COPY ./scripts/ /bin/
COPY --from=builder /modokid /bin
COPY ./schema /schema

ENTRYPOINT [ "/bin/modokid-entrypoint.sh" ]
