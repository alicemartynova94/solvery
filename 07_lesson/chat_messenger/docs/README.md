# QA Notes

## Prerequisites  
Start application:  
`make run`  
  
After successful startup:  
Swagger UI: http://localhost:9000/swagger/index.html
Expected result:  
Swagger UI page is available.  
  
## Session API  
  
### Create Session  
Creates a new session for a client. The created session can subsequently be used as: X-Session-ID.
  
**Request:**  
POST http://localhost:8080/api/v1/sessions  
Content-Type: application/json
```
{}
```  
  
**Response:**  
```
{
    "id": "550e8400-e29b-41d4-a716-446655440000"
}
```  
HTTP Status: 200 OK  
  
### Delete Session  
Terminates existing session.  
  
**Request:**  
DELETE http://localhost:8080/api/v1/sessions/{id}  
```
{}
```  
  
**Expected Response:**  
```
{}
```  
HTTP Status: 204 No Content  
  
### Error Cases  
  
**Max Sessions Reached**  
Response:  
```
{
    "error": "session limit reached"
}
```  
HTTP Status: 409 Conflict  
  
**Delete Non-existent Session**  
Response:  
```
{
    "error": "session not found"
}
```  
HTTP Status: 404 Not Found  
  
**Delete Active Chat Session**
Response:  
```
{
    "error": "session owns active chats" 
}
```  
HTTP Status: 409 Conflict  
  
### gRPC API  
GRPC http://localhost:8090/messenger.v1
  
Available methods:  
- CreateSession
- DeleteSession
  
## Chat API  
All requests that require a session context must contain: X-Session-ID: <session-id>.  
  
**Session Roles in Relation to a Chat**  
- Chat owner — the session specified by chat_owner when the chat is created. 
- Chat member — a session included in the chat members list. 
  
**Session Roles Rules**  
- Chat owner cannot be removed from the chat. 
- Chat members can only be added by chat owner, removed by chat owner or chat member itself. 
- Chat owner session cannot be deleted while the session owns an active chat. 
  
**Expired Chat Rules**  
- Becomes read only chat 
- Member management is forbidden 
- Settings update is forbidden 
  
### Create Chat  
Creates a new chat for the session (chat owner). The owner must also be considered a chat member.  
  
**Max Messages Rules**  
- Changing to a lower value (than the current) is forbidden 
- Negative value is forbidden 
  
**Request:**  
POST http://localhost:8080/api/v1/chats  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
Content-Type: application/json  
```
{
    "max_messages": 100,
    "expires_at": "2026-08-10T12:00:00Z",
    "read_only": false,
    "members":[ 
    "550e8400-e29b-41d4-a716-446655440001", 
    "550e8400-e29b-41d4-a716-446655440002"
    ]
}
```  
  
**Response:**  
```
{
    "id": "660e8400-e29b-41d4-a716-446655440000"
}
```  
HTTP Status: 200 OK  
  
### Get Chat  
Gets basic chat info, including chat settings and members' session IDs.  
  
**Request:**  
GET http://localhost:8080/api/v1/chats/{id}  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000   
```
{}
```  
  
**Response:**  
```
{
    "max_messages": 100,
    "expires_at": "2026-08-10T12:00:00Z",
    "read_only": false,
    "members":[ 
    "550e8400-e29b-41d4-a716-446655440000", 
    "550e8400-e29b-41d4-a716-446655440001", 
    "550e8400-e29b-41d4-a716-446655440002"
    ],
    "created_at": "2026-08-10T10:00:00Z",
    "updated_at": "2026-08-10T10:00:00Z"
}
```  
HTTP Status: 200 OK  
  
### Update Chat   
Updates chat settings. Available for the chat owner.  
Only following fields can be updated: max_messages, expires_at, read_only.  
  
**Request:**  
PATCH http://localhost:8080/api/v1/chats/{id}  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
Content-Type: application/json  
```
{
    "max_messages": 100,
    "expires_at": "2026-08-10T12:00:00Z",
    "read_only": false
}
```  
  
**Response:**  
```
{}
```  
HTTP Status: 200 OK  
  
### Delete Chat  
Deletes an existing chat. Available for the chat owner.  
  
**Request:**  
DELETE http://localhost:8080/api/v1/chats/{id}  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
```
{}
```  
  
**Response:**  
```
{}
```  
HTTP Status: 204 No Content  
  
### Add Chat Member  
Adds an existing session to the chat.   
  
**Request:**  
POST http://localhost:8080/api/v1/chats/{id}/members  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
Content-Type: application/json  
```
{
    "session_id": "550e8400-e29b-41d4-a716-446655440001"
}
```  
  
**Response:**  
```
{}
```  
HTTP Status: 200 OK
  
### Delete Chat Member  
Removes a member's session from the chat.  
Here X-Session-ID is either the owner or a member of the chat.  
  
**Request:**  
DELETE http://localhost:8080/api/v1/chats/{id}/members  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
```
{}
```  
  
**Response:**  
```
{}
```  
HTTP Status: 204 No Content  

### Error Cases  
  
**Missing Session**  
Response:  
```
{
    "error": "session required"
}
```  
HTTP Status: 401 Unauthorized  
  
**Non-existent Session**  
Response:  
```
{
    "error": "session not found"
}
```  
HTTP Status: 401 Unauthorized  
  
**Permission Denied**  
Response:  
```
{
    "error": "forbidden"
}
```  
HTTP Status: 403 Forbidden  
  
**Chat not Found**  
Response:  
``` 
{
    "error": "chat not found"
}
```  
HTTP Status: 404 Not Found  
  
**Requester Forbidden**  
Response:  
```
{
    "error": "requester doesn't have rights for the action"
}
```  
HTTP Status: 403 Forbidden  
  
**Duplicate Member**  
Response:  
``` 
{
    "error": "member already exists"
}
```  
HTTP Status: 409 Conflict  
  
## Messages API  
Messages belong o a chat. Every message has an author, which is determined in the header (X-Session-ID).  
  
### Create Message  
Creates a new message in a chat.  
  
**Request:**  
POST http://localhost:8080/api/v1/chats/{chat_id}/messages  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
Content-Type: application/json  
```
{
    "text": "Example text."
}
```  
  
**Response:**  
```
{ 
    "id": "770e8400-e29b-41d4-a716-446655440000", 
    "chat_id": "660e8400-e29b-41d4-a716-446655440000", 
    "sender_id": "550e8400-e29b-41d4-a716-446655440000", 
    "text": "Example text.", 
    "created_at": "2026-08-10T10:00:00Z"
}
```  
HTTP Status: 200 OK  
  
### Get Messages  
Returns messages from a chat. Only chat members can retrieve messages.   
  
Constraints for pagination:  
- default limit - 50 
- offset >= 0 
- ordering by created_at 

**Request:**  
GET http://localhost:8080/api/v1/chats/{chat_id}/messages  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
```  
{}
```  
  
**Response:**  
```  
{ 
    "messages": [ 
        { 
            "id": "770e8400-e29b-41d4-a716-446655440000", 
            "chat_id": "660e8400-e29b-41d4-a716-446655440000", 
            "sender_id": "550e8400-e29b-41d4-a716-446655440000", 
            "text": "Example text.", 
            "created_at": "2026-08-10T10:00:00Z", 
            "updated_at": "2026-08-10T10:00:00Z" 
        } 
    ]
}
```  
HTTP Status: 200 OK  
  
### Delete Message  
Deletes an existing message. Only the author and chat owner can delete a message.  
  
**Request:**  
DELETE http://localhost:8080/api/v1/chats/{id}/messages/{id}  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
```  
{}
```  
  
**Response:**  
```  
{}
```  
HTTP Status: 204 No Content  
  
### Edit Message  
Updates the text of an existing message. Only the author can edit the message.  

**Request:**  
PATCH http://localhost:8080/api/v1/chats/{id}/messages/{id}  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
Content-Type: application/json  
```  
{ 
    "text": "Example update message."
}
```    
  
**Response:**  
```
{
    "id": "770e8400-e29b-41d4-a716-446655440000",
    "chat_id": "660e8400-e29b-41d4-a716-446655440000",
    "sender_id": "550e8400-e29b-41d4-a716-446655440000",
    "text": "Example update message.",
    "updated_at": "2026-08-10T10:05:00Z"
}
```  
HTTP Status: 200 OK  
  
### Error Cases  
  
**Non-existent Chat**  
Response:  
```
{
    "error": "chat not found"
}
```  
HTTP Status: 404 Not Found  
  
**Requester is not a Member**  
Response:  
```
{
   "error": "forbidden"
}
```  
HTTP Status: 401 Unauthorized  
  
**Non-existent Message**  
Response:  
```
{
   "error": "message not found"
}
```  
HTTP Status: 404 Not Found  
  
**Read-only Chat**  
Response:  
```
{
    "error": "chat is read-only"
}
```  
HTTP Status: 409 Conflict  
  
**Expired Chat**  
Response:  
```
{
    "error": "chat expired"
}
```  
HTTP Status: 409 Conflict  
  
**Empty Message**  
Response:  
``` 
{ 
    "error": "message text is required" 
}
```  
HTTP Status: 400 Bad Request  
  
**Message Length**  
Response:  
``` 
{ 
    "error": "message too long"
}
```  
HTTP Status: 400 Bad Request  