FROM jrei/systemd-centos:7

COPY ReviewService.viper.yaml /etc/ReviewService/ReviewService.viper.yaml
COPY review-service /usr/local/bin/review-service

CMD ["review-service", "run"]
