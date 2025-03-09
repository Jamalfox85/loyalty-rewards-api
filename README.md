# loyalty-rewards-api

## start up

```
    go run main.go

    ngrok http --url=supposedly-simple-hamster.ngrok-free.app 8080 #Preferred
    or
    ngrok http 8080
```

## ngrok

Connecting to BigCommerce APIs require a public facing api. By utilizing ngrok we can expose this api publicly, allowing us to hit BC routes

```
    ngrok http --url=supposedly-simple-hamster.ngrok-free.app 8080 #Preferred
    or
    ngrok http 8080
```

exposts terminal at localhost:8080 - adjust to match host if needed
