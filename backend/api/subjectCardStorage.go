package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) GetUserFavoriteSubjects(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		var subjectCards []SubjectCard

		sqlString := `
			select t.* from (
				select s.id, s.name, s.category_id, c.name, (
					select count(*) from notes n
					inner join users u on u.id = ? and n.author_id = u.id
					where n.subject_id = s.id
				) posts_count
				from subjects s
				inner join categories c on c.id = s.category_id
				order by posts_count desc, s.name
				limit 5
			) t
			where t.posts_count > 0`

		err := api.db.Raw(sqlString, id).Scan(&subjectCards)
		if err != nil {
			log.Println(err)
		}

		if err := json.NewEncoder(c.Writer).Encode(subjectCards); err != nil {
			log.Printf("Cannot encode note cards to JSON, error: %v\n", err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
