package common

import (
    "time"
)

type HAnswer struct {
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
	Mytext         string     // mytext
	Spend            int16     // spend
	Mychoice            int16     // mychoice
	EtoColor         string     // not in table
}

type MCategoryName struct {
	CategoryID          int       // category_id
	CategoryName        string    // category_name
	UpdatedAt           time.Time // updated_at
	CategoryDescription string    // category_description
}

type MCategoryQuestion struct {
	QuestionID int       // question_id
	CategoryID int       // category_id
	UpdatedAt  time.Time // updated_at
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

type CommentCreatedAt []TComment
func (p CommentCreatedAt) Len() int {
    return len(p)
}
func (p CommentCreatedAt) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p CommentCreatedAt) Less(i, j int) bool {
    return p[i].CreatedAt.After(p[j].CreatedAt)
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

