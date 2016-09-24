import React from "react"

import ResultBody2   from "../components/result/body2.jsx"

export default class ResultBody extends React.Component {
  constructor(props) {
    super(props)
  }

  //
  render() {

    let answers = this.props.answers.map(function (answer) {
      return (
        <ResultBody2 key={index} answer={answer.answer} />
      )
    })

    return (
      <tr>
        <th scope="row">{this.props.idx}</th>
        <td>{this.props.email}</td>

        <td>
        </td>

      </tr>

      <th>{this.props.title}</th>
    )
  }
}
