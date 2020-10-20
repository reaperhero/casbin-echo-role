package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reaperhero/casbin-echo-role/hander/http"
	"os"
)



func main() {
	enforcer, _ := casbin.NewEnforcer("./casbin_auth_model.conf", "./casbin_auth_policy.csv")
	e := echo.New()
	e.Use(middleware.Recover())
	//e.Use(casbin_mw.Middleware(enforcer))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: os.Stdout,
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out},"user_agent":"${user_agent}"}` + "\n",
	}))
	http.NewHTTPHandler(e,enforcer)
	e.Start(":8000")
}
