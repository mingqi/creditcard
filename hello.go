// go run ./main.go
package main

import (
	"example/hello/calculate"
	"example/hello/category"
	"example/hello/csvtolist"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wcharczuk/go-chart/v2"
)

type Person struct {
	Age    int
	Name   string
	Height int
}

func findPeople(name string, people []Person) Person {
	var theperson Person
	for i := 0; i < 10; i = i + 1 {
		if people[i].Name == name {
			theperson = people[i]
			break
		}
	}
	return theperson
}
func main() {

	http.HandleFunc("/hh", func(w http.ResponseWriter, r *http.Request) {
		path := "/Users/shaom/codes/creditcard/pies" // replace with your directory path
		files, err := ioutil.ReadDir(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		/*t, err := template.ParseFiles("hellotemplate.html")
		if err != nil {
			log.Fatal(err)
		}
		*/
		body := `
			<!DOCTYPE html>
		<html lang="en">
		  <head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<meta http-equiv="X-UA-Compatible" content="ie=edge" />
			<title>Document</title>
		  </head>
		  <body>
			<form
			  enctype="multipart/form-data"
			  action="/upload"
			  method="post"
			>
			  <input type="file" name="myFile" /></br>
			  year:<input type="text" id="year" name="year" /></br>
      		  month:<input type="text" id="month" name="month" /></br>
			  <input type="submit" value="upload" />
			  <ul>`

		for _, file := range files {
			body = body + `<li><a href=/upload/image?path=` + file.Name() + `>` + file.Name() + `</a></li>`
		}
		body = body + `
				</ul>
			</form>
	 	</body>
	</html>
		`
		fmt.Fprintf(w, body)
		/*err = t.Execute(os.Stdout, data)
		if err != nil {
			log.Fatal(err)
		}
		*/
	})
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/upload/image", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("path")
		absPath, err := filepath.Abs("pies")
		if err != nil {
			fmt.Println(err)
			return
		}

		buf, err := ioutil.ReadFile(absPath + "/" + name)
		if err != nil {
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Type", "image/png")

		w.Write(buf)
	})
	fmt.Println("-----------------------------")
	http.ListenAndServe(":8081", nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	year := r.FormValue("year")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	month := r.FormValue("month")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	err = ioutil.WriteFile("uploadedfile", fileBytes, 999999999)
	if err != nil {
		fmt.Println(err)
	}
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	txList, err := csvtolist.Tolist(dir)
	category_output := category.Tocategory(txList)
	calculate_output := calculate.Calculate(category_output)
	var values []chart.Value
	for key, value := range calculate_output {
		values = append(values, chart.Value{Label: key, Value: float64(value)})
	}
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: values,
	}
	dir, err = os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	f, _ := os.Create(dir + year + month + ".png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
