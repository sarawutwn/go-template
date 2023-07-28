package QuestionRepository

import (
	"fmt"
	"go-template/database"
	"time"

	"github.com/google/uuid"
)

func CreateQuestion(title string, branch_type string) (*string, error) {
	db := database.DB
	query := `
		INSERT INTO question(
			question_id, 
			title, 
			branch_type, 
			created_at, 
			updated_at
		) VALUES ($1, $2, $3, $4, $5) RETURNING question_id
	`
	var ID string
	newUUID := uuid.New()
	err := db.QueryRow(query, newUUID.String(), title, branch_type, time.Now(), time.Now()).Scan(&ID)
	fmt.Print(err)
	if err != nil {
		return nil, err
	}
	return &ID, nil
}

func CreateDescription(DescriptionText string, DescriptionType string, DescriptionTrueString string, DescriptionFalseString string, Priority int64, QuestionID string) (*string, error) {
	db := database.DB
	query := `
		INSERT INTO question_description(
			question_description_id, 
			description_text, 
			description_type, 
			description_true_string, 
			description_false_string,
			priority,
			created_at, 
			updated_at,
			question_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING question_description_id
	`
	var ID string
	newUUID := uuid.New()
	err := db.QueryRow(query, newUUID.String(), DescriptionText, DescriptionType, DescriptionTrueString, DescriptionFalseString, Priority, time.Now(), time.Now(), QuestionID).Scan(&ID)
	if err != nil {
		return nil, err
	}
	return &ID, nil
}
