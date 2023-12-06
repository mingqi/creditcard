// go run ./main.go
package main

import (
	"example/hello/calculate"
	"example/hello/category"
	"example/hello/csvtolist"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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

type Filename struct {
	Year  string
	Month string
	File  string
}
type Data struct {
	Filearrey []Filename
}

func main() {
	http.HandleFunc("/hh", func(w http.ResponseWriter, r *http.Request) {
		var v Filename
		path := "/Users/shaom/codes/creditcard/pies" // replace with your directory path
		files, err := ioutil.ReadDir(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t, err := template.ParseFiles("hellotemplate.html")
		if err != nil {
			log.Fatal(err)
		}
		var data Data
		for i, file := range files {
			filename := []byte(file.Name())
			if len(filename) != 10 {
				continue
			}
			if string([]byte(file.Name())[7])+string([]byte(file.Name())[8])+string([]byte(file.Name())[9]) != "png" {
				continue
			}

			year, err := strconv.Atoi(string(filename[0:4]))
			if err != nil {
				continue
			}
			if 2000 >= year || year >= 2999 {
				continue
			}
			month, err := strconv.Atoi(string(filename[4:6]))
			if err != nil {
				continue
			}
			if month < 1 || month > 12 {
				continue
			}

			data.Filearrey = append(data.Filearrey, v)
			data.Filearrey[i].File = file.Name()
			data.Filearrey[i].Year = fmt.Sprint(year)
			data.Filearrey[i].Month = fmt.Sprint(month)
		}

		fmt.Print(data.Filearrey)
		err = t.Execute(w, data)
		fmt.Println("aaa")
		if err != nil {
			log.Fatal(err)
		}
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
