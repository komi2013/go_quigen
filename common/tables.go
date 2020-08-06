package common

import (
    "time"
)

type HAnswer struct {
	QuestionID    int       // question_id
	QuestionTxt   string    // question_txt
	GeneratorID   int       // generator_id
	GeneratorImg  string    // generator_img
	AskedAt       time.Time // asked_at
	Choice0       string    // choice_0
	Choice1       string    // choice_1
	Choice2       string    // choice_2
	Choice3       string    // choice_3
	Reference     string    // reference
	QuestionType  int16     // question_type
	CategoryID    int       // category_id
	QuestionImg   string    // question_img
	RespondentID  int       // respondent_id
	RespondentImg string    // respondent_img
	Sequence      int       // sequence
	Mytext        string    // mytext
	Mychoice      int16     // mychoice
	Count         int       // count
	Explanation   string    // explanation
	CreatedAt     time.Time // created_at
	EtoColor      string    // not in table
}

type MCategoryName struct {
	CategoryID          int       // category_id
	CategoryName        string    // category_name
	UpdatedAt           time.Time // updated_at
	CategoryDescription string    // category_description
}

type MCategoryQuestion struct {
	QuestionID      int       // question_id
	CategoryID      int       // category_id
	UpdatedAt       time.Time // updated_at
	QuestionTxt     string    // question_txt
}

type MCategoryTree struct {
	LeafID    int       // leaf_id
	Level1    int       // level_1
	Level2    int       // level_2
	Level3    int       // level_3
	Level4    int       // level_4
	Level5    int       // level_5
	Level6    int       // level_6
	Level7    int       // level_7
	Level8    int       // level_8
	UpdatedAt time.Time // updated_at
}

type MPublicMessage struct {
	PublicMessageID  int       // public_message_id
	PublicMessageTxt string    // public_message_txt
	UpdatedAt        time.Time // updated_at
}

type TComment struct {
	CommentID  int       // comment_id
	CommentTxt string    // comment_txt
	UsrID      int       // usr_id
	QuestionID int       // question_id
	CreatedAt  time.Time // created_at
	UsrImg     string    // usr_img
	EtoColor   string     // not in table
}

type TMessage struct {
	MessageID  int       // message_id
	Sender     int       // sender
	Receiver   int       // receiver
	MessageTxt string    // message_txt
	CreatedAt  time.Time // created_at
	MessageImg string    // message_img
}

type TQuestion struct {
	QuestionID       int       // question_id
	QuestionTxt      string    // question_txt
	UsrID            int       // usr_id
	UsrImg           string    // usr_img
	CreatedAt        time.Time // created_at
	UpdatedAt        time.Time // updated_at
	Choice0          string    // choice_0
	Choice1          string    // choice_1
	Choice2          string    // choice_2
	Choice3          string    // choice_3
	Reference        string    // reference
	QuestionType     int16     // question_type
	CategoryID       int       // category_id
	QuestionImg      string    // question_img
	PreviousQuestion int       // previous_question
	NextQuestion     int       // next_question
	Sequence         int       // sequence
	Sound            string    // sound
	Explanation      string    // explanation
}

type TUsr struct {
	UsrID        int       // usr_id
	PvUID        string    // pv_u_id
	Provider     int       // provider
	UsrName      string    // usr_name
	UsrImg       string    // usr_img
	UpdatedAt    time.Time // updated_at
	Point        int       // point
	Introduce    string    // introduce
	Nice         int       // nice
	Certify      int       // certify
	Quiz         int       // quiz
	Forum        int       // forum
	ForumComment int       // forum_comment
	Latitude     float64   // latitude
	Longitude    float64   // longitude
	PushTokens   []byte    // push_tokens
}

