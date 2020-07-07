// Package models concentrates all the types for the db and API
//
// If you want to create a table for your models, also add its
// corresponding line in the **database.Migrate()** function, and
// its corresponding seeds (if necessary) in **database.Seed()**
package models

// Problem contains the information to challenge a user with
type Problem struct {
	ID   string `json:"id"`
	Name string `gorm:"size:255"`

	Description string `json:"description"`
}

// Lesson contains one step in the learning roadmap
type Lesson struct {
	ID string `json:"id"`

	Title      string `json:"title" gorm:"size:255"`
	Subtitle   string `json:"subtitle"`
	Content    string `json:"content"`
	Excercises string `json:"excercises"`

	NextID    string `json:"nextId"`
	NextTitle string `json:"nextTitle"`
	NextLevel string `json:"nextLevel"`
	PrevID    string `json:"prevId"`
	PrevTitle string `json:"prevTitle"`
	PrevLevel string `json:"prevLevel"`

	LevelID string `json:"levelId"`
	Level   Level  `json:"level"`
}

// Level contains the information of a learning level
type Level struct {
	ID       string   `json:"id"`
	Title    string   `json:"title" gorm:"size:255"`
	Subtitle string   `json:"subtitle"`
	Lessons  []Lesson `json:"lessons"`
}
