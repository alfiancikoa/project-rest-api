# Restful API Project Erajaya

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)

<br>

# Table of Content

- [Description](#description)
- [Requirements](#Requirements)
- [How to use](#how-to-use)
- [Our Feature](#Our-Feature)
- [Endpoints](#endpoints)

<br>


# Description

Project Erajaya Basic Rest-API.
Project API

<br>


# Requirements

* Visual Studio Code
* Postman
* Mysql Workbench


<br>

# High Level Architecture
<img src="https://github.com/alfiancikoa/project-rest-api/blob/main/img/HLA.jpg">

# How to use
- For More instructions details klik link <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/instruction.txt">link text</a>
- Install Go, Postman, MySQL Workbench
- Create Environment file .env
```
MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_HOST=172.17.0.2
MYSQL_PORT=3306
MYSQL_DBNAME=dbproject
MYSQL_ROOT_PASSWORD=password
```
- Clone this repository in your $PATH:
```
https://github.com/alfiancikoa/project-rest-api.git
```
* CREATE DATABASE IF NOT EXISTS `project1`;
* USE `project1`;
* Run `main.go`
```
$ go run main.go
```
* Open Postman run with your localhost, follow the routes in the Visual Studio Code folder.

<br>


# Our Feature
* CREATE: Products
* READ: Products

<br>
<br>

# Endpoints

| Method | Endpoint | Description| Authentication | Authorization
|:-----|:--------|:----------| :----------:| :----------:|
| POST  | /products | Insert New Products | No | No
| GET | /products | Get All Products List | No | No
| GET    | /products/filter/:request | Get All Products Filter by | No | No

<br>

# Filter Use

Field the filter to use
* Ascending ( Starting from A )
* Descending ( Starting from Z )
* Lower ( From lower price )
* Upper ( From higher price )

<br>
