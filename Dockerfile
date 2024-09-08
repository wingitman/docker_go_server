ENV PORT=$PORT
FROM debian:stable-slim
COPY docker_go_server /bin/goserver

CMD ["/bin/goserver"]
