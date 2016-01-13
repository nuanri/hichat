import React, {Component, PropTypes} from 'react'
import {reduxForm} from 'redux-form'
import { pushPath } from 'redux-simple-router'
import {api_post} from '../api'
import { readAccountInfo } from '../actions/account_read'

export const fields = ['userName', 'passWord'];

class AuthSignUpForm extends Component {

  constructor(props) {
    super(props)
    this._submit = this._submit.bind(this)
  }

  render() {
    const {
      fields: {userName, passWord},
      handleSubmit,
      submitting
    } = this.props;

    return (
      <form className="siginform form-horizontal" onSubmit={handleSubmit(this._submit)}>
        <div className="form-group">
          <label className="col-sm-2 control-label">
            <span className="glyphicon glyphicon-user" aria-hidden="true"></span>
          </label>
          <div className="col-sm-10">
            <input type="text" placeholder="用户名" {...userName}/>
          </div>
        </div>

        <div className="form-group">
          <label className="col-sm-2 control-label">
            <span className="glyphicon glyphicon-lock" aria-hidden="true"></span>
          </label>
          <div className="col-sm-10">
            <input type="password" placeholder="密码" {...passWord}/>
          </div>
        </div>

        <div className="form-group">
          <div className="col-sm-offset-2 col-sm-10">
            <button  className="btn btn-default" type="submit" disabled={submitting} onClick={handleSubmit(this._submit)}>
              登录
            </button>
          </div>
        </div>
      </form>
    );
  }

  _submit(values, dispatch) {
    console.log('values = ', values)
    console.log('dispatch = ', dispatch)

    return new Promise((resolve, reject) => {
      return api_post(`/auth/signin`, {
        body: JSON.stringify(values)
      })
      // .then(res => res.json())
      .then(json => {
        console.log('json = ', json)
        if (json.error) {
          reject(Object.assign({}, json.errors, {
            _error: json.error,
          }))
        } else {
          if (json.sid) {
            localStorage.setItem('sid', json.sid);
            dispatch(readAccountInfo())
            dispatch(pushPath("/"))
            resolve();
          } else {
            reject({_error: '没有发现sid'})
          }
        }
      })
      .catch(error => {
        reject({_error: error})
      })
    });
  }
}

AuthSignUpForm.propTypes = {
  fields: PropTypes.object.isRequired,
  handleSubmit: PropTypes.func.isRequired,
  submitting: PropTypes.bool.isRequired
}


export default reduxForm({
    form: 'signup',
    fields: ['userName', 'passWord']
  })(AuthSignUpForm)
