import React     from "react"
import ReactDOM  from "react-dom"
import { Router, Route, IndexRoute, hashHistory } from "react-router"

import Layout    from "./pages/layout.jsx"

import List      from "./pages/list.jsx"
import Creation  from "./pages/creation.jsx"
import Result    from "./pages/result.jsx"


export default class App extends React.Component {
  render() {
    return (
      <Router history={hashHistory}>
        <Route path="/" component={Layout}>
          <IndexRoute component={List}></IndexRoute>
          <Route path="creation" name="creation" component={Creation}></Route>
          <Route path="result(/:id)" name="result" component={Result}></Route>
        </Route>
      </Router>
    )
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('root')
)
