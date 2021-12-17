package hello

import (
	"context"
)

//encore:api public path=/hello/:name
func World(ctx context.Context, name string) (*Response, error) {
	msg := "Hello, " + name + "!"
	return &Response{Message: msg}, nil
}

type Response struct {
	Message string
}

// ==================================================================

// Encore comes with a built-in development dashboard for
// exploring your API, viewing documentation, debugging with
// distributed tracing, and more. Visit your API URL in the browser:
//
//     http://localhost:4060
//

// ==================================================================

// Next steps
//
// 1. Deploy your application to the cloud with a single command:
//
//     git push encore
//
// 2. To continue exploring Encore, check out one of these topics:
//
//    Building a Slack bot:  https://encore.dev/docs/tutorials/slack-bot
//    Building a REST API:   https://encore.dev/docs/tutorials/rest-api
//    Using SQL databases:   https://encore.dev/docs/develop/sql-database
//    Authenticating users:  https://encore.dev/docs/develop/auth
