// Package models concentrates all the types for the db and API
//
// If you want to create a table for your models, also add its
// corresponding line in the **database.Migrate()** function, and
// its corresponding seeds (if necessary) in **database.Seed()**
package models

// Problem contains the information to challenge a user with
type Problem struct {
	ID   int    `json:"id"`
	Name string `gorm:"size:255"`

	Description string `json:"description"`
}

// Lesson contains one step in the learning roadmap
type Lesson struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"size:255"`

	Content    string `json:"content"`
	Excercises string `json:"excercises"`

	Next string `json:"next"`
	Prev string `json:"prev"`
}
