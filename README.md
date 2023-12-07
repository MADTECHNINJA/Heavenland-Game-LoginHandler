# Websocket API with Gin + Gorilla
This repository contains one basic REST API call and another websocket call

## Endpoints
/ endpoint will return base rest response
/api/test will connect it to webserver

## Controllers
Controllers contains function for websocket connection, it takes the user message and returns it to user. If message is "ping" then it returns "pong"

## Base Route
All the routes are written in Gin, specific endpoint such as /api has groups of endpoints listed under them
Currently It servers only /test, more groups can be added under it. Gorilla Websocket is used under this endpoint.

## Integration with Gorilla and Gin
Gorilla needs upgrader which will be used to upgrade connection to websocket
Gin will have default http handlers and routes, Gorilla will require to edit those
Those will be c.Writer, c.Request under gin.Context Implementations can be found under Controllers

## Package dependencies
Go is package dependent so this package is written as "bitbucket.org/heavenland/hl-game-loginhandler" More information can be found under go.mod file

## JSON Marshalling
Example added to Controllers Line 55-66
UnMarshalling added on Line 78 as function
Route will be on /api/unmarshall