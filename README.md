# Restful API Project2 Airbnb

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)

<br>

# Table of Content

- [Description](#description)
- [Requirements](#Requirements)
- [How to use](#how-to-use)
- [Our Feature](#Our-Feature)
- [Endpoints](#endpoints)
- [Credits](#credits)

<br>


# Description

Project Base task Alterra Academy ini dibuat untuk mengimplementasikan SCRUM method using GO Language STRUCTURING , ECHO , JWT , API. 

<br>


# Requirements

* Visual Studio Code
* Postman
* Mysql Workbench


<br>


# How to use
- Install Go, Postman, MySQL Workbench
- Clone this repository in your $PATH:
```
https://github.com/armuh16/Project2-Airbnb.git
```
* CREATE DATABASE IF NOT EXISTS `alta_airbnb`;
* USE `alta_airbnb`;
* Run `main.go`
```
$ go run main.go
```
* Open Postman run with your local host, follow the routes in Visual Studio Code folder.


<br>

# Our Feature
* CREATE : User, Homestays, Reservations
* READ : User, Homestays, Reservations
* UPDATE User, Homestays
* DELETE : User, Homestays, Reservations

<br>
<br>

# Endpoints

| Method | Endpoint | Description| Authentication | Authorization
|:-----|:--------|:----------| :----------:| :----------:|
| POST  | /register | Register a new user | No | No
| POST | /login | Login existing user| No | No
|---|---|---|---|---|
| GET    | /users/:id | Get list of all user | Yes | Yes
| PUT | /users/:id | Update user profile | Yes | Yes
| DELETE | /users/:id | Delete user profile | Yes | Yes
|---|---|---|---|---|
| POST   | /homestays | Add homestays | Yes | Yes
| GET   | /homestays | Get homestays | No | No
| GET   | /homestays/:id | Get homestays by id homestays | No | No
| GET   | /homestays/type/:type | Get homestays by type feature filter | No | No
| GET   | /homestays/feature/:type | Get homestays by ex. wifi, pool, ac | No | No
| GET   | /homestays/my | Get homestays own homestay (own hosting) | Yes | Yes
| PUT | /homestays/:id | Update homestays by id homestays | Yes | Yes
| DELETE   | /homestays/:id | Delete homestays by id homestays | Yes | Yes
|---|---|---|---|---|
| GET | /reservations | Get list of all reservations | Yes | Yes
| GET | /reservations/:id | Get list of id reservations | Yes | Yes
| POST | /reservations | Add reservations | Yes | Yes
| POST | /reservations/check | Check reservations avail | No | No
| DELETE | /reservations/:id | Delete reservations by id reservations | Yes | Yes
|---|---|---|---|---|


<br>


# Credits

1. https://github.com/alfiancikoa
2. https://github.com/Nathannov24
3. https://github.com/armuh16
