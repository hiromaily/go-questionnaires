import React from "react"
import { IndexLink, Link } from "react-router"

export default class Nav extends React.Component {
  constructor() {
    super()
    this.state = {
      collapsed: true,
    }
  }

  toggleCollapse() {
    const collapsed = !this.state.collapsed
    this.setState({collapsed})
  }

  render() {
    const { location } = this.props
    const { collapsed } = this.state
    const listClass = location.pathname === "/" ? "active" : ""
    const creationClass = location.pathname.match(/^\/creation/) ? "active" : ""
    const resultClass = location.pathname.match(/^\/result/) ? "active" : ""
    const navClass = collapsed ? "collapse" : ""

    return (
      <nav className="navbar navbar-inverse navbar-fixed-top" role="navigation">
        <div className="container">
          <div className="navbar-header">
            <button type="button" className="navbar-toggle" onClick={this.toggleCollapse.bind(this)} >
              <span className="sr-only">Toggle navigation</span>
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
            </button>
          </div>
          <div className={"navbar-collapse " + navClass} id="bs-example-navbar-collapse-1">
            <ul className="nav navbar-nav">
              <li className={listClass}>
                <IndexLink to="/" onClick={this.toggleCollapse.bind(this)}>List</IndexLink>
              </li>
              <li className={creationClass}>
                <Link to="creation" onClick={this.toggleCollapse.bind(this)}>Creation</Link>
              </li>
              <li className={resultClass}>
                <Link to="result" onClick={this.toggleCollapse.bind(this)}>Result</Link>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    )
  }
}
