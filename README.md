# Restful API Project1 Group 3 E-Commerces

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)


# Table of Content

- [Description](#description)
- [Requirements](#Requirements)
- [How to use](#how-to-use)
- [Our Feature](#Our-Feature)
- [Structuring](#Structuring)
- [Endpoints](#endpoints)
- [Credits](#credits)


# Description

Project Base task Alterra Academy ini dibuat untuk mengimplementasikan SCRUM method using GO Language STRUCTURING , ECHO , JWT , API. 

# Requirements

* Visual Studio Code
* Postman
* Mysql Workbench

# How to use
- Install Go, Postman, MySQL Workbench
- Clone this repository in your $PATH:
```
https://github.com/armuh16/Project2-Airbnb.git
```
* CREATE DATABASE IF NOT EXISTS `project1_kelompok3`;
* USE `project1_kelompok3`;
* Run `main.go`
```
$ go run main.go
```
* Open Postman run with your local host, follow the routes in Visual Studio Code folder.


# Our Feature
* CREATE : User, 
* READ : User, 
* UPDATE User, 
* DELETE : User, 


# Structuring
```
📦project1_kelompok3
 ┣ 📂.vscode
 ┃   ┗ 📜settings.json
 ┣ 📂config
 ┃   ┗ 📜config.go
 ┣ 📂constants
 ┃   ┗ 📜constant.go
 ┣ 📂controllers
 ┃   ┗ 📜orderController.go
 ┃   ┗ 📜productController.go
 ┃   ┗ 📜productController_test.go
 ┃   ┗ 📜shoppingCartController.go
 ┃   ┗ 📜shoppingCartController_test.go
 ┃   ┗ 📜userController.go
 ┣ 📂lib
 ┃   ┗ 📂database
 ┃     ┗ 📜order.go
 ┃     ┗ 📜product.go
 ┃     ┗ 📜shoppingCart.go
 ┃     ┗ 📜user.go
 ┃   ┗ 📂response
 ┃     ┗ 📜response.go
 ┣ 📂middlewares
 ┃   ┗ 📜logMiddleware.go
 ┃   ┗ 📜middlewares.go
 ┣ 📂models
 ┃   ┗ 📜address.go
 ┃   ┗ 📜order_detail.go
 ┃   ┗ 📜orders.go
 ┃   ┗ 📜payment_methods.go
 ┃   ┗ 📜products.go
 ┃   ┗ 📜shopping_carts.go
 ┃   ┗ 📜users.go
 ┣ 📂routes
 ┃   ┗ 📜route.go
 ┣ 📜.env
 ┣ 📜.gitignore
 ┣ 📜cover.out
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┣ 📜main.go
 ┣ 📜profile.cov
 ┗ 📜README.MD
```

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
| POST   | /products | Add products | Yes | Yes
| GET   | /products | Get products list | No | No
| GET   | /products/:id | Get products by id product | No | No
| GET   | /products/users | Get products by userid | Yes | Yes
| PUT | /products/:id | Update products by id products | Yes | Yes
| DELETE   | /products/:id | Delete products by id products | Yes | Yes
|---|---|---|---|---|
| GET | /shopping_carts | Get list of all shopping carts | Yes | Yes
| POST | /shopping_carts | Add shopping carts | Yes | Yes
| PUT | /shopping_carts/:id | Update shopping carts by id products | Yes | Yes
| DELETE | /shopping_carts/:id | Delete shopping_cart by id products | Yes | Yes
|---|---|---|---|---|
| POST | /orders | Add orders | Yes | Yes
| POST | /orders | Add deatil orders | Yes | Yes
| GET | /orders | Get list of all orders| Yes | Yes
| GET | /history | Get list of all history| Yes | Yes
| GET | /cancel | Get list of all cancel| Yes | Yes
|---|---|---|---|---|

<br>


# Credits

1. https://github.com/alfynf
2. https://github.com/Nathannov24
3. https://github.com/armuh16
