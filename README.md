# GO-SMTP
go-smtp is a tool that sends email through different smtp protocols

It uses distributed messaging systems to trigger the action to send mails whenever an event occured.

As distributed messaging system, go-smtp uses [Kafka](https://kafka.apache.org/).

## Running Application
Before running application, make sure that you set up these parameters with `.env` file

```
SMTP_PORT=587
SMTP_HOST=mail.example.com
SMTP_USER=user@example.com
SMTP_PASSWD=<Your_Password_Here>
KAFKA_HOST=localhost
KAFKA_PORT=9092
KAFKA_BROKERS=localhost:9092
KAFKA_TOPIC=mail
KAFKA_PARTITION=0
KAFKA_GROUP_ID=go-smtp
```