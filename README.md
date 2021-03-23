# whats-api

This is a REST API client that handles [jsdaniell/whats-cli](https://github.com/jsdaniell/whats-cli) to servers on HTTP Requests, connect and send messages on Whatsapp using HTTP requests.

## How to use

- **Clone the repository**

`git clone https://github.com/jsdaniell/whats-api.git`

- **Download dependencies**

`go mod download`

- **Run server locally**

`go run main.go`

### On backgroud

Basically this is a wrapper for [jsdaniell/whats-cli](https://github.com/jsdaniell/whats-cli) program that's allow to get the QRCode to connect and send messages using REST API, the binary of whats-cli is downloaded and, the commands is runned locally inside the requests controllers.

## Testing / Requesting

The API's endpoints until now are:

**GET**: https://localhost:9000/getQrCode <br/>
returns the QRCode PNG on the response.

GET: https://localhost:9000/disconnect <br/>
disconnect the actually logged session on the server.

POST: https://localhost:9000/send <br/>
send a message to some number.
```json
{
 "number": "558599999999",
 "message": "Message"
}
```





