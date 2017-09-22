package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ory/dockertest"
)

var db *sql.DB

func main() {

	i := 1

	fmt.Printf("%d\n", i)
	i++
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	fmt.Printf("%d\n", i)
	i++
	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	fmt.Printf("%d\n", i)
	i++
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		fmt.Printf("retry: %d\n", i)
		i++
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	fmt.Printf("%d\n", i)
	i++
	//code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// form := url.Values{}
	// form.Set("user_id", "180")
	// url := "http://localhost:50923/user/swimlane/get"
	// url = "http://httpbin.org/post"

	// // resp, err := http.PostForm(url, form)

	// req, err := http.NewRequest("POST", url, nil) //strings.NewReader(form.Encode()))
	// req.Form = form
	// req.PostForm = form
	// req.Form.Set("user_id", "please")
	// req.PostForm.Set("user_id", "please2222")
	// dieOn(err)
	// dieOn(err, "can't make request")
	// resp, err := http.DefaultClient.Do(req)

	// dieOn(err, "can't make request")
	// b, err := ioutil.ReadAll(resp.Body)
	// dieOn(err, "can't make request")
	// fmt.Printf(string(b))
	// return
}

func dieOn(err error, msg ...string) {
	if err != nil {
		log.Fatalf("%v: %v", err, msg)
	}
}
