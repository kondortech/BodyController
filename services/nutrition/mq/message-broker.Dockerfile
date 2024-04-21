FROM rabbitmq:3.13-management
COPY conf/rabbitmq.conf /etc/rabbitmq/
COPY conf/definitions.json /etc/rabbitmq/