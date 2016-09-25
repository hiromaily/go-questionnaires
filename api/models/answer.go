package models

import (
	"fmt"
	lg "github.com/hiromaily/golibs/log"
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
