package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app=fiber.New(fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		c.Status(fiber.StatusInternalServerError)
		return c.SendString("Error: " + err.Error())
	},
})

func TestErrorHandling(t*testing.T)  {
	app.Get("/error", func(c *fiber.Ctx) error {
		return errors.New("ups")
	})

	request:= httptest.NewRequest("GET", "/error", nil)
	response, err:= app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 500, response.StatusCode)
	bytes, err:= io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Error: ups", string(bytes))
}

func TestRoutingHelloWorld(t *testing.T){
	app.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("hello dunia")
	})
	
	request:=httptest.NewRequest("GET","/",nil)
	resp, err:=app.Test(request)
	assert.Nil(t,err)
	
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err:= io.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, "hello dunia", string(bytes))
}

func TestCtx(t *testing.T){
	app.Get("/hello",func(c *fiber.Ctx) error {
		name:=c.Query("name","Guest")
		return c.SendString("hello "+ name)
	})
	
	request:=httptest.NewRequest("GET","/hello?name=Jonathan",nil)
	resp, err:=app.Test(request)
	assert.Nil(t,err)
	
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err:= io.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, "hello Jonathan", string(bytes))

	request = httptest.NewRequest("GET","/hello",nil)
	resp, err = app.Test(request)
	assert.Nil(t,err)
	
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err = io.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, "hello Guest", string(bytes))
}

func TestHttpRequest(t *testing.T){
	app.Get("/request",func(c *fiber.Ctx) error {
		first:=c.Get("firstname")
		last:=c.Cookies("lastname")

		return c.SendString("hello " + first + " " + last)
	})
	
	request:=httptest.NewRequest("GET","/request",nil)
	request.Header.Set("firstname","Jonathan")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Farrel"})
	resp, err:=app.Test(request)
	assert.Nil(t,err)
	
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err:= io.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, "hello Jonathan Farrel", string(bytes))
}

func TestRouteParams(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error { // Removed space in URL pattern
		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("Get order " + orderId + " from user " + userId)
	})
	
	request := httptest.NewRequest("GET", "/users/jonathan/orders/10", nil)
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, "Get order 10 from user jonathan", string(bytes))
}

func TestFormReq(t *testing.T) {
	app.Post("/hello", func(c *fiber.Ctx) error { // Removed space in URL pattern
		name:= c.FormValue("name")
		return c.SendString("hello "+name)
	})
	//body request
	body:= strings.NewReader("name=Jonathan")
	request := httptest.NewRequest("POST", "/hello",body)
	request.Header.Set("Content-Type","application/x-www-form-urlencoded")
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello Jonathan", string(bytes))
}

//go:embed source/example.txt
var contohFile []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(c *fiber.Ctx) error { // Removed space in URL pattern
		file,err  := c.FormFile("file")
		if  err != nil {
			return err 
		}
		 err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}
		return c.SendString("Upload success")
	})
	//body request
	body:= new(bytes.Buffer)
	writer:= multipart.NewWriter(body)
	file, err:= writer.CreateFormFile("file", "example.txt")
	assert.Nil(t,err)
	file.Write(contohFile)
	writer.Close() 

	request := httptest.NewRequest("POST", "/upload",body)
	request.Header.Set("Content-Type",writer.FormDataContentType())
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload success", string(bytes))
}


type LoginRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestReqBody(t *testing.T) {
	app.Post("/Login", func(c *fiber.Ctx) error { // Removed space in URL pattern
		body:=c.Body()
		request:=new(LoginRequest)
		err := json.Unmarshal(body, request)
		if err != nil {
			return err
		}
		return c.SendString("hello " +request.Username) 
	})
	//body request
	body:= strings.NewReader(`{"username": "Jonathan", "password":"rahasia"}`)
	request := httptest.NewRequest("POST", "/Login",body)
	request.Header.Set("Content-Type","application/json")
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello Jonathan", string(bytes))
}

type RegisterReq struct{
	//add tag
	Username string `json:"username" xml:"username" form:username`
	Passsword string `json:"password" xml:"password" form:password`
	Name string `json:"name" xml:"name" form:name`
}

//test body parser

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(c *fiber.Ctx) error { // Removed space in URL pattern
		request:=new(RegisterReq)
		 err:= c.BodyParser(request)
		if  err!= nil {
			return nil
		}
		
		return c.SendString("register success " +request.Username) 
	})
	//body request

}

//JSON
func TestBodyParserJSON(t *testing.T) {
	TestBodyParser(t)
	//body request
	body:= strings.NewReader(`{"username": "Jonathan", "password":"rahasia","name":"Jonathan Farrel "}`)
	request := httptest.NewRequest("POST", "/register",body)
	request.Header.Set("Content-Type","application/json")
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "register success Jonathan", string(bytes))
}

//FORM

func TestBodyParserFORM(t *testing.T) {
	TestBodyParser(t)
	//body request
	body:= strings.NewReader(`username=Jonathan&password=rahasia&name=Jonathan+Farrel`)
	request := httptest.NewRequest("POST", "/register",body)
	request.Header.Set("Content-Type","application/x-www-form-urlencoded")
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "register success Jonathan", string(bytes))
}



//XML

func TestBodyParserXML(t *testing.T) {
	TestBodyParser(t)
	//body request
	body:= strings.NewReader(
	`<RegisterRequest>
		<username>Jonathan</username>
		<password>rahasia</password>
		<name>Jonathan Farrel</name>
	</RegisterRequest>
	`)
	
	request := httptest.NewRequest("POST", "/register",body)
	request.Header.Set("Content-Type","application/xml")
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "register success Jonathan", string(bytes))
}

func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(c *fiber.Ctx) error { // Removed space in URL pattern
	return c.JSON(fiber.Map{
		"username":"Jonathan",
		"name":"Jonathan Farrel",
	} )	
	})
	//body request
	// body:= strings.NewReader(`{"username": "Jonathan", "password":"rahasia"}`)
	request := httptest.NewRequest("GET", "/user",nil)
	request.Header.Set("Accept","application/json")
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t,`{"name":"Jonathan Farrel","username":"Jonathan"}`, string(bytes))
}


func TestDownloadFile(t *testing.T) {
	app.Get("/download", func(c *fiber.Ctx) error { // Removed space in URL pattern
		return c.Download("./source/example.txt","example.txt")
	})
	//body request

	request := httptest.NewRequest("GET", "/download",nil)
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	assert.Equal(t,`attachment; filename="example.txt"`, resp.Header.Get("Content-Disposition"))

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t,"this is sample file for upload", string(bytes))
}

func TestRoutingGroup(t *testing.T) {
	helloworld:= func (c *fiber.Ctx) error{
		return c.SendString("hello world")
	}
	api:=app.Group("/api")
	api.Get("/hello", helloworld) //api/hello
	api.Get("/world", helloworld) // api/world

	web:=app.Group("/web")
	web.Get("/hello", helloworld) //web/hello
	web.Get("/world", helloworld)	// web/world

	
	//body request

	request := httptest.NewRequest("GET", "/api/hello",nil)
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t,"hello world", string(bytes))
}

func TestStatic(t *testing.T) {
	app.Static("/public", "./source")
	//body request

	request := httptest.NewRequest("GET","/public/example.txt", nil)
	resp, err := app.Test(request) // Added space for consistency
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "this is sample file for upload", string(bytes))
}

func TestClient(t*testing.T)  {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)
	agent := client.Get("https://example.com/")
	//get request response
	status, response , errors := agent.String()
	assert.Nil(t, errors)
	assert.Equal(t,200, status)
	assert.Contains(t,response,"Example Domain")

	fmt.Println(response)
}







