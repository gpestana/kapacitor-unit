FROM kapacitor:1.5

COPY init.sh /init.sh
COPY kapacitor.conf /etc/kapacitor/kapacitor-unit.conf

EXPOSE 9092
CMD ["./init.sh"]
