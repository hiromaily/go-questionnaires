import React from "react"
import { Link } from "react-router"


export default class ListDetail extends React.Component {
  constructor(props) {
    super(props)

    this.clickBtnEvt = this.clickBtnEvt.bind(this)
  }

  //Click del button
  clickBtnEvt(e) {
    console.log("[ListDetail]:clickBtnEvt(e)")

    //call event for post btn click
    this.props.btnDel.call(this, this.props.id)
  }

  render() {

    return (
      <tr>
        <th scope="row">{this.props.id}</th>
        <td>{this.props.title}</td>
        <td>
          <Link to={"/result/" + this.props.id} query={{title: this.props.title}}>result</Link>
        </td>
        <td><button type="button" onClick={this.clickBtnEvt} className="btn btn-danger btn-sm">Delete</button></td>
      </tr>
    )
  }
}
