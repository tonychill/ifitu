# ifitu

## Adding funciton

1. add api to the router in init_http.go

- r.app.Post("/checkout", r.handleCheckout)

2. create the handler for the route that you have just created

3. decode the response from the fiber context (c \*fiber.Ctx)

4. call the downstream services; the router is able to call what every service the router is connected to (currently finance service)

5. Return the responses as you see fit

## Creating a task in clickup

1. using the fiber lib <https://docs.gofiber.io/api/client> make a get or post request to click with the payload needed as per clickup docs/postman
2. handle the resopnse by calling the finance service (finImpl.WhateverApi...) to update stirpe and or the db.
3. return to clinet

http_handlers_task

decide if you want to call the coordinator or the service directly.

go to the proto folder and update the respective service's proto definition

compile the proto definition using the prodogen.sh script.

implement the api created by the protogen.sh script with in the actual service in /service/apis.go

implement the business logic
