[![Stories in Ready](https://badge.waffle.io/devopsdays/gather-flag.svg?label=ready&title=Ready)](http://waffle.io/devopsdays/gather-flag) [![Stories in Progress](https://badge.waffle.io/devopsdays/gather-flag.svg?label=in%progress&title=In%20Progress)](http://waffle.io/devopsdays/gather-flag) [![Needs Review](https://badge.waffle.io/devopsdays/gather-flag.svg?label=needs-review&title=Needs%20Review)](http://waffle.io/devopsdays/gather-flag)
[![Build Status](https://travis-ci.org/devopsdays/gather-flag.svg?branch=master)](https://travis-ci.org/devopsdays/gather-flag)
[![Coverage Status](https://coveralls.io/repos/github/devopsdays/gather-flag/badge.svg?branch=master)](https://coveralls.io/github/devopsdays/gather-flag?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/devopsdays/gather-flag)](https://goreportcard.com/report/github.com/devopsdays/gather-flag)
[![license](https://img.shields.io/github/license/devopsdays/gather-flag.svg)]()

You can see progress on tasks at http://waffle.io/devopsdays/gather-flag

[![Throughput Graph](https://graphs.waffle.io/devopsdays/gather-flag/throughput.svg)](https://waffle.io/devopsdays/gather-flag/metrics)
# gather-flag

`gather-flag` is a web application for organizing open spaces, built with :heart: by [mattstratton](https://github.com/mattstratton) in [Go](https://golang.org/).

Requirements
===========

* Docker 1.12
* Docker Compose 1.8

Starting services
==============================

```
docker-compose up -d
```

Stopping services
==============================

```
docker-compose stop
```

Including new changes
==============================

If you need change some source code you can deploy it typing:

```
docker-compose build
```

That being said, things *should* reload on the fly. 

## Documentation
### Topic Service
This service is used to get information about a topic. It provides the topic title, the description, and other details.

*Routes:*

* GET - /topics : Get all topics
* POST - /topics : Create topic
* GET - /topics/{id} : Get topic by id
* DELETE - /topics/{id} : Remove topic by id

### Account Service
This service is used to manage user accounts, specifically around login, etc.

*Routes*

* GET - /accounts : Get all accounts
* POST - /accounts: Create  account
* GET - /accounts/{id} : Get account by id
* DELETE - /accounts/{id} : Remove account by id


### Room Service
This service manages rooms, including who has access to them, etc

*Routes*

* GET - /rooms : Get all rooms
* POST - /rooms : Create new room
* GET - /rooms/{i} : Get room by id
* DELETE - /rooms/{i} : Remove room by id
