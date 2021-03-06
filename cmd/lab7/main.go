// by setting package as main, Go will compile this as an executable file.
// Any other package turns this into a library
package main

// These are your imports / libraries / frameworks
import (
	// this is Go's built-in sql library
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	// this allows us to run our web server
	"github.com/gin-gonic/gin"
	// this lets us connect to Postgres DB's
	_ "github.com/lib/pq"
	// this allows us to better format JSON responses
	"github.com/coopernurse/gorp"
)

var (
	// this is the pointer to the database we will be working with
	// this is a "global" variable (sorta kinda, but you can use it as such)
	db *sql.DB
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Holds the items that're returned for a single shelter
	// type Shelter struct {
	// 	Id int `db:"id"`
	// 	Name 	string  `db:"name"`  // <--- EDIT THESE LINES
	// 	Desc 	string	`db:"desc"` //<--- ^^^^
	// 	Phone string	`db:"phone"`
	// 	Email string 	`db:"email"`//<--- ^^^^
	// 	Url 	string	`db:"url"`
	// }
	type Shelter struct {
		Id 		int
		Name 	string
		Desc 	string
		Phone string
		Email string
		Url 	string
	}

	var errd error
	// here we want to open a connection to the database using an environemnt variable.
	// This isn't the best technique, but it is the simplest one for heroku
	db, errd := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}
  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Shelter{}, "shelter").SetKeys(true, "Id")
  // create the table. in a production system you'd generally
  // use a migration tool, or create the tables via scripts
  err2 := dbmap.CreateTablesIfNotExists()
	if err2 != nil {
		log.Fatalf("Create Table failed: %q", err2)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("html/*")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})

	router.GET("/client", func(c *gin.Context) {
		c.HTML(http.StatusOK, "client.html", nil)
	})

	router.GET("/shelter", func(c *gin.Context) {
		log.Println("============\nGetting All Shelters\n============")
		var shelters []Shelter
    _, err := dbmap.Select(&shelters, "SELECT id, name, shelter.desc, phone, email, url FROM shelter where shelter.desc != '' and name != '' and phone != '' and email != '' and url != '' LIMIT 50")
		if err != nil {
			log.Fatalf("Select failed: %q", err)
		}
    content := gin.H{}
    for k, v := range shelters {
			if &v == nil {
				content[strconv.Itoa(k)] = ""
			} else {
				content[strconv.Itoa(k)] = v
			}
    }
    c.JSON(200, content)
	})

	router.GET("/shelter/:shelter_id", func(c *gin.Context) {
		log.Println("============\nGetting Shelter\n============")
		shelter_id := c.Params.ByName("shelter_id")
		s_id, _ := strconv.Atoi(shelter_id)
		shelter := Shelter{}
		err := dbmap.SelectOne(&shelter, "SELECT id, name, shelter.desc, phone, email, url FROM shelter WHERE id=$1", s_id)
		if err != nil {
			log.Fatalf("SelectOne failed: %q", err)
		}
		content := gin.H{"name": shelter.Name, "id": shelter.Id, "desc": shelter.Desc,
											"phone": shelter.Phone, "email": shelter.Email, "url": shelter.Url}
		c.JSON(200, content)
	})

	router.POST("/shelter/:shelter_id", func(c *gin.Context) {
		log.Println("============\nUpdating Shelter\n============")
		var json Shelter
		c.Bind(&json)
		shelter := Shelter{
			Id:	json.Id,
			Desc: json.Desc,
			Phone: json.Phone,
			Email: json.Email,
			Name: json.Name,
			Url: json.Url}
		log.Printf("JSON:%q", json)
		count, err := dbmap.Update(&shelter)
		log.Printf("\nUpdated values: %q", count)
		if err != nil {
			log.Fatalf("Update failed: %q", err)
			c.JSON(500, gin.H{"result": "An error occured"})
		}
		c.JSON(200, gin.H{"result":"Success!"})
	})

	router.GET("/ping", func(c *gin.Context) {
		ping := db.Ping()
		if ping != nil {
			// our site can't handle http status codes, but I'll still put them in cause why not
			c.JSON(http.StatusOK, gin.H{"error": "true", "message": "db was not created. Check your DATABASE_URL"})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "false", "message": "db created"})
		}
	})

	//-----------------------------------------------
	//   BRITTNEY'S CLIENT VIEW CODE!!!!
	//-----------------------------------------------

	router.GET("/query1", func(c *gin.Context) {

		table := "<table class='table'><thead><tr>"
		// put your query here

		rows, err := db.Query("SELECT s.name, a.city FROM shelter s, address a WHERE s.addressId = a.id AND a.city = 'Seattle'") // <--- EDIT THIS LINE
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}
		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// declare all your RETURNED columns here
		var name string      // <--- EDIT THESE LINES
		var description string //<--- ^^^^

		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&name, &description) // <--- EDIT THIS LINE
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			table += "<tr><td>" + name + "</td><td>" + description + "</td></tr>" // <--- EDIT THIS LINE
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

	router.GET("/query2", func(c *gin.Context) {
		table := "<table class='table'><thead><tr>"
		// put your query here

		rows, err := db.Query("SELECT name, desc, url FROM shelter") // <--- EDIT THIS LINE
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}
		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// columns
		var firstName string
		var lastName string

		for rows.Next() {
			rows.Scan(&firstName, &lastName)
			// rows.Scan() // put columns here prefaced with &
			table += "<tr><td>" + firstName + "</td><td>" + lastName + "</td></tr>" // <--- EDIT THIS LINE
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

	// router.GET("/query3", func(c *gin.Context) {
	// 	table := "<table class='table'><thead><tr>"
	// 	// put your query here
	// 	rows, err := db.Query("") // <--- EDIT THIS LINE
	// 	if err != nil {
	// 		// careful about returning errors to the user!
	// 		c.AbortWithError(http.StatusInternalServerError, err)
	// 	}
	// 	// foreach loop over rows.Columns, using value
	// 	cols, _ := rows.Columns()
	// 	if len(cols) == 0 {
	// 		c.AbortWithStatus(http.StatusNoContent)
	// 	}
	// 	for _, value := range cols {
	// 		table += "<th class='text-center'>" + value + "</th>"
	// 	}
	// 	// once you've added all the columns in, close the header
	// 	table += "</thead><tbody>"
	// 	// columns
	// 	var total int
	// 	for rows.Next() {
	// 		rows.Scan(&total)
	// 		// rows.Scan() // put columns here prefaced with &
	// 		table += "<tr><td>" + strconv.Itoa(total) + "</td></tr>" // <--- EDIT THIS LINE
	// 	}
	// 	// finally, close out the body and table
	// 	table += "</tbody></table>"
	// 	c.Data(http.StatusOK, "text/html", []byte(table))
	// })

	// NO code should go after this line. it won't ever reach that point
	router.Run(":" + port)
}



/*
Example of processing a GET request

// this will run whenever someone goes to last-first-lab7.herokuapp.com/EXAMPLE
router.GET("/EXAMPLE", func(c *gin.Context) {
    // process stuff
    // run queries
    // do math
    //decide what to return
    c.JSON(http.StatusOK, gin.H{
        "key": "value"
        }) // this returns a JSON file to the requestor
    // look at https://godoc.org/github.com/gin-gonic/gin to find other return types. JSON will be the most useful for this
})

*/
