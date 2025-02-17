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

## Sending email
go-smtp accepts inputs as JSON format. for sending email via smtp you should publish a json that has a format as follows through your message producers

```json
{
    "To": [
        "ensar@example.com"
    ],
    "Subject": "Sample Subject",
    "ContentType": "text/plain; charset=UTF-8",
    "Content": "The peace be upon you"
}
```