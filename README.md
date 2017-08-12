# go-questionnaires

[![Go Report Card](https://goreportcard.com/badge/github.com/hiromaily/go-questionnaires)](https://goreportcard.com/report/github.com/hiromaily/go-questionnaires)
[![codebeat badge](https://codebeat.co/badges/5531a607-8a8c-4768-a05b-4e4cc3dee101)](https://codebeat.co/projects/github-com-hiromaily-go-questionnaires-master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/999e25ce8e1e44a69b4bc2620b0d2743)](https://www.codacy.com/app/hiromaily2/go-questionnaires?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=hiromaily/go-questionnaires&amp;utm_campaign=Badge_Grade)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/hiromaily/go-questionnaires/master/LICENSE)

questionnaires services.  
This project was developed as company's assessment within 3 days when I applied for there.


## Directory composition
### 1.api
* For api server created by Golang

### 2.back-offcice
* For back-office server created by Node.js and React

### 3.front-offcice
* For front-office server created by Node.js and React

### 4.docker
* setting files of docker environment


## Installation on Local
####[Release]
```
$ make ecsbld
$ make ecs_push_image
$ docker-compose -f docker-compose-from-image.yml up
```

####[Development]
```
$ docker-compose build
$ docker-comppose up
```


## docker environment
| Container            | Explain                 | Port       |
|:---------------------|:------------------------|:-----------|
| question-nginx       | Nginx for reverse proxy | 8080->80   |
| question-mysql       | MySQL5.7                | 4306->3306 |
| question-api         | API server              | 8083->8083 |
| question-backoffice  | Back-Office server      | 8082->8082 |
| question-frontoffice | Front-Office server     | 8081->8081 |

## Endpoint (For localhost on Nginx)
### Front-Office Server
[http://localhost:8080/](http://localhost:8080/)

### Back-Office Server
[http://localhost:8080/admin/](http://localhost:8080/admin/)

### API Server

#### 1. Get questionnaires list 
```
[GET]
localhost:8080/api/ques
```

#### 2. Post new questionnaire 
```
[POST]
localhost:8080/api/ques
```

#### 3. Delete questionnaire 
```
[DELETE]
localhost:8080/api/ques/{id}
```

#### 4. Get answer by questionnaire_id
```
[GET]
localhost:8080/api/answer/{id}
```

#### 5. Post new answer by questionnaire_id
```
[POST]
localhost:8080/api/answer/{id}
```

