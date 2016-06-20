FROM debian:jessie

RUN mkdir -p /opt/alert-service/conf

COPY marketwatcher-alert-service /opt/alert-service/service
COPY conf/ /opt/alert-service/conf/

CMD ["/opt/alert-service/service"]
