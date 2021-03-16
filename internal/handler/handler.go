package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pgx "github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Conn *pgx.Conn
}

func New(dsn string) *Handler {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	return &Handler{
		Conn: conn,
	}
}

// /register?name=?&surname=?&email=?
func (h *Handler) Register(c echo.Context) error {
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")
	email := c.QueryParam("email")

	_, err := h.Conn.Exec(context.Background(), "insert into public.user (name, surname, email) values($1, $2, $3)", name, surname, email)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "user created")
}

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

// /getUser?userId=$1
func (h *Handler) GetUser(c echo.Context) error {
	userID := c.QueryParam("userId")

	var user User
	err := h.Conn.QueryRow(context.Background(), "select id, name, surname, email from public.user where id=$1", userID).
		Scan(&user.ID, &user.Name, &user.Surname, &user.Email)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, string(bytes))
}

// /createBlog?name=?&body=?&url_pic=?id_user=?
func (h *Handler) CreateBlog(c echo.Context) error {
	name := c.QueryParam("name")
	body := c.QueryParam("body")
	url_pic := c.QueryParam("url_pic")
	id_user := c.QueryParam(("id_user"))

	_, err := h.Conn.Exec(context.Background(), "insert into public.blog (name, body, url_pic, id_user) values ($1, $2, $3, $4)", name, body, url_pic, id_user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "blog created")
}

type Blog struct {
	ID        int    `json: "id"`
	Name      string `json: "name"`
	Body      string `json: "body"`
	Url_pic   string `json: "url_pic"`
	ID_user   int    `json: "id_user"`
	Name_user string
}

// /getBlog?blogId=$1
func (h *Handler) GetBlog(c echo.Context) error {
	blogID := c.QueryParam("blogId")

	var blog Blog
	err := h.Conn.QueryRow(context.Background(), "select id, name, body, url_pic, id_user from public.blog where id=$1", blogID).
		Scan(&blog.ID, &blog.Name, &blog.Body, &blog.Url_pic, &blog.ID_user)

	bytes, err := json.Marshal(blog)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, string(bytes))
}

// /getBlogOwner?blogId=$1
func (h *Handler) GetBlogOwner(c echo.Context) error {
	blogID := c.QueryParam("blogId")

	var blog Blog
	row := h.Conn.QueryRow(context.Background(), "select blog.name, \"user\".name from public.blog join public.\"user\" on blog.id_user = \"user\".id where blog.id = $1", blogID)
	err := row.Scan(&blog.Name, &blog.Name_user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	blogOwner := map[string]string{
		"user_name": blog.Name_user,
		"blog_name": blog.Name,
	}

	bytes, err := json.Marshal(blogOwner)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(bytes))
}

// /updateUser?id=?&name=?&surname=?&email=?
func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")
	email := c.QueryParam("email")

	_, err := h.Conn.Exec(context.Background(), "update public.user set name=$1, surname=$2, email=$3 where id=$4;", name, surname, email, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "user updated")
}

// /deleteUser?id=?
func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.QueryParam("id")

	_, err := h.Conn.Exec(context.Background(), "delete from public.user where id=$1;", id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "user deleted")
}
