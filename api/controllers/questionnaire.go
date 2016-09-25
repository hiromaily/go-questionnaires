package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	m "github.com/hiromaily/go-questionnaires/api/models"
	lg "github.com/hiromaily/golibs/log"
	"net/http"
)

// ResponseQuestionnaires is for response of questionnaires list
type ResponseQuestionnaires struct {
	List []FormattedQuestionnaires `json:"list"`
}

// FormattedQuestionnaires is to store formated Questionnaires
type FormattedQuestionnaires struct {
	ID        int      `column:"questionnaire_id" json:"id"`
	Title     string   `column:"title" json:"title"`
	Questions []string `column:"questions" json:"questions"`
}

// FormattedAnswers is to store formated Answers
type FormattedAnswers struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Questions []string `json:"questions"`
	Answers   []Answer `json:"answers"`
}

// Answer is child of FormattedAnswers
type Answer struct {
	Email  string   `json:"email"`
	Answer []string `json:"answer"`
}

/*
{
  "questionnaire_id": 1,
  "questions": ["Q1 aaaaaa", "Q2 bbbbbb", "Q3 ccccc", "Q4 ddddd", "Q5 eeeee"],
  "answers": [{
    "email": "aaa@gmail.com",
    "answer": ["answer1", "answer2", "answer3", "answer4", "answer5"]
  }, {
    "email": "bbb@gmail.com",
    "answer": ["answer-b-1", "", "answer-b-3", "answer-b-4", "answer-b-5"]
  }, {
    "email": "ccc@gmail.com",
    "answer": ["answer-c-1", "answer-c-2", "answer-c-3", "answer-c-4", "answer-c-5"]
  }]
}
*/

// RequestQuestionnaires is request data when posting new data for questionnaires
type RequestQuestionnaires struct {
	Title     string   `json:"title"`
	Questions []string `json:"questions"`
}

// validation for post request for creation of questionnaire
func postRequestParamAndValid(c *gin.Context, req *RequestQuestionnaires) (err error) {
	lg.Info("postRequestParamAndValid()")

	contentType := c.Request.Header.Get("Content-Type")
	lg.Debug(" Content-Type is ", contentType)

	if contentType == "application/json" {
		err = c.BindJSON(req)
		if err != nil {
			return err
		}
	} else {
		return errors.New("contentType is invalid")
	}

	lg.Debugf(" Request Body: %#v", req)

	//Validation
	if req.Title == "" {
		return errors.New("blank is not allowed at title")
	}
	if len(req.Questions) == 0 {
		return errors.New("at reast one question is required")
	}
	for _, v := range req.Questions {
		if v == "" {
			return errors.New("blank is not allowed at question")
		}
	}

	return nil
}

// GetQuesAction is to get questionnaire list
func GetQuesAction(c *gin.Context) {
	lg.Info("GetQuesAction()")

	// get questionnaires
	var q []m.Questionnaires
	err := m.GetDB().GetQuestionnaireList(&q)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	//convert data
	//var f []FormatedQuestionnaires
	f := make([]FormattedQuestionnaires, 0, len(q))
	for _, v := range q {
		//json
		var questions []string
		json.Unmarshal([]byte(v.Questions), &questions)

		f = append(f, FormattedQuestionnaires{
			ID:        v.ID,
			Title:     v.Title,
			Questions: questions,
		})
	}

	response := ResponseQuestionnaires{
		List: f,
	}

	//{"list":[{"id":1,"title":"title1","created":"2016-09-24 21:43:15","updated":"2016-09-24 21:43:15"},{"id":2,"title":"title2","created":"2016-09-24 21:43:15","updated":"2016-09-24 21:43:15"},{"id":3,"title":"title3","created":"2016-09-24 21:43:15","updated":"2016-09-24 21:43:15"}]}
	c.JSON(http.StatusOK, response)
}

// PostRegQuesAction to register new questionnaire
func PostRegQuesAction(c *gin.Context) {
	lg.Info("PostRegQuesAction()")

	//Param & Check valid
	var request RequestQuestionnaires
	err := postRequestParamAndValid(c, &request)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	//convert questions to json
	retByte, _ := json.Marshal(request.Questions)
	lg.Debugf("%s\n", string(retByte))

	// Insert questionnaire
	newID, err := m.GetDB().InsertQuestionnaire(request.Title, string(retByte))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	//respolnse
	c.JSON(http.StatusOK, gin.H{
		"newID": newID,
	})
}

// DeleteQuesAction is to Delete questionnaire by ID
func DeleteQuesAction(c *gin.Context) {
	lg.Info("DeleteQuesAction()")

	//Param
	if c.Param("id") == "" {
		c.AbortWithError(400, errors.New("missing id on request parameter"))
		return
	}

	//Delete
	affected, err := m.GetDB().DeleteQuestionnaire(c.Param("id"))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	if affected == 0 {
		lg.Debug("there was no updated data.")
	}

	c.JSON(http.StatusOK, gin.H{
		"affected": affected,
	})
}

// GetAnswerAction is to get answer by ID
func GetAnswerAction(c *gin.Context) {
	lg.Info("GetAnswerAction()")

	//Param
	if c.Param("id") == "" {
		c.AbortWithError(400, errors.New("missing id on request parameter"))
		return
	}

	var ans []m.Answers
	err := m.GetDB().GetAnswersByID(&ans, c.Param("id"))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	//lg.Debug(ans)
	//[{1 title1 abc@gmail.com ["answer1", "answer2", "answer3"]} {1 title1 xxxx@gmail.com ["aaaaa111", "bbbbb222", "ccccc333"]}]

	//convert data
	var f FormattedAnswers
	//var a []Answer
	a := make([]Answer, 0, len(ans))

	f.ID = ans[0].ID
	f.Title = ans[0].Title
	json.Unmarshal([]byte(ans[0].Questions), &f.Questions)

	for _, v := range ans {
		//json
		var answers []string
		json.Unmarshal([]byte(v.Answers), &answers)

		a = append(a, Answer{
			Email:  v.Email,
			Answer: answers,
		})
	}
	f.Answers = a

	//lg.Debug(f)

	c.JSON(http.StatusOK, f)
}
