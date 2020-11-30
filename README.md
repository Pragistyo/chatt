# chatt

Back-end simple chatt app usign GO and postgreSQL \
This app using gorilla/mux as router and pgx as postgreSQL driver \
This app not include web-socket, only rest api logic \
This default app run on port 9091

# How to try This app
Install GO on your machine. \
Download Go and installation procedure -> [installGO](https://golang.org/doc/install)

## Clone this app
  using ssh 
```
git clone git@github.com:Pragistyo/chatt.git
```
  using https

```
git clone https://github.com/Pragistyo/chatt.git
```
## Install Dependencies
  from terminal on this project root folder, run command:
```sh
$ go get
```

## Setup Env

### Change .envs name
  - In root folder change file .envs to .env
  - In root/migrations folder change file .envs to .env

### Input your environtment variable
#open your .env file on

    - root folder
    - root/migrations

#change 

    - your_db_name
    - your_db_password
    - your_db_host
    - your_db_port

to your own credential

### Recommendation 
    use your cloud postgreSQL db

## DO MIGRATIONS TABLE
  from terminal on this project root folder, run command:
```sh
$ cd migrations
$ go run migrations.go
```
  then back to your root folder (hope no error happen :D )
```
$ cd ..
```

## RUN APP
from root folder
```sh
$ go run main.go
```

## ROUTES

#### go-chat Routes

| Route                                                                     |  HTTP | Description |
| ------------------------------------------------------------------------- |------ | --------------|
| `/go-chat/api/v1/create-user/`                                            | POST  | Create one user data
| `/go-chat/api/v1/login/`                                                  | POST  | retrieve one user data
| `/go-chat/api/v1/chat-room/`                                              | POST  | Create chat room between two users
| `/go-chat/api/v1/message/`                                                | POST  | User Post Message based on chat_room_name and user_id
| `/go-chat/api/v1/message-chat-room/{chat_room_name}/{opposite_user_id}/`  | GET   | User get his message, based on chat_room_name and opposite user_id. Users read all unread messages in conversation
| `/go-chat/api/v1/conversation-card/`                                     | GET   | Get List of users conversation (opposite name, last msg, unread msg) with other user 


## USING ROUTES FROM RESTFULL DEVELOPMENT TOOL (POST MAN/ INSOMNIA/ etc)
 - [download_insomnia](https://www.postman.com/downloads/)
 - [download_postman](https://insomnia.rest/download/) \
 This App input only from Request body type: Multipart From Data

### 1. CREATE USER
    url: localhost:9091/go-chat/api/v1/create-user/
    method: POST
    Body Type: multipart Form Data
    
Example Request Body 1:

    | Request Body (Multipart Form)     |            VALUE            | 
    | --------------------------------- | ----------------------------|
    | email                             |  Mourinho@TottenhamFC.co.uk |
    | name                              |  Jose Mourinho              |

Example Request Body 2:

    | Request Body (Multipart Form)     |        VALUE            | 
    | --------------------------------- | ------------------------|
    | email                             | Klopp@LiverpoolFC.co.uk |
    | name                              | Jurgen Klopp            |

Example of Response if Success:
```yaml
{
    "message":"user created",
    "new_id": 13,
    "status": 201,
}
```

### 2. LOGIN    
    url: localhost:9091/go-chat/api/v1/login/
    method: POST
    Body Type: multipart Form Data

Example Request Body 1:

    | Request Body (Multipart Form)     |          VALUE              | 
    | --------------------------------- | --------------------------- |
    | email                             |  Mourinho@TottenhamFC.co.uk |

Example of Response if Success :
```yaml
{
  "message": "success",
  "status": 200,
  "Users": {
    "id": 13,
    "Email": "Mourinho@TottenhamFC.co.uk",
    "Name": "Jose Mourinho"
}
```

### 3. CREATE CHAT ROOM
    url: localhost:9091/go-chat/api/v1/chat-room/ 
    method: POST 
    Body Type: multipart Form Data

Example Request Body 1:

    | Request Body (Multipart Form)     |          VALUE             | 
    | --------------------------------- | ---------------------------|
    | user_id                           |             13             |
    | oppose_user_id                    |              8             |
    | user_email                        | Mourinho@TottenhamFC.co.uk |
    | oppose_user_email                 | Klopp@LiverpoolFC.co.uk    |

Example of Response if Success :
```yaml
{
  "Message": "success create chat room",
  "Status": 201,
  "chat_room_name": "Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk"
}
```
Response if duplicate room:
```yaml
{
  "Message": "duplicate chat_room_name",
  "Status": 400,
}
```
### 4. POST MESSAGE 
url: localhost:9091/go-chat/api/v1/message/
method: POST
Body Type: multipart Form Data

Example Request Body 1:

    | Request Body (Multipart Form)   |                        VALUE                        | 
    | ------------------------------- | ----------------------------------------------------|
    | message                         | Not Really a good day in champions league though    |
    | chat_room_name                  | Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk  |
    | user_post_id                    |                          8                          |

Example Request Body 2:

    | Request Body (Multipart Form)   |                       VALUE                         | 
    | ------------------------------- | ----------------------------------------------------|
    | message                         | It is a good day in european league                 |
    | chat_room_name                  | Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk  |
    | user_post_id                    |                         13                          |

Example of Response if Success :
```yaml
{
  "Message": "success post message",
  "Status": 201,
  "id_message": 11
}
```
### 5. GET MESSAGE CHAT ROOM 
This api for user to retrieve all message in chat room which each of them registered. \
This end point will mark all unread opposite message whom user has conversation with as read

url : 
```
/go-chat/api/v1/message-chat-room/{chat_room_name}/{opposite_user_id}/
```

EXAMPLE url: 
```
localhost:9091/go-chat/api/v1/message-chat-room/Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk/13/
```
method: GET

In example above it is described: \
params chat_room_name 
```
Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk
```
opposite_user_id
```
13
```

Example of Response if Success :
```yaml
{
  "message": "success get messages",
  "status": 200,
  "updated_read_count": 3,
  "Users": [
    {
      "id": 5,
      "message": "You know, I 'sink' Tottenham will challange for league trophy this season",
      "senttime": "2020-11-27T06:48:25.057236Z",
      "readtime": {
        "Time": "2020-11-27T08:58:26.912476Z",
        "Valid": true
      },
      "chat_room_name": "Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk",
      "user_id": 13
    },
    {
      "id": 6,
      "message": "Ah yeah for sure. In this moment, Tottenham DO really have the rythm. BUT we will see, ",
      "senttime": "2020-11-27T06:50:06.971065Z",
      "readtime": {
        "Time": "2020-11-27T06:50:30.66142Z",
        "Valid": true
      },
      "chat_room_name": "Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk",
      "user_id": 8
    },
    {
      "id": 7,
      "message": "Just lost this champions league match, it is what it is ",
      "senttime": "2020-11-27T08:41:48.226355Z",
      "readtime": {
        "Time": "0001-01-01T00:00:00Z",
        "Valid": false
      },
      "chat_room_name": "Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk",
      "user_id": 8
    },
    {
      "id": 8,
      "message": "Either way, mo not playing full time, i could count as advantage in the long term",
      "senttime": "2020-11-27T08:58:20.074085Z",
      "readtime": {
        "Time": "0001-01-01T00:00:00Z",
        "Valid": false
      },
      "chat_room_name": "Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk",
      "user_id": 8
    }
  ]
}
```

### 6. GET LIST CONVERSATION CARD WHICH USER HAVE
url:
```
localhost:9091/go-chat/api/v1/conversation-card/{user_id}/
```
EXAMPLE url
```
localhost:9091/go-chat/api/v1/conversation-card/8/
```

Example of Response if success:
```yaml
{
  "message": "success get list conversation card",
  "status": 200,
  "Users": [
    {
      "id": 10,
      "name": "Brendan@LiverpoolFC.co.uk",
      "chat_room_name": "Klopp@LiverpoolFC.co.uk-Brendan@LiverpoolFC.co.uk",
      "unread_count": {
        "Int64": 1,
        "Valid": true
      },
      "last_msg": {
        "String": "yes, brendan we believe it",
        "Valid": true
      }
    },
    {
      "id": 12,
      "name": "Henderson@LiverpoolFC.co.uk",
      "chat_room_name": "Klopp@LiverpoolFC.co.uk-Henderson@LiverpoolFC.co.uk",
      "unread_count": {
        "Int64": 0,
        "Valid": false
      },
      "last_msg": {
        "String": "",
        "Valid": false
      }
    },
    {
      "id": 13,
      "name": "Mourinho@TottenhamFC.co.uk",
      "chat_room_name": "Mourinho@TottenhamFC.co.uk-Klopp@LiverpoolFC.co.uk",
      "unread_count": {
        "Int64": 1,
        "Valid": true
      },
      "last_msg": {
        "String": "thank you mou",
        "Valid": true
      }
    }
  ]
}
```
