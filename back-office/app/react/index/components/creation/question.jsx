import React from "react"


export default class Question extends React.Component {
  constructor(props) {
    super(props)

    this.changeQuestionEvt = this.changeQuestionEvt.bind(this)
    this.clickBtnEvt = this.clickBtnEvt.bind(this)
  }

  //Change question event
  changeQuestionEvt(e) {
    console.log("[Question]:changeQuestionEvt(e)")
    //call
    this.props.chgQ.call(this, e.target.value, this.props.idx-1)
  }

  //Click del button
  clickBtnEvt(e) {
    console.log("[Question]:clickBtnEvt(e)")

    if(this.props.total == 1){
        //error
        swal("error!", "At least one question is required.", "error")
        return
    }

    //call event for post btn click
    this.props.btnDel.call(this, this.props.idx-1)
  }

  render() {

    return (
      <div>
      <div className="input-group">
        <span className="input-group-addon" id="sizing-addon2">Question{this.props.idx}</span>
        <input type="text" className="form-control" placeholder={"question" + this.props.idx}
          onChange={this.changeQuestionEvt} aria-describedby="sizing-addon2" defaultValue={this.props.question} />
        <span className="input-group-btn">
          <button className="btn btn-secondary" onClick={this.clickBtnEvt} type="button">Delete</button>
        </span>
      </div>
      <br />
      </div>
    )
  }
}
