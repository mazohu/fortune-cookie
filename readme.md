# Credit

Example app taken from auth0.com \(pts [1](https://auth0.com/blog/developing-golang-and-angular-apps-part-1-backend-api/) & [2](https://auth0.com/blog/developing-golang-and-angular-apps-part-2-angular-front-end/)\)

# Demoing the app
## Client
Paste the following commands in a new terminal
```cd ui \
npm install \
$env:NODE_OPTIONS="--openssl-legacy-provider" \
npm install http-server -g \
ng build --prod \
http-server dist/ui
```
## Server
Paste the following commands in a new terminal
```cd .. \
go run main.go
```