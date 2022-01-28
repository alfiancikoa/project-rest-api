# Restful API Project Erajaya

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)

<br>

# Table of Content

- [Description](#description)
- [Requirements](#Requirements)
- [How to use](#how-to-use)
- [Start Docker](#Getting-Start-With-Docker)
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
<h5>For More instructions details click link <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/instruction.txt">instruction</a></h5>

- Clone this repository in your $PATH:
```
$ git clone https://github.com/alfiancikoa/project-rest-api.git
```
<h3>If you are using port from your localhost</h3>

- Create Environment file (.env) and fill with
  
```
MYSQL_USER=root           // your user MySQL 
MYSQL_PASSWORD=password   // your password MySQL
MYSQL_HOST=localhost      // your IP local default localhost
MYSQL_PORT=3306           // default port from MySQL
MYSQL_DBNAME=dbproject    // database name
```

- Create Database If Not Exist On the Terminal


```
# login mysql
$ sudo mysql -u root -p

# on menu mysql
mysql> create database <db-name>;

# to show all databases list
mysql> show databases;

# to exit
mysql> exit;
```
- Running the Program
```
$ go run main.go
```
# Getting Start With Docker

Common commands
```
# pull image from container registry
docker pull <image-name>

# login to container registry
docker login -u <user-name>

# build image
docker build -t <image-name>[:tag] .

# push image to container registry
docker push <image-name>

# create and run container
docker run -p <host-port>:<container-port> -e <env-name>=<env-value> -v <host-volume>:<container-volume> --name <container-name> <image-name>

# run existing container
docker start <container-name>
docker container start <container-name>

# stop running container
docker stop <container-name>
docker container stop <container-name>

# remove container
docker rm <container-name>
docker container rm <container-name>

# show images
docker image list

# show container
docker container list
docker ps
docker ps -a

# access/run command in a container
docker exec -it <container-name> <command>
docker exec -it mysql bash
```
<br>

# Using Docker

<br>
If you are using docker-desktop, the containers can access host os by using host.docker.internal name.
Otherwise, you can use default host IP address: 172.17.0.1
<br>
```
# create and run dockerMySQL container
docker run -p 3307:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=dbproject --name dockerContainer -d mysql:latest
```
Set Environtmen on file .env with your configuration
```
MYSQL_USER=root           // your user MySQL 
MYSQL_PASSWORD=password   // make sure this is same with before command (MYSQL_ROOT_PASSWORD=password)
MYSQL_HOST=172.17.0.1     // default IP Docker
MYSQL_PORT=3307           // changing port from 3306 -> 3307
MYSQL_DBNAME=dbproject    // make sure this is same with before command (MYSQL_DATABASE=dbproject)
```
Make sure your IP of Docker Container of MySQL
```
docker inspect -f '{{.Name}} - {{.NetworkSettings.IPAddress }}' $(docker ps -aq)
```
Build The Image and then Run the Container
```
# Build the image
docker build -t nama-image:latest .

# create and run appContainer
docker run --name nameContainerAPI -p 8080:8080 nama-image:latest
```


* CREATE DATABASE IF NOT EXISTS `dbproject`;
* USE `dbproject`;
* Run `main.go`
```
$ go run main.go
```
* Open Postman run with your localhost, follow the routes in the Visual Studio Code folder.
* <h4>more instruction details click link <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/instruction.txt">instruction</a></h4>

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

```
- localhost:8080/products
- localhost:8080/products/filter/new
```

<br>

# Filter Use
<br>

Field the filter to use on :require
* ascending ( Starting from A )
* descending ( Starting from Z )
* lower ( From lower price )
* upper ( From higher price )
* new ( Starting from latest Product )

<br>
<h3>For More instructions details click link <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/instruction.txt">instructions</a></h3>
<br>
<h4>For Unit Testing Coverage Report <a href="https://github.com/alfiancikoa/project-rest-api/blob/main/img/Coverage%20Report%20-%20.png">coverage</a></h4>
<br>
