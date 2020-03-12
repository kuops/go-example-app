FROM centos:7

RUN yum -y install gcc
WORKDIR /app
COPY build/app /app/server
COPY build/templates /app/templates
COPY build/database /app/database
CMD ["/app/server"]
