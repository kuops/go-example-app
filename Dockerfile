FROM centos:7

WORKDIR /app
COPY build/app /app/server
CMD ["/server"]
