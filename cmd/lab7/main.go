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
	type Shelter struct {
		name  string // <--- EDIT THESE LINES
		desc  string //<--- ^^^^
		phone string
		email string //<--- ^^^^
		url   string
	}

	var errd error
	// here we want to open a connection to the database using an environemnt variable.
	// This isn't the best technique, but it is the simplest one for heroku
	db, errd := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

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

	router.GET("/shelters", func(c *gin.Context) {
		var shelters []Shelter
		_, errd := dbmap.Select(&shelters, "SELECT * FROM public.shelter LIMIT 10")
		if errd != nil {
			log.Fatalf("Select failed", errd)
		}
		content := gin.H{}
		for k, v := range shelters {
			content[strconv.Itoa(k)] = v
		}
		c.JSON(200, content)
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
		locationInput := func(r *http.Request) { r.FormValue("location") }
		table := "<table class='table'><thead><tr>"
		// put your query here

		rows, err := db.Query("SELECT name, \"desc\" FROM shelter JOIN address ON shelter.addressId = address.id"+"JOIN state ON state.abbrev = address.stateAbbrev WHERE zip = $1 OR abbrev = $1 OR city = $1 OR"+"\"full\"= $1", locationInput) // <--- EDIT THIS LINE
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
		var name string        // <--- EDIT THESE LINES
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
	//---------------------------------------------------
	//---------------------------------------------------

	router.GET("/query2", func(c *gin.Context) {
		table := "<table class='table'><thead><tr>"
		// put your query here

		rows, err := db.Query("SELECT first, last FROM tickets t JOIN users u on t.userid = u.userid JOIN flights f on t.flightid = f.flightid JOIN locations l ON l.locationid = f.destinationid WHERE l.city = 'Atlanta' AND l.state = 'GA' GROUP BY first, last HAVING MIN(quantity) > 2") // <--- EDIT THIS LINE
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

	router.GET("/query3", func(c *gin.Context) {
		table := "<table class='table'><thead><tr>"
		// put your query here
		rows, err := db.Query("SELECT SUM(quantity) FROM tickets t JOIN flights f ON t.flightid = f.flightid WHERE f.destinationid IN (SELECT locationid FROM locations WHERE state = 'CA')") // <--- EDIT THIS LINE
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
		var total int
		for rows.Next() {
			rows.Scan(&total)
			// rows.Scan() // put columns here prefaced with &
			table += "<tr><td>" + strconv.Itoa(total) + "</td></tr>" // <--- EDIT THIS LINE
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

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
