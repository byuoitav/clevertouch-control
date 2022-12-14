FROM grc.io/distroless/static

ARG NAME

COPY ${NAME} /app

ENTRYPOINT ["/app"]