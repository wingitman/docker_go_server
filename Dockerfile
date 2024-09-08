FROM debian:stable-slim
COPY docker_go_server /bin/goserver

ENV PORT=8080
CMD ["/bin/goserver"]
