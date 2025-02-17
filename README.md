# GO-SMTP
Go-smtp is a Go library designed for event-driven email sending. 

It integrates with [Kafka](https://kafka.apache.org/) to receive email sending requests and supports multiple SMTP protocols for reliable email delivery.

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

After adding environment variables, you can simply execute command `go run .`