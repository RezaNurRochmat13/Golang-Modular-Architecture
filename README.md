# Research Modular Architecture Using Golang

This project is example research for how building modular architecture in Golang. 

## Table of Contents

List of all readme contents:

- [Key Features](#key-features)
- [Built With](#built-with)
- [Deployment](#deployment)
- [API Reference](#api-reference)
- [Release History](#release-history)
- [Authors](#authors)

## Key Features

List of all features/functionalities handled by this app/service:

 - Key feature A
   - Sub feature A1
   - Sub feature A2
   - Sub feature A3
 - Key feature B
   - Sub feature B1
   - Sub feature B2
 - Key feature C
 - Key feature D
 - Key feature E

## Built With
List of all core technologies used to build this app/service along with their functions, versions (if any), and link to their online pages, grouped based on development stack:

### Backend
- [Golang](https://golang.org/) - Programming Language [v1.12.4]
- [Gin Gonic](https://github.com/gin-gonic/gin) - HTTP Routing Framework for Golang [v1.3.0]
- [Echo](https://github.com/labstack/echo) - HTTP Routing Framework for Golang [v4.0.0]
- [Dep Package Management](https://github.com/golang/dep) - Dep Package Management For Golang [v0.5.1]
- [Viper Configuration Tools](https://github.com/spf13/viper) - Viper Configuration Tools [v1.3.2]
- [Go SQL Driver](https://github.com/go-sql-driver/mysql) - SQL Driver For MySQL DB [v1.4.1]
- [Assert Unit Testing](https://github.com/stretchr/testify) - Assert Testify Unit Test Tools [v1.3.0]

### Frontend
- [Angular](https://angular.io/) - JS Framework [v.7.0.0]
- [Bootstrap](http://getbootstrap.com/) - CSS Framework [v4.0.0]

### Automation Test
- For Running All Test, run command go test

### Database
- [MySQL](https://www.percona.com/) - Core Database [v5.7]

## Deployment

List all required steps to deploy this app/service in server, like environment variables, server requirements, amount of compute resources (CPU, RAM), and dependency services that will communicate with this app/service.

### Environment Variables

**MySQL Config:**

- `MYSQL_DATABASE` Name of the database using MySQL.
- `MYSQL_USER` Username for the database using MySQL.
- `MYSQL_PASSWORD` Password for the database using MySQL.
- `MYSQL_HOST` Hostname of the database server using MySQL.

**Dependency Services Config:**

- `USER_SERVICE` FQDN of User Service.
- `STUDENT_SERVICE` FQDN of Student Service.
- `NOTIF_SERVICE` FQDN of Notification Service.
- `LOGGING_SERVICE` FQDN of Logging Service.

**Other Config:**

- `LOG_LEVEL` Log level of the service. Accepted values: `debug`, `info`, or `notice`.
- `JWT_SECRET` The key that will be used to sign JSON Web Token.

## API Reference

Depending on the size of this app/service APIs, if it is small and simple enough the reference docs can be added to this README. For medium to larger app/service please use Gitlab wiki or provide a link to where the API reference docs live.

## Release History

List of all released versions of this app/service along with their version logs, sorted from newest to oldest:

- 0.2.1
  - CHANGE: Update docs (module code remains unchanged)
- 0.2.0
  - CHANGE: Remove `setDefaultXYZ()`
  - ADD: Add `init()`
- 0.1.1
  - FIX: Crash when calling `baz()`
- 0.1.0
  - The first proper release
  - CHANGE: rename `foo()` to `bar()`
- 0.0.1
  - Work in progress

## Authors
List of all people working on this app/service along with their respective roles & contact emails:

- [Dian Sigit Prastowo](mailto:bruce.wayne@uii.ac.id) - Project Manager (Team Leader)
- [Manggala Pramuditya W](mailto:billy.batson@uii.ac.id) - DevOps and Backend Engineer
- [Ahmad Haris Fahmy](mailto:diana.prince@uii.ac.id) - Backend Engineer
- [Reja Nur Rochmat](mailto:hal.jordan@uii.ac.id) - Frontend and Backend Engineer
- [Rofi Abul Hasani](mailto:barry.allen@uii.ac.id) - UI/UX and Frontend Developer
- [Maghfirah Suyuthi](mailto:clark.kent@uii.ac.id) - Frontend Developer