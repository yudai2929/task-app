package middleware

import (
	"fmt"
	"log/slog"

	"github.com/ogen-go/ogen/middleware"
)

func AccessLog() middleware.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		slog.Info(fmt.Sprintf("%s %s", req.Raw.Method, req.Raw.URL))

		res, err := next(req)
		if err != nil {
			slog.Error(fmt.Sprintf("Error %v", err))
		} else {
			slog.Info(fmt.Sprintf("Response %T", res.Type))
		}

		return res, err
	}
}
