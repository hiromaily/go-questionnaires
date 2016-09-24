import React from "react"


export default class ResultHeader extends React.Component {
  constructor(props) {
    super(props)
  }

  //
  render() {
    return (
      <th>{this.props.title}</th>
    )
  }
}
