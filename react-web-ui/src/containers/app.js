var React = require('react');
var ReactDOM = require('react-dom')
import { connect } from 'react-redux'

import { Link } from 'react-router'

import { readAccountInfo } from '../actions/account_read'

class App extends React.Component {

  componentDidMount() {
    let { dispatch } = this.props
    dispatch(readAccountInfo())
  }

  render(){
    let { account } = this.props
    console.log('App: account = ', this.props.account)

    if (account && !account.error) {
      let current = account.current
      return (
        <div>
          <nav className="navbar navbar-default">
            <div className="container-fluid">
              <div className="navbar-header">
                <Link className="navbar-brand" to={`/`}>Brand</Link>
              </div>
              <ul className="nav navbar-nav navbar-right">
                <li><Link to="#">{current.username}</Link></li>
                <li><Link to={`/auth/signout`}>注销</Link></li>
              </ul>
            </div>
            </nav>
          <div>
            { this.props.children }
          </div>
        </div>
      )
    } else {
      return (
        <div>
          <nav className="navbar navbar-default">
            <div className="container-fluid">
              <div className="navbar-header">
                <Link className="navbar-brand" to={`/`}>Brand</Link>
              </div>
              <ul className="nav navbar-nav navbar-right">
                <li><Link to="#">注册</Link></li>
                <li><Link to={`/auth/signup`}>登录</Link></li>
              </ul>
            </div>
            </nav>
          <div>
            { this.props.children }
          </div>
        </div>
      )
    }
  }
}

export default connect(state => ({
  account: state.main.account,
}))(App)
