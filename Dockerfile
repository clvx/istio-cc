FROM alpine:latest
COPY src/bin/cb /
ENTRYPOINT ["/cb"]
