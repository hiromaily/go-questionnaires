package models

import (
	"errors"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	u "github.com/hiromaily/golibs/utils"
)

// table name
const tblAnswers = "t_answers"

// Answers is joined data of t_answers and t_questionnaires table structure.
type Answers struct {
	ID        int    `column:"questionnaire_id" json:"id"`
	Title     string `column:"title" json:"title"`
	Questions string `column:"questions" json:"questions"`
	Email     string `column:"user_email" json:"email"`
	Answers   string `column:"answers" json:"answers"`
	//Created   string `column:"created" json:"created"`
	//Updated   string `column:"updated" json:"updated"`
}

// GetAnswersByID is to get answer list
func (m *Models) GetAnswersByID(questionnaires interface{}, id string) error {
	lg.Info("GetAnswersByID()")

	sql := `
SELECT q.questionnaire_id, q.title, q.questions, a.user_email, a.answers
 FROM %s AS q
 LEFT JOIN %s AS a ON q.questionnaire_id=a.questionnaire_id
 WHERE q.delete_flg=?
 AND a.delete_flg=?
 AND q.questionnaire_id=?
`
	sql = fmt.Sprintf(sql, tblQuestionnaires, tblAnswers)
	lg.Debugf("sql is %s", sql)

	//When Test for result is 0 record, set 1 to delFlg
	delFlg := 0
	_ = m.Db.SelectIns(sql, delFlg, delFlg, id).Scan(questionnaires)

	return m.Db.Err
}

// IsEmailExisting is to check existing email
func (m *Models) IsEmailExisting(id, email string) error {
	lg.Info("IsEmailExisting()")

	sql := `
SELECT count(answers_id)
 FROM %s
 WHERE questionnaire_id=?
 AND user_email=?
 AND delete_flg="0"
`
	sql = fmt.Sprintf(sql, tblAnswers)

	// get count
	count, err := m.Db.SelectCount(sql, id, email)

	if err != nil {
		lg.Debugf("IsEmailExisting() Error: %s", err)
		return err
	}

	if count != 0 {
		return errors.New("this email has already been registered")
	}

	return nil
}

// InsertAnswer is to post new answer
func (m *Models) InsertAnswer(id string, email string, answers string) (int64, error) {
	lg.Info("GetQuestionnaireList()")

	//Insert t_answers
	sql := fmt.Sprintf("INSERT INTO %s (questionnaire_id, user_email, answers) VALUES (?, ?, ?)", tblAnswers)
	newID, err := m.Db.Insert(sql, u.Atoi(id), email, answers)
	if err != nil {
		return 0, err
	}

	return newID, nil
}
