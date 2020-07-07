// Package seeds handles the creation of initial data in the DB
package seeds

import (
	"racegex/database"
	"racegex/models"
)

// Seed will create an initial set of data in the db
func Seed() {
	seedProblems()
	seedLessons()
}

func seedProblems() {
	problems := []models.Problem{
		{ID: 1, Name: "Emails"},
		{ID: 2, Name: "Hostnames"},
		{ID: 3, Name: "IP Addresses"},
	}

	for _, problem := range problems {
		database.DBConn.Create(&problem)
	}

}

func seedLessons() {
	lessons := []models.Lesson{
		{ID: 1,
			Name:       "Introduction",
			Content:    "This is the introduction! An this is some **markdown** content.",
			Excercises: "Excercise 1",
		},
		{ID: 2,
			Name:       "Literals",
			Content:    "A literal matches the very same value you use. Example: `/a/`, `/b/`, `/2/` will _literally_ match `a`, `b`, and `2`.",
			Excercises: "Excercise 2",
		},
		{ID: 3,
			Name:       "Metacharacters",
			Content:    "Metacharacters represent sets of characters. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
		},
	}

	for i, lesson := range lessons {
		if i > 0 {
			lesson.Prev = lessons[i-1].Name
		}
		if i < len(lessons)-1 {
			lesson.Next = lessons[i+1].Name
		}

		database.DBConn.Create(&lesson)
	}

}
