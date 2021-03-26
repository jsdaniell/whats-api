# whatsapp-api

Connect and send messages as a whatsapp client sending REST requests, **this is only for experimental use**.

## How to use locally

**Requirements:**

- Have the go installed on your machine.

##### Download the repository

`git clone https://github.com/jsdaniell/whats-api.git`

##### Get dependencies

`go mod download || go mod vendor`

##### Running

`go run main.go`

##### Testing the requests

GET: https://localhost:9000/getQrCode <br>
returns the QRCode PNG on the response. <br>

`curl --location --request GET 'http://localhost:9000/getQrCode' \
 --header 'Authorization: anyIDString'`

GET: https://localhost:9000/disconnect <br>
disconnect the actually logged session on the server. <br>

`curl --location --request GET 'http://localhost:9000/disconnect' \
 --header 'Authorization: anyIDString'`

POST: https://localhost:9000/sendMessage <br>
send a message to some number. <br>

`curl --location --request POST 'http://localhost:9000/sendMessage' \
 --header 'Authorization: anyIDString' \
 --header 'Content-Type: application/json' \
 --data-raw '{
 	"number": "558599999999",
 	"message": "message"
 }'`
 
 ## How to use on Heroku Instance
 
 **Requirements:**
 
 - Have the [heroku cli](https://devcenter.heroku.com/articles/heroku-cli) installed on your machine and logged account.
 
 ##### Download the repository
 
 `git clone https://github.com/jsdaniell/whats-api.git`
 
 ##### Create your project on your heroku account
 
 `heroku create {name-of-your-project}`
 
 ##### Deploying to heroku
 
 `heroku container:push -a {name-of-your-project} web` <br>
 
 `heroku container:release -a {name-of-your-project} web`
 
 ##### Testing
 
 `curl https://{name-of-your-project}.herokuapp.com/`
 
 *The response must be: Server Running...*
 
 See logs with `heroku logs --tail` on the root folder of project.



