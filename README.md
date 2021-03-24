# whatsapp-cli

Connect and send messages as a whatsapp client directly from command line.

## How to works

##### Install

`npm i -g whats-cli`

##### Get the QRCode on command line

`whats-cli connect-qr`

##### Get the string of QRCode on command line

`whats-cli connect`

##### Send message

`whats-cli send "number" "message"`

Number have to be without +55 prefix (8599999999) DDD + Number, +55 prefix is configured by default.

#### Dealing with bad restored sessions:

Sometimes if you want to reset the stored session you can erase the `whatsappSession.gob` file.

`rm -rf $TMPDIR/whatsappSession.gob`



