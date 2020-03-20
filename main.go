package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

const port = "5010"

type User struct {
	Name  string `json:"name"`
	Age   string `json:"age"`
	Email string `json:"email"`
}

type Users *[]User

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Response struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:message`
	Status  string `json:status`
}

type AuthInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	initRouting(e)

	e.Logger.Fatal(e.Start(":" + port))
}

func initRouting(e *echo.Echo) {
	e.GET("/", index)
	e.GET("/users/:name", getUserName)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/users", saveUser)
	e.POST("/send", sendMessage)
	e.POST("/auth", auth)
}

func auth(c echo.Context) error {
	authInfo := new(AuthInfo)
	if err := c.Bind(authInfo); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, authInfo)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "hello, booker!")
}

func getUserName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	// get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	users, err := dataAdd(u.Name, u.Age, u.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func sendMessage(c echo.Context) error {
	m := new(Message)
	if err := c.Bind(m); err != nil {
		return err
	}
	r := new(Response)
	r.Name = m.Name
	r.Email = m.Email
	r.Message = m.Message
	r.Status = "success"
	return c.JSON(http.StatusOK, r)
}

func firebaseInit(ctx context.Context) (*firestore.Client, error) {
	// use a service account
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return client, nil
}

func dataAdd(name string, age string, email string) ([]*User, error) {
	ctx := context.Background()
	client, err := firebaseInit(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// add data
	_, err = client.Collection("users").Doc(name).Set(ctx, map[string]interface{}{
		"age":   age,
		"email": email,
	})
	if err != nil {
		log.Fatalf("Failed adding aloveloca: %v", err)
	}

	// read data
	allData := client.Collection("users").Documents(ctx)
	// get all documents
	docs, err := allData.GetAll()
	if err != nil {
		log.Fatalf("Failed adding getAll: %v", err)
	}

	users := make([]*User, 0)
	for _, doc := range docs {
		u := new(User)
		mapToStruct(doc.Data(), &u)
		u.Name = doc.Ref.ID
		users = append(users, u)
	}

	//close
	defer client.Close()

	return users, err
}

func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
