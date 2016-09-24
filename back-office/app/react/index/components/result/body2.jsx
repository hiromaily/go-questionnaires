import React from "react"


export default class ResultBody2 extends React.Component {
  constructor(props) {
    super(props)
  }

  //
  render() {
    return (
      <th>{this.props.answer}</th>
    )
  }
}
