import React     from "react"
import $         from 'jquery'

import Questions from "../components/creation/question.jsx"

export default class Creation extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
        title: '',
        questions: [''],
        error:{
            title:'',
            question:[]
        },
        update: false
    }

    this.changeTitleEvt = this.changeTitleEvt.bind(this)
    this.questionChange = this.questionChange.bind(this)
    this.delBtnEvt = this.delBtnEvt.bind(this)
    this.addBtnEvt = this.addBtnEvt.bind(this)
    this.createBtnEvt = this.createBtnEvt.bind(this)
  }

  //Ajax
  callAjax(passedURL, sendData) {
    console.log("[Creation]:callAjax")

    let that = this
    let method = 'post'
    let contentType = "application/json"
    if(sendData != ''){
      sendData = JSON.stringify(sendData);
    }

    $.ajax({
      url: encodeURI(passedURL),
      type: method,
      //cache    : false,
      crossDomain: false,
      contentType: contentType,
      dataType:    'json', //data type from server
      data:        sendData
    }).done(function (data, textStatus, jqXHR) {
      console.log(data)
      //TODO:reset form
      that.setState({
        title: '',
        questions: [''],
        update: true
       })
      swal("success!", "done!", "success")
    }).fail(function (jqXHR, textStatus, errorThrown) {
      console.error(passedURL, textStatus, errorThrown.toString())
      swal("error!", "validation error was occurred!", "error")
    })
  }

  //change title
  changeTitleEvt(e) {
    console.log("[Creation]:changeTitleEvt()")
    //console.log(e.target.value)

    this.setState({
      title: e.target.value,
      update: true
    })
  }

  //change value of question form
  questionChange(value, index) {
    console.log("[Creation]:questionChange()")
    //console.log(value, index)

    let q = this.state.questions
    q[index] = value

    this.setState({
      questions: q,
      update: false
    })
  }

  //Click delete btn
  delBtnEvt(index) {
    console.log("[Creation]:delBtnEvt()")
    console.log(" index is ", index)

    let q = this.state.questions
    q.splice(index, 1)
    console.log(q)

    this.setState({
      questions: q,
      update: true
    })
  }

  //Click add button
  addBtnEvt(e) {
    console.log("[Creation]:addBtnEvt(e)")

    let q = this.state.questions
    q.push('')

    this.setState({
      questions: q,
      update: true
    })
  }

  // validation
  checkValidation(){
    //validation check
    let rtnVal = true
    let tError = ''
    if(this.state.title == ''){
      rtnVal = false
      tError = 'blank is not allowed'
    }

    let qError = this.state.questions.map(function (data, index) {
      //console.log(data)
      if(data == ''){
        rtnVal = false
        return 'blank is not allowed'
      }else{
        return ''
      }
    })
    this.setState({
      error : {
        title: tError,
        question: qError
      },
      update: true
    })
    return rtnVal
  }

  //Click create button
  createBtnEvt(e) {
    console.log("[Creation]:createBtnEvt(e)")

    //validation
    let bRet = this.checkValidation()
    //console.log(bRet)
    if(!bRet){
      return
    }

    //call ajax
    let url = '/api/ques'

    //send data
    //{"title":"title4", "questions":["q1","q2","q3"]}
    let sendData = new Object()
    sendData.title = this.state.title
    sendData.questions = this.state.questions

    //console.log(sendData)
    this.callAjax(url, sendData)
  }

  //when stopping render, return false
  shouldComponentUpdate(nextProps, nextState){
    return nextState.update
  }

  //
  render() {
    console.log("[Creation]render()")
    //console.log(this.state.error)

    let that = this
    let total = this.state.questions.length
    let questions = this.state.questions.map(function (data, index) {
      let key = `question_${index+1}_${data}`
      let err = that.state.error.question[index]
      //console.log("error is ", err)
      return (
        <Questions key={key} chgQ={that.questionChange} btnDel={that.delBtnEvt} total={total} idx={index}
          question={data} error={err} />
      )
    })

    return (
      <div>
        <h1>Creation</h1>
        <p className="text-danger">{this.state.error.title}</p>
        <div className="input-group input-group-lg">
          <span className="input-group-addon" id="sizing-addon1">Title</span>
          <input type="text" className="form-control" placeholder="title"
            onChange={this.changeTitleEvt} aria-describedby="sizing-addon1"
            value={this.state.title} />
        </div>
        <br />
        {questions}
        <button className="btn btn-primary" onClick={this.addBtnEvt} type="button">Add new question</button>
        <br />
        <br />
        <button className="btn btn-primary" onClick={this.createBtnEvt} type="button">Create</button>
        <br />
        <br />
      </div>
    )
  }
}
