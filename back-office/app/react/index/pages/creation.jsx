import React from "react"

import Questions from "../components/creation/question.jsx"

//TODO:To create q from server

export default class Creation extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
        questions: [{title: ''}],
        update: false
    }

    this.questionChange = this.questionChange.bind(this)
    this.delBtnEvt = this.delBtnEvt.bind(this)
    this.addBtnEvt = this.addBtnEvt.bind(this)
    this.createBtnEvt = this.createBtnEvt.bind(this)
  }

  //change value of question form
  questionChange(value, index) {
    console.log("[Creation]:questionChange()")
    console.log(value, index)

    let q = this.state.questions
    q[index].title = value

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
    console.log("[Question]:addBtnEvt(e)")

    let q = this.state.questions
    q.push({title: ''})

    this.setState({
      questions: q,
      update: true
    })
  }

  //Click create button
  //TODO:Ajax
  createBtnEvt(e) {
    console.log("[Question]:createBtnEvt(e)")

    //call ajax
    //let url = '/json/questionnaireList.json'
    //this.callAjax(url, 'json', '')
  }

  //when stopping render, return false
  shouldComponentUpdate(nextProps, nextState){
    return nextState.update
  }

  render() {
    console.log("[Creation]render()")
    let that = this
    let total = this.state.questions.length
    let questions = this.state.questions.map(function (data, index) {
      index++
      let key = `question_${index}_${data.title}`
      return (
        <Questions key={key} chgQ={that.questionChange} btnDel={that.delBtnEvt} total={total} idx={index} question={data.title} />
      )
    })

    return (
      <div>
        <h1>Creation</h1>
        <div className="input-group input-group-lg">
          <span className="input-group-addon" id="sizing-addon1">Title</span>
          <input type="text" className="form-control" placeholder="title" aria-describedby="sizing-addon1" />
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
