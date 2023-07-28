package QuestionService

import (
	QuestionRepository "go-template/api/questions/repository"
	Schema "go-template/api/questions/schema"
)

func CreateQuestion(question Schema.Question, descriptions []Schema.Description) error {
	QuestionID, err := QuestionRepository.CreateQuestion(question.Title, question.BranchType)
	if err != nil {
		return err
	}
	var ID []string
	for _, description := range descriptions {
		descriptionID, err := QuestionRepository.CreateDescription(
			description.DescriptionText,
			description.DescriptionType,
			description.DescriptionTrueString,
			description.DescriptionFalseString,
			description.Priority,
			*QuestionID)
		if err != nil {
			ID = append(ID, *descriptionID)
		}
	}
	return nil
}
