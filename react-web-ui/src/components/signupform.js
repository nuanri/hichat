import React, {Component, PropTypes} from 'react';
import {reduxForm} from 'redux-form';
import { pushPath } from 'redux-simple-router'
import {api_post} from '../api'

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

    return (<form onSubmit={handleSubmit(this._submit)}>
        <div>
          <label>用户名</label>
          <div>
            <input type="text" placeholder="用户名" {...userName}/>
          </div>
        </div>

        <div>
          <label> 密码 </label>
          <div>
            <input type="text" placeholder="密码" {...passWord}/>
          </div>
        </div>

        <div>
          <button disabled={submitting} onClick={handleSubmit(this._submit)}>
            Submit
          </button>
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
      .then(res => res.json())
      .then(json => {
        console.log('json = ', json)
        if (json.error) {
          reject(Object.assign({}, json.errors, {
            _error: json.error,
          }))
        } else {
          if (json.sid) {
            localStorage.setItem('sid', json.sid);
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
  })(AuthSignUpForm);
