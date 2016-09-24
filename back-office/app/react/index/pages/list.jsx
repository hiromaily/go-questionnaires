import React       from "react"
import $           from 'jquery'

import ListDetails from "../components/list/details.jsx"

//TODO:To get q list from server: change url (now it's static json)
//TODO:To delete q from server

export default class List extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
        list: []
    }

    this.getQuestionnaireList = this.getQuestionnaireList.bind(this)
    this.delBtnEvt = this.delBtnEvt.bind(this)
  }

  //Ajax
  callAjax(passedURL, sendData) {
    console.log("[List]:callAjax")

    let that = this
    let method = 'get'
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
      that.setState({
        list: data.list
      })
    }).fail(function (jqXHR, textStatus, errorThrown) {
      console.error(url, textStatus, errorThrown.toString())
      swal("error!", "validation error was occurred!", "error")
    })
  }

  //Get Questionnaire List
  getQuestionnaireList() {
    console.log("[List]:getQuestionnaireList()")

    //call ajax
    let url = '/json/questionnaireList.json'
    this.callAjax(url, '')
  }

  //Click delete btn
  delBtnEvt(id) {
    console.log("[List]:delBtnEvt()")
    console.log(" id is ", id)
  }

  //Only once before first render()
  componentWillMount() {
    console.log("[List]:componentWillMount()")
    this.getQuestionnaireList()
  }

  //
  render() {
    console.log("[List]render()")

    let that = this
    let list = this.state.list.map(function (data) {
      let key='questionnaire_' + data.id
      return (
        <ListDetails key={key} id={data.id} title={data.title} btnDel={that.delBtnEvt} />
      )
    })

    return (
      <div>
        <h1>List</h1>
        <table className="table">
          <thead>
            <tr>
              <th>No.</th>
              <th>Title</th>
              <th>Result</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            {list}
          </tbody>
        </table>
      </div>
    )
  }
}
