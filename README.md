# go-questionnaires
questionnaires services.


## Directory composition
### 1.api
* For api server created Golang

### 2.back-offcice
* For back-office server created Node.js and React

### 3.front-offcice
* For front-office server created Node.js and React

### 4.docker
* setting files of docker environment for development

### 5.docker2
* setting files of docker environment for release


## Installation on Local
```
$ ./docker2/create_containers.sh
 or
$ cd ./docker;./docker-create.sh
```

## docker environment
| Container            | Explain                 | Port       |
|:---------------------|:------------------------|:-----------|
| question-nginx       | Nginx for reverse proxy | 8080->80   |
| question-mysql       | MySQL5.7                | 4306->3306 |
| question-api         | API servver             | 8083->8083 |
| question-backoffice  | Back-Office server      | 8082->8082 |
| question-frontoffice | Front-Office server     | 8082->8082 |

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

