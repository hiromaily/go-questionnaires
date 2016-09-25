import React       from "react"
import $           from 'jquery'

import ListDetails from "../components/list/details.jsx"

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
  callAjax(passedURL, method, id) {
    console.log("[List]:callAjax")

    let that = this
    //let method = 'get'
    let contentType = "application/json"

    $.ajax({
      url: encodeURI(passedURL),
      type: method,
      //cache    : false,
      crossDomain: false,
      contentType: contentType,
      dataType:    'json', //data type from server
      data:        ''
    }).done(function (data, textStatus, jqXHR) {
      console.log(data)
      if(method == 'get'){
        that.setState({
          list: data.list
        })
      }else if(method == 'delete'){
        //TODO:remove element
        let newList = that.state.list.map(function (data, index) {
          console.log(data)
          if(data.id == id){
            return
          }else{
            return data
          }
        })
        console.log(newList)
        //that.setState({
        //  list: newList
        //})
        //undefined

      }
    }).fail(function (jqXHR, textStatus, errorThrown) {
      console.error(passedURL, textStatus, errorThrown.toString())
      swal("error!", "validation error was occurred!", "error")
    })
  }

  //Get Questionnaire List
  getQuestionnaireList() {
    console.log("[List]:getQuestionnaireList()")

    //call ajax
    //let url = '/admin/json/questionnaireList.json'
    let url = '/api/ques'
    this.callAjax(url, 'get', 0)
  }

  //Click delete btn
  delBtnEvt(id) {
    console.log("[List]:delBtnEvt()")
    console.log(" id is ", id)

    let url = '/api/ques/'+id
    this.callAjax(url, 'delete', id)
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
              <th>ID</th>
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
