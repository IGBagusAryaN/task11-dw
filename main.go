package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/form-add-project", formAddProject)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/testimonial", testimonial)
	e.POST("/add-project", addProject)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func formAddProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	nameProject := c.FormValue("inputProjectName")
	startDate := c.FormValue("inputStartDate")
	endDate := c.FormValue("inputEndDate")
	description := c.FormValue("inputDesc")
	// checkbox1 := c.FormValue("inputCheckboxHtml")
	// checkbox2 := c.FormValue("inputCheckboxCss")
	// checkbox3 := c.FormValue("inputCheckboxReact")
	// checkbox4 := c.FormValue("inputCheckboxJs")


	println("nameProject : " + nameProject)
	println("startDate : " + startDate)
	println("startEnd : " + endDate)
	println("description : " + description)
	// println("checkbox1 : " + checkbox1)
	// println("checkbox2 : " + checkbox2)
	// println("checkbox3 : " + checkbox3)
	// println("checkbox4 : " + checkbox4)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id":      id,
		"Title":   "Game Web Store",
		"Duration" : "1 bulan",
		"Content": "Masuki alam semesta virtual yang memikat di mana para pemain game dari segala usia dan preferensi dapat menemukan pelarian yang sempurna. Toko web kami adalah gerbang ke koleksi game yang dirangkai dengan teliti, mulai dari petualangan penuh aksi yang mendebarkan hingga teka-teki yang membingungkan, simulasi olahraga yang menggetarkan hati, dan epik bermain peran yang mendalam. Dengan katalog yang luas yang mencakup permata indie dan judul blockbuster, toko web kami menawarkan beragam keajaiban gaming yang akan membuat Anda ketagihan berjam-jam.",
	}

	var tmpl, err = template.ParseFiles("views/project-detail1.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}
