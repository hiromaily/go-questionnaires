import React from "react"

import ResultBody2   from "./body2.jsx"

export default class ResultBody extends React.Component {
  constructor(props) {
    super(props)
  }

  //
  render() {

    let answers = this.props.answer.map(function (answer, index) {
      return (
        <ResultBody2 key={index} answer={answer} />
      )
    })

    return (
      <tr>
        <th scope="row">{this.props.idx}</th>
        <td>{this.props.email}</td>
        {answers}
      </tr>

    )
  }
}
