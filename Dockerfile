FROM debian:jessie

RUN mkdir -p /opt/alert-service/conf

COPY marketwatcher-alert-service /opt/alert-service/service
COPY conf/ /opt/alert-service/conf/
COPY db/ /data
COPY scripts/wait-for-it.sh /usr/local/bin
COPY scripts/start.sh /usr/local/bin

ENTRYPOINT ["/opt/alert-service/service"]
