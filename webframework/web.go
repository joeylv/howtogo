package main

import (
	"github.com/kataras/iris"
	"module"
)

// User is just a bindable object structure.
//type DirInfo struct {
//	name    string `json:"pathname"`
//	dirName string `json:"name"`
//}
//
//type Top struct {
//	dirName string `json:"pathname"`
//	tuName  string `json:"tuname"`
//	count   string `json:"count"`
//	time   string `json:"time"`
//}

func main() {
	app := iris.New()

	// Define templates using the std html/template engine.
	// Parse and load all files inside "./views" folder with ".html" file extension.
	// Reload the templates on each request (development mode).
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))

	// Register custom handler for specific http errors.
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		// .Values are used to communicate between handlers, middleware.
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}
		ctx.Writef("(Unexpected) internal server error")
	})

	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	// app.Done(func(ctx iris.Context) {})

	// Method POST: http://localhost:8080/decode
	app.Post("/decode", func(ctx iris.Context) {
		var dir module.DirInfo
		ctx.ReadJSON(&dir)
		ctx.Writef("%s %s is %d years old and comes from %s", dir.Name, dir.DirName)
	})

	// Method GET: http://localhost:8080/encode
	app.Get("/encode", func(ctx iris.Context) {
		doe := module.DirInfo{
			Name:    "Johndoe",
			DirName: "John",
		}

		ctx.JSON(doe)
	})

	// Method GET: http://localhost:8080/profile/anytypeofstring
	app.Get("/profile/{name:string}", profileByUsername)

	// Want to use a custom regex expression instead?
	// Easy: app.Get("/profile/{username:string regexp(^[a-zA-Z ]+$)}")
	//
	// If parameter type is missing then it's string which accepts anything,
	// i.e: /{paramname} it's exactly the same as /{paramname:string}.

	usersRoutes := app.Party("/users", logThisMiddleware)
	{
		// Method GET: http://localhost:8080/users/42
		usersRoutes.Get("/{id:uint64 min(1)}", getUserByID)
		// Method POST: http://localhost:8080/users/create
		usersRoutes.Post("/create", createUser)
	}

	// Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}

func logThisMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())

	// .Next is required to move forward to the chain of handlers,
	// if missing then it stops the execution at this handler.
	ctx.Next()
}

func profileByUsername(ctx iris.Context) {
	// .Params are used to get dynamic path parameters.
	username := ctx.Params().Get("username")
	ctx.ViewData("Username", username)
	// renders "./views/users/profile.html"
	// with {{ .Username }} equals to the username dynamic path parameter.
	ctx.View("users/profile.html")
}

func getUserByID(ctx iris.Context) {
	userID := ctx.Params().Get("id") // Or convert directly using: .Values().GetInt/GetUint64/GetInt64 etc...
	// your own db fetch here instead of user :=...
	user := module.DirInfo{Name: "username" + userID}

	ctx.XML(user)
}

func createUser(ctx iris.Context) {
	var dir module.DirInfo
	err := ctx.ReadForm(&dir)
	if err != nil {
		ctx.Values().Set("error", "creating user, read and parse form failed. "+err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	// renders "./views/users/create_verification.html"
	// with {{ . }} equals to the User object, i.e {{ .Username }} , {{ .Firstname}} etc...
	ctx.ViewData("", dir)
	ctx.View("users/create_verification.html")
}