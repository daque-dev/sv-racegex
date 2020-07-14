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
			Subtitle:   "With regular expressions you can describe general text patterns. Find that email in the haystack!",
			Content:    "This is the introduction! An this is some **markdown** content.",
			Excercises: "Excercise 1",
			LevelID:    "beginner",
		},
		{ID: "literals",
			Title:      "Literals",
			Subtitle:   "Literals will _literally_ (yeah, it's accurate to use it here) match the character you describe. If you're look for `trouble`, you'll find `trouble`.",
			Content:    "A literal matches the very value you describe. If you're looking for `trouble`, you'll find `trouble`.",
			Excercises: "Excercise 2",
			LevelID:    "beginner",
		},
		{ID: "metacharacters",
			Title:      "Metacharacters",
			Subtitle:   "**The meat of the matter**. They will match not one character, but a group of them. This is when lazyness pays off.",
			Content:    "Metacharacters represent sets of characters. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
			LevelID:    "beginner",
		},
		{ID: "modifiers",
			Title:      "Pattern modifiers",
			Subtitle:   "Change your regular expressions behavior just by sending one of these. Go through multiple lines, ignore the casing, or find all the matched patterns.",
			Content:    "Metacharacters represent sets of characters. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
			LevelID:    "beginner",
		},
		{ID: "quantifiers",
			Title:      "Quantifiers",
			Subtitle:   "Give some steroids to your characters (or groups) by quantifying them. `?` will make them optional, `+` will allow for 1 or more, and `*` to 0 or more.",
			Content:    "Indicate how many of the previous character it should capture. Example: `/a/`, `/b/`, `/2/` will _literally_ match `a`, `b`, and `2`.",
			Excercises: "Excercise 2",
			LevelID:    "intermediate",
		},
		{ID: "group-capturing",
			Title:      "Group capturing",
			Subtitle:   "Wrap your characters in a group, so they're now another building block that you can quantify, or to do some back-referencing.",
			Content:    "Capture or non-capture the matched group. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
			LevelID:    "intermediate",
		},
		{ID: "backreferencing",
			Title:      "Backreferencing",
			Subtitle:   "I don't get enough by finding pieces of text. I want to reuse them. Morph them into something different.",
			Content:    "Capture or non-capture the matched group. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
			LevelID:    "intermediate",
		},
		{ID: "lookups",
			Title:      "Lookups",
			Subtitle:   "Sometimes, describing what you want to match is not enough. Sometimes you also need to describe the context.",
			Content:    "Capture or non-capture the matched group. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
			LevelID:    "advanced",
		},
		{ID: "performance",
			Title:      "Performance",
			Subtitle:   "You can achieve the task you want in many different ways, so it becomes necessary to discard some of them by how complex and inefficient they are.",
			Content:    "Capture or non-capture the matched group. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
			LevelID:    "advanced",
		},
		{ID: "redos",
			Title:      "REDOS",
			Subtitle:   "When writing sufficiently loose expressions turns them Evil, risking you to _Regular Expression Denial of Service_. Doesn't sound good. And it isn't.",
			Content:    "Capture or non-capture the matched group. Example: `.` will match any character, `^` will match the start of a string, and `$` the end.",
			Excercises: "Excercise 2",
			LevelID:    "advanced",
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
			Subtitle: "You are able to read and write some regular expressions."},
		{ID: "advanced",
			Title:    "Advanced",
			Subtitle: "You know how to read and understand most of the expressions you find."},
	}

	for _, level := range levels {
		database.DBConn.Create(&level)
	}
}
