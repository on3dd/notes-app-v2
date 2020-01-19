package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (api *API) GetNote(c *gin.Context) {
	// Check if the article ID is valid
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		var note Note
		api.db.Where("id = ?", id).First(&note)

		if note.Id == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		if err := json.NewEncoder(c.Writer).Encode(note); err != nil {
			log.Printf("Cannot encode note to JSON, error: %v\n", err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (api *API) GetNotes(c *gin.Context) {
	var notes []Note
	api.db.Order("id asc").Find(&notes)

	if len(notes) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := json.NewEncoder(c.Writer).Encode(notes); err != nil {
		log.Printf("Cannot encode notes to JSON, error: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func (api *API) AddNote(c *gin.Context) {
	err := c.Request.ParseMultipartForm(0)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	var note Note
	note.AuthorId, _ = strconv.Atoi(c.Request.FormValue("author"))
	note.CategoryId, _ = strconv.Atoi(c.Request.FormValue("category_id"))
	note.TeacherId, _ = strconv.Atoi(c.Request.FormValue("teacher_id"))
	note.Title = c.Request.FormValue("title")
	note.Description = c.Request.FormValue("description")

	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatalf("Error retrieving the file: %v", err)
	}
	defer file.Close()

	//sep := string(os.PathSeparator)
	sep := "/"
	path := sep + "downloads" + sep + "category-" + strconv.Itoa(note.CategoryId) + sep + "teacher-" + strconv.Itoa(note.TeacherId)

	CreateDirIfNotExist("../frontend/public" + path)

	hash := "." + RandStringBytesMaskImprSrcSB(8)
	note.Link = path + sep + handler.Filename + hash

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Cannot read the file: %v", err)
	}

	err = ioutil.WriteFile("../frontend/public"+note.Link, fileBytes, 0644)
	if err != nil {
		log.Fatalf("Cannot write to the file: %v", err)
	}

	var temp Note
	api.db.Table("notes").Select("id").Last(&temp)

	note.Id = temp.Id + 1
	note.PostedAt = time.Now()

	api.db.Create(&note)

	if err := json.NewEncoder(c.Writer).Encode(note); err != nil {
		log.Printf("Cannot encode note to JSON, error: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}