﻿# websocket_echo_clean_architecture

## How to Use
1. Clone this repository.

2. Run the application.

```sh
go run main.go
```

3. Try it out using client application like Postman.

## How to Create a Room

1. Create a new http request POST

2. Enter the Url with this format `http://localhost:8000/ws/createRoom`

3. Input Room_Id and Room_Name with JSON format, and Send request.
```json
{
    "id":"1",
    "name":"room1"
}
```

## How to Test using Postman

1. Create a new request called `first user`.

2. Enter the Web Socket URL with this format `ws://localhost:8000/ws/joinRoom/:roomId?userId=YOUR_USERID&username=YOUR_USERNAME`.

Example: `ws://localhost:8000/ws/joinRoom/1?userId=1&username=dylee`

3. Save the request then try to connect to the Web Socket by clicking `Connect` button.

4. Create another request called `second user`.

5. Enter the Web Socket URL with this format `ws://localhost:8000/ws/joinRoom/:roomId?userId=YOUR_USERID&username=YOUR_USERNAME`.

Example: `ws://localhost:8000/ws/joinRoom/1?userId=2&username=innaka`

6. Save the request then try to connect to the Web Socket by clicking `Connect` button.

7. In first request, try to send a message to the user 2 from user 1.

```
    hallo
```

8. The message will be received in second request. In the second request, try to send a message back to the user 1.
```
    hi, user 1
```
