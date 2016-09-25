package models

import (
	"fmt"
	"github.com/hiromaily/golibs/db/mysql"
	lg "github.com/hiromaily/golibs/log"
	u "github.com/hiromaily/golibs/utils"
)

// table name
const tblQuestionnaires = "t_questionnaires"

// Questionnaires is of t_questionnaires table structure.
type Questionnaires struct {
	ID        int    `column:"questionnaire_id" json:"id"`
	Title     string `column:"title" json:"title"`
	Questions string `column:"questions" json:"questions"`
	//Created   string `column:"created" json:"created"`
	//Updated   string `column:"updated" json:"updated"`
}

// QuestionnairesByInsert is of t_questionnaires table structure.
type QuestionnairesByInsert struct {
	Title string `column:"title" json:"title"`
}

// GetQuestionnaireList is to get questionnaire list
func (m *Models) GetQuestionnaireList(questionnaires interface{}) error {
	lg.Info("GetQuestionnaireList()")

	sql := "SELECT %s FROM %s WHERE delete_flg=?"
	sql = fmt.Sprintf(sql, mysql.ColumnForSQL(questionnaires), tblQuestionnaires)
	lg.Debugf("sql is %s", sql)

	//When Test for result is 0 record, set 1 to delFlg
	delFlg := 0
	_ = m.Db.SelectIns(sql, delFlg).Scan(questionnaires)

	return m.Db.Err
}

// InsertQuestionnaire is to post new questionnaire
func (m *Models) InsertQuestionnaire(title string, questions string) (int64, error) {
	lg.Info("GetQuestionnaireList()")

	//TODO:transaction if needed

	//Insert t_questionnaires
	sql := fmt.Sprintf("INSERT INTO %s (title, questions) VALUES (?, ?)", tblQuestionnaires)
	newID, err := m.Db.Insert(sql, title, questions)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// DeleteQuestionnaire is to delete questionnaire by ID
func (m *Models) DeleteQuestionnaire(id string) (int64, error) {
	lg.Info("DeleteQuestionnaire()")

	sql := fmt.Sprintf("DELETE FROM %s WHERE questionnaire_id=?", tblQuestionnaires)
	return m.Db.Exec(sql, u.Atoi(id))
}
