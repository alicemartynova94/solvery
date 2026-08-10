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
Creates new session for a client. The created session can subsequently be used as: X-Session-ID: 550e8400-e29b-41d4-a716-446655440000.  
  
**Request:**  
POST http://localhost:8080/api/v1/sessions  
Content-Type: application/json  

```
{}
```  
  
**Example Response:**  
```
{
    "id": "550e8400-e29b-41d4-a716-446655440000"
}
```  
HTTP Status: 200 OK  
  
### Delete Session  
Terminate existing session.  

**Request:**  
DELETE http://localhost:8080/api/v1/sessions/{id}  

**Expected Response:**  
HTTP Status: 204 No Content  
  
### Error Cases  
  
**Max Sessions Reached**  
Request:  
POST http://localhost:8080/api/v1/sessions  
Content-Type: application/json
```
{}
```  
  
Expected Response:  
```
{
    "error": "session limit reached"
}
```  
HTTP Status: 409 Conflict  
  
**Delete Non-existent Session**  
Request:  
DELETE http://localhost:8080/api/v1/sessions/{id} 
  
Expected Response:  
```
{
    "error": "session not found"
}
```  
HTTP Status: 404 Not Found  
  
**Delete Active Chat Session**  

Request:  
DELETE http://localhost:8080/api/v1/sessions/{id}  
  
Expected Response:  
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
X-Session-ID identifies the session that performs the request.  
  
### Session roles
A session can have the following roles in relation to a chat:  
Chat owner — the session specified by chat_owner when the chat is created.  
Chat member — a session included in the chat members list.  
  
Chat owner cannot be removed from the chat. Chat members can only be added or removed by chat owner.  
Chat owner session cannot be deleted while the session owns an active chat.  

### Expired Chat  
- Becomes read only chat
- Member management is forbidden 
- Settings update is forbidden 

### Max Messages  
- Changing to a lower value (than the current) is forbidden 
- Negative value is forbidden 
  
### Session Related Error Cases  
**Missing session**  
Expected Response:  
```
{
    "error": "session required"
}
```
HTTP Status: 401 Unauthorized  
  
**Non-existent session**  
Expected Response:  
```
{
    "error": "session not found"
}
```  
HTTP Status: 401 Unauthorized  
  
**Permission denied**  
Expected Response:  
```
{
    "error": "forbidden"
}
```  
HTTP Status: 403 Forbidden  
  
### Create Chat  
Creates a new chat for the session (chat owner). The owner must also be considered a chat member.  
  
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
    "550e8400-e29b-41d4-a716-446655440002", 
    "550e8400-e29b-41d4-a716-446655440000"
    ]
}
```  
  
**Example Response:**  
```
{
    "id": "660e8400-e29b-41d4-a716-446655440000"
}
```  
HTTP Status: 200 OK  
  
### Get Chat  
Get basic chat info, including chat settings and member session IDs.    
  
**Request:**  
GET http://localhost:8080/api/v1/chats/{id}  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
Content-Type: application/json  
  
```
{}
```  
  
**Example Response:**  
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
HTTP Status: 200 OK  
  
### Error Cases  
  
**Chat not found**  
Example Response:
``` 
{
    "error": "chat not found"
}
```   
HTTP Status: 404 Not Found  
  
**Requester is not a member**  
Example Response: 
```
{
    "error": "forbidden"
}
```  
HTTP Status: 403 Forbidden  
  
### Update Chat   
Update chat settings. Available only for chat owner.  
  
**Request:**  
PATCH http://localhost:8080/api/v1/chats/{id}  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
Content-Type: application/json  

**Example Request:**  
Only following fields can be updated.  
```
{
    "max_messages": 100,
    "expires_at": "2026-08-10T12:00:00Z",
    "read_only": false
}
```  
  
**Example Response:**  
HTTP Status: 200 OK  

**Requester is not an owner**  
Example Response:  
```
{
    "error": "forbidden"
}
```  
  
### Delete Chat  
Delete an existing chat. Available only for chat owner.  

**Request:**  
DELETE http://localhost:8080/api/v1/chats/{id}  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  

**Expected Response:**  
HTTP Status: 204 No Content  
  
### Error Cases  
  
**Chat does not exist**  
Expected Response:  
``` 
{
    "error": "chat not found"
}
```   
HTTP Status: 404 Not Found  
  
**Requester is not the chat owner**  
Expected Response:  
```
{
    "error": "forbidden"
}
```  
HTTP Status: 403 Forbidden  
  
### Add Chat Member  
Add an existing session to the chat. Available only for chat owner.  
  
**Request:**  
POST http://localhost:8080/api/v1/chats/{id}/members  
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000  
Content-Type: application/json  
  
```
{
    "session_id": "550e8400-e29b-41d4-a716-446655440001"
}
```  
  
**Example Response:**  
HTTP Status: 200 OK  
  
### Error Cases  
  
**Chat does not exist**  
Expected Response:  
``` 
{
    "error": "chat not found"
}
```   
HTTP Status: 404 Not Found  
  
**Requester is not the chat owner**  
Expected Response:  
```
{
    "error": "forbidden"
}
```  
HTTP Status: 403 Forbidden  
  
**Duplicate Member**  
Expected Response:  
``` 
{
    "error": "member already exists"
}
```  
HTTP Status: 409 Conflict  
  
### Delete Chat Member  
Remove a session from the chat. The chat owner can remove any member, a member can only remove itself.  
Here X-Session-ID is either the owner or a member of the chat. 
  
**Request:**  
DELETE http://localhost:8080/api/v1/chats/{id}/members 
X-Session-ID: 550e8400-e29b-41d4-a716-446655440000 

**Example Request:**
```
{}
```  
  
**Example Response:**  
HTTP Status: 204 No Content  
  
### Error Cases  
**Attempt to remove the chat owner or a member by another member**  
Example Response:  
```
{
    "error": "Forbidden"
}
```  
HTTP Status: 403 Forbidden  
  
**Target session is not a member**  
Example Response:  
```
{
    "error": "member not found"
}
```  
HTTP Status: 404 Not Found  
  
**Chat does not exist**  
Example Response:  
```
{
"error": "chat not found"
}
```  
HTTP Status: 404 Not Found  

## Messages API  
