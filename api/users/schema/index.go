package UsersSchema

type RequestSignUp struct {
	UserCode   string `json:"user_code" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Firstname  string `json:"firstname" validate:"required"`
	Lastname   string `json:"lastname" validate:"required"`
	BranchCode string `json:"branch_code" validate:"required"`
}

type RequestSignIn struct {
	UserCode string `json:"user_code" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Users struct {
	QuestionUserID string `json:"question_user_id"`
	UserCode       string `json:"user_code"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	BranchCode     string `json:"branch_code"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type UserSignIn struct {
	QuestionUserID string `json:"question_user_id"`
	UserCode       string `json:"user_code"`
	BranchCode     string `json:"branch_code"`
	Password       string `json:"password"`
}
