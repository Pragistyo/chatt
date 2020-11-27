# chatt
back-end simple chatt app

## ROUTES

#### User Routes

| Route                                                                     |  HTTP | Description |
| ------------------------------------------------------------------------- |------ | --------------|
| `/go-chat/api/v1/create-user/`                                            | POST  | Create one user data
| `/go-chat/api/v1/login/`                                                  | POST  | retrieve one user data
| `/go-chat/api/v1/chat-room/`                                              | POST  | Create chat room between two users
| `/go-chat/api/v1/message/`                                                | POST  | User Post Message based on chat_room_name and user_id
| `/go-chat/api/v1/message-chat-room/{chat_room_name}/{opposite_user_id}/`  | GET   | User get his message, based on chat_room_name and opposite user_id
| `/go-chat/api/v1//conversation-card/`                                     | GET   | Get List of users conversation (opposite name, last msg, unread msg) with other user 