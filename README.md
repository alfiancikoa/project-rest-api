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

Project Erajaya Basic Rest-API
<br>
<p>Project API sederhana dengan menggunakan 2 enpoint untuk menambah dan mengurutkan list product, project ini dapat dijalankan di localhost
dan juga dapat dibungkus ke dalam container dengan mengguankn docker</p>

<br>


# Requirements

* Visual Studio Code
* Postman
* Mysql Workbench
* Docker


<br>

# High Level Architecture
<img src="https://github.com/alfiancikoa/project-rest-api/blob/main/img/HLA.jpg">

# How to use
- <h3>For More instructions details click link <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/instruction.txt">link text</a></h3>
- Install Go, Postman, MySQL Workbench, Docker
- Clone this repository in your $PATH:
```
https://github.com/alfiancikoa/project-rest-api.git
```
- Create Environment file (.env)
```
MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_HOST=172.17.0.2
MYSQL_PORT=3306
MYSQL_DBNAME=dbproject
MYSQL_ROOT_PASSWORD=password
```

* CREATE DATABASE IF NOT EXISTS `dbproject`;
* USE `dbproject`;
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
Example in localhost
<br>
- localhost:8080/products <br>
- localhost:8080/products/filter/new
<br>

# Filter Use
<br>
<br>
Field the filter to use on :require
* ascending ( Starting from A )
* descending ( Starting from Z )
* lower ( From lower price )
* upper ( From higher price )
* new ( Starting from latest Product )

<br>
<h3>For More instructions details click link <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/instruction.txt">link text</a></h3>
<br>
