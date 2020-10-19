FROM gcr.io/distroless/static

COPY octail /

ENTRYPOINT ["/octail"]
