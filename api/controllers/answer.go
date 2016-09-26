package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	m "github.com/hiromaily/go-questionnaires/api/models"
	lg "github.com/hiromaily/golibs/log"
	"net/http"
)

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

// RequestAnswers is request data when posting new data for answers
type RequestAnswers struct {
	Email   string   `json:"email"`
	Answers []string `json:"answers"`
}

// validation for post request for creation of answer
func validationAnswer(c *gin.Context, req *RequestAnswers) (err error) {
	lg.Info("validationAnswer()")

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
	//Param
	if c.Param("id") == "" {
		c.AbortWithError(400, errors.New("missing id on request parameter"))
		return
	}

	if req.Email == "" {
		return errors.New("blank is not allowed at email")
	}
	if len(req.Answers) == 0 {
		return errors.New("at reast one question is required")
	}

	//TODO: check if same as number of question

	for _, v := range req.Answers {
		if v == "" {
			return errors.New("blank is not allowed at question")
		}
	}

	return nil
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
		//when no data,
		lg.Error(err)
		c.AbortWithError(400, err)
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

// PostAnswerAction is to register answers by ID and Email
func PostAnswerAction(c *gin.Context) {
	lg.Info("PostAnswerAction()")

	//Param & Check valid
	var request RequestAnswers
	err := validationAnswer(c, &request)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	//convert questions to json
	retByte, _ := json.Marshal(request.Answers)
	lg.Debugf("%s\n", string(retByte))

	//TODO:check email has already been registered.
	//TODO:update may be better
	//Search existing email or not
	err = m.GetDB().IsEmailExisting(c.Param("id"), request.Email)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	// Insert questionnaire
	newID, err := m.GetDB().InsertAnswer(c.Param("id"), request.Email, string(retByte))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	//respolnse
	c.JSON(http.StatusOK, gin.H{
		"newID": newID,
	})

}
