# Simple Message API
This is a simple message rest API with some web socket feature

## How to use
- Run the app using `go run main.go` or build it first and then run the executable
- Access `http://localhost:3000/messages-board`
- Create some data using this request:
```curl
curl --location --request POST 'localhost:3000/messages' \
--header 'Content-Type: application/json' \
--data-raw '{
	"content":"some message"
}'
```
- Watch the page updated the message
- Get the message using this request:
```
curl --location --request GET 'localhost:3000/messages'
```