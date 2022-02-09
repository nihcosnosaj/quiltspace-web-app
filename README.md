# Quiltspace Web App 🧵 --- W.I.P. ---

This web application was developed as a project for a local quilter who wanted some sort of digital way to interact with their quilting projects. As of now, it is just a simple project management space where you can perform basic CRUD operations on a database of your quilting projects. For now, I've tailored it specifically to the art of quilting. In the future, I may refactor it to be a blank slate for managing any sort of creative hobby.

## Notes

I've decided to use a third-party ServeMux, [httprouter](https://godoc.org/github.com/julienschmidt/httprouter) instead of Go's standard net/http DefaultServeMux as I find the syntax cleaner and appropriate for my needs with this project. 

I've also decided to not use any web frameworks. Frameworks give up control and explicitness of code in favor of development speed and fast deployment. These two things are at odds with what Go is all about. Go is all about clarity, readability, and explicitness. I'm not opposed to using frameworks, but I also wanted to learn the barebone fundamentals of web development and get a lower-level understanding of what goes on underneath the hood of all these frameworks. 

## Set-Up

You will need to configure a PostgreSQL database (or I suppose any SQL relational database would work fine) with the following fields:
- Qid (an autoincrementing (serial primary key) number)
- Name 
- Style
- Pattern 

Then, you will need to create a login string for your PostgreSQL database and import it into the config/db.go file as an environment variable. I used the godotenv package to do this with the following code:
``` Go
	// load database config string as environment variable
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dotenv := os.Getenv("psqlconfig")

```

The login string should follow this format:
```
"postgres://user:password@localhost/database?sslmode=disable"
```

You may want to consider which sslmode you are using. I chose to disable it as I am only working locally for now. When I deploy, I will enable SSL to encrypt the connection between the client and server. 