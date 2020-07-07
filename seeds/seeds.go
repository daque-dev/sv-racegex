// Package seeds handles the creation of initial data in the DB
package seeds

import (
	"racegex/database"
	"racegex/models"
)

// Seed will create an initial set of data in the db
func Seed() {
	seedProblems()
	seedLevels()
	seedLessons()
}

func seedProblems() {
	problems := []models.Problem{
		{ID: "emails", Name: "Emails"},
		{ID: "hostnames", Name: "Hostnames"},
		{ID: "ip-addresses", Name: "IP Addresses"},
	}

	for _, problem := range problems {
		database.DBConn.Create(&problem)
	}

}

func seedLessons() {
	lessons := []models.Lesson{
		{ID: "intro",
			Title:      "What is a regular expression",
			Subtitle:   "A brief introduction",
			Content:    "This is the introduction! An this is some **markdown** content.",
			Excercises: "Excercise 1",
			LevelID:    "beginner",
		},
		{ID: "literals",
			Title:      "Literals",
			Subtitle:   "a, 2, 4, Joe, 2020",
			Content:    "A literal matches the very same value you use. Example: `/a/`, `/b/`, `/2/` will _literally_ match `a`, `b`, and `2`.",
			Excercises: "Excercise 2",
			LevelID:    "beginner",
		},
		{ID: "metacharacters",
			Title:      "Metacharacters",
			Subtitle:   "., ^, $",
			Content:    "Metacharacters represent sets of characters. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
			LevelID:    "beginner",
		},
		{ID: "quantifiers",
			Title:      "Quantifiers",
			Subtitle:   "*, +, ?",
			Content:    "Indicate how many of the previous character it should capture. Example: `/a/`, `/b/`, `/2/` will _literally_ match `a`, `b`, and `2`.",
			Excercises: "Excercise 2",
			LevelID:    "intermediate",
		},
		{ID: "group-capturing",
			Title:      "Group capturing",
			Subtitle:   "(), (?:)",
			Content:    "Capture or non-capture the matched group. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
			LevelID:    "intermediate",
		},
	}

	for i, lesson := range lessons {
		if i > 0 {
			lesson.PrevID = lessons[i-1].ID
			lesson.PrevTitle = lessons[i-1].Title
			lesson.PrevLevel = lessons[i-1].LevelID
		}
		if i < len(lessons)-1 {
			lesson.NextID = lessons[i+1].ID
			lesson.NextTitle = lessons[i+1].Title
			lesson.NextLevel = lessons[i+1].LevelID
		}

		database.DBConn.Create(&lesson)
	}

}

func seedLevels() {
	levels := []models.Level{
		{ID: "beginner",
			Title:    "Beginner",
			Subtitle: "You may know what regular expressions are for, but are unfamiliar with their syntax."},
		{ID: "intermediate",
			Title:    "Intermediate",
			Subtitle: "You are able to understand and write some regular expressions."},
	}

	for _, level := range levels {
		database.DBConn.Create(&level)
	}
}
