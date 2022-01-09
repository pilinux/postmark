# postmark

[![Go Report Card](https://goreportcard.com/badge/github.com/pilinux/postmark)][01]
[![CodeFactor](https://www.codefactor.io/repository/github/pilinux/postmark/badge)][02]
[![codebeat badge](https://codebeat.co/badges/cd6b8c1c-c682-4535-9c15-e0c15124838d)][03]
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)][04]

## Real-time notifications from [Postmark][11] [webhooks][12]

&#9889; Receive notifications via webhooks

&#9889; Save events in MySQL / PostgreSQL / SQLite

### Best use-case

&#9889; Handle repeated abuses on the application-side before
sending transactional emails to a fake/bad email address

### Features

Version: `0.1.1`

| Message type  | Supported          |
| ------------- | ------------------ |
| Transactional | :white_check_mark: |
| Inbound       | :x:                |
| Broadcasts    | :x:                |

### Setup

- You need a MySQL / PostgreSQL instance, or you can use SQLite
- Download the binary `go install github.com/pilinux/postmark@latest`
- Create a free project on [Sentry][13] to track errors
  - Choose `Go` as the platform
  - Save the `DNS` (format: `https://secret_code@abc.ingest.sentry.io/secret_number`)
- Change the filename from `.env.sample` to `.env`
- Update the following variables in the `.env` file (DO NOT DELETE other non-used variables)
  - `APP_PORT`
  - `SentryDSN`
  - `USERNAME`
  - `PASSWORD`
  - `DBDRIVER`
  - `DBUSER`
  - `DBPASS`
  - `DBNAME`
  - `DBHOST`
  - `DBPORT`
  - `DBTIMEZONE`
- Execute the binary file. It will automatically create a new table
  `postmark_outbounds` and migrate the database.
- On Postmark,
  - select your server and browse the `transactional` message stream
  - select `Add webhook`
  - add the following information:
    - Webhook URL: `http(s)://<your_ip_or_domain:port>/webhooks/v1/outbound-events`
      If you run the application behind a reverse proxy (NGINX / Apache) and configure a
      domain, then the URL will be `http(s)://<your_domain>/webhooks/v1/outbound-events`
    - add `Basic auth credentials`
    - select the events Postmark should forward to your server

### `postmark_outbounds` database

| Name          | Comment                                                               |
| ------------- | --------------------------------------------------------------------- |
| `id`          | primary key                                                           |
| `created_at`  | data creation time in the database                                    |
| `updated_at`  |                                                                       |
| `deleted_at`  |                                                                       |
| `record_type` | Delivery / Bounce / SpamComplaint / Open / Click / SubscriptionChange |
| `type`        | HardBounce / SpamComplaint                                            |
| `type_code`   | 1 (for bounce) / 100001 (for spam complaint) / 0 (others)             |
| `message_id`  | Postmark message ID                                                   |
| `tag`         | user-defined tag                                                      |
| `from`        | from email                                                            |
| `to`          | destination email                                                     |
| `event_at`    | timestamp from Postmark                                               |
| `server_id`   | Postmark server ID                                                    |

## License

&#169; piLinux 2022

Released under the [MIT license][04]

This project is built with [GoREST][21].

[01]: https://goreportcard.com/report/github.com/pilinux/postmark
[02]: https://www.codefactor.io/repository/github/pilinux/postmark
[03]: https://codebeat.co/projects/github-com-pilinux-postmark-main
[04]: LICENSE
[11]: https://postmarkapp.com
[12]: https://postmarkapp.com/developer/webhooks/webhooks-overview
[13]: https://sentry.io
[21]: https://github.com/piLinux/GoREST
