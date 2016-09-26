import React        from "react"
import $            from 'jquery'

import ResultHeader from "../components/result/header.jsx"
import ResultBody   from "../components/result/body.jsx"


export default class Result extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
        result: {
            'id': 0,
            'title': '',
            'questions': [],
            'answers': []
        }
    }

    this.getResult = this.getResult.bind(this)
  }

  //Ajax
  callAjax(passedURL) {
    console.log("[Result]:callAjax")

    let that = this
    let method = 'get'
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
      that.setState({
        result: data
      })
    }).fail(function (jqXHR, textStatus, errorThrown) {
      console.error(passedURL, textStatus, errorThrown.toString())
      swal("error!", "validation error was occurred or no result!!", "error")
    })
  }

  //Get Result
  getResult() {
    console.log("[Result]:getResult()")

    //call ajax
    //let url = '/admin/json/result1.json'
    let url = '/api/answer/'+this.props.params.id

    this.callAjax(url)
  }

  //Only once before first render()
  componentWillMount() {
    console.log("[Result]:componentWillMount()")
    //console.log(this.props.params.id)
    //console.log(this.props.location.query)

    this.getResult()
  }

  //
  render() {
    console.log("[Result]render()")
    let header = this.state.result.questions.map(function (title, index) {
      return (
        <ResultHeader key={index} title={title} />
      )
    })

    let body = this.state.result.answers.map(function (data, index) {
      return (
        <ResultBody key={index} idx={index+1} email={data.email} answer={data.answer} />
      )
    })

    return (
      <div>
        <h1>Result</h1>
        <h3>{this.props.location.query.title}</h3>
        <table className="table">
          <thead>
            <tr>
              <th>No.</th>
              <th>Email</th>
              {header}
            </tr>
          </thead>
          <tbody>
            {body}
          </tbody>
        </table>
      </div>
    )
  }
}
