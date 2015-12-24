var React = require('react');
var ReactDOM = require('react-dom')

import { Link } from 'react-router'

export default class App extends React.Component {
  render(){
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
    );
  }
}
