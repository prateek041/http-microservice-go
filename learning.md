# Major learning points

## Even though underlying types might be same, the methods need to be exactly the type to the one they were attached

## When we define a method on a slice of type, we can also call that method using that type itself

What is the reason of using subrouters, try testing the use of root routers.

We divide our router into multiple sub routers, that handles only certain type of requests. It helps in overall refactoring of code.

subrouter creates a subroute for a route. the subrouter is only called when the route is matched

The ServeMux type is an HTTP handler that routes incoming requests to the appropriate handler based on the request's URL. It has a list of registered routes, and it matches the request's URL path to these routes to determine which handler should be executed.

So, the serveMux is the router, and the route is a combination of a pattern and the corresponding handler function.

There are many reasons for creating subrouters, which you will get a gist of as you keep moving ahead

## Mux package

There are two things, a router and a route. A router has a list of routes, that it goes through whenever it recieves a request from the client. and call the handler registered on the route, who's criterias are closest to the incoming request.

A route is a matching between a "pattern" and what "handler" to call when the pattern matches.

A pattern can be many things, a URL path, an HTTP method, a HOST name etc. all these are handled seperately using methods attached to route struct.

Create the middleware and attach it to the (routers). It first runs the middleware present in the Use (thing) and then runs the handler.

This is not finished, refactoring in the learning docs needed.