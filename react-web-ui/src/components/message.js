import React, {Component, PropTypes} from 'react'
import { connect } from 'react-redux'
import {reduxForm} from 'redux-form'
import { pushPath } from 'redux-simple-router'
import {api_post} from '../api'
import { NiceDate } from './date';

import { sendingMessage, listMessage } from '../actions/message'

export const fields = ['body']

class MessageList extends Component {

  componentDidMount() {
    this.props.dispatch(listMessage())
  }

  render() {
    console.log('this.props = ', this.props)
    let { lists } = this.props.message

    return (
      <div>
        {lists.map((m, i) => (
          <div>
          <p>{m.username}<NiceDate date={m.add_time}/></p>
          <p key={i}>{m.msg}</p>
          </div>
        ))}
      </div>
    )
  }
}

MessageList = connect(state => ({
  message: state.main.message,
}))(MessageList)


export default class Message extends Component {

  constructor(props) {
    super(props)
      this._submit = this._submit.bind(this)
  }

  render(){
    const {
      fields,
      handleSubmit,
      submitting,
    } = this.props;

    let { onlineUsers } = this.props.message

    return (
      <div className="row">
        <div className="col-lg-3">
          <div className="panel panel-default status-friend">
            <div className="panel-heading">
              在线成员
            </div>
            <div className="panel-body onlineuser">
              <div>
                {onlineUsers.map((m, i) => (
                  <p>
                    <span className="glyphicon glyphicon-user" aria-hidden="true"></span>
                    <span  className="user-list" key={i}>{m}</span>
                  </p>
                ))}
              </div>
            </div>
          </div>
        </div>

        <div className="col-lg-9">
          <div className="panel panel-default">
            <div className="panel-body send-body">
              <MessageList />
            </div>
            <div className="textarea-wrapper">
              <form onSubmit={handleSubmit(this._submit)}>
                <textarea className="form-control message-input" rows="3" placeholder="Hello, World!" {...fields.body}>
                </textarea>
                <span className="input-group-btn">
                   <button   disabled={submitting} onClick={handleSubmit(this._submit)} className="btn btn-default" type="button">发送</button>
                </span>
              </form>
            </div>
          </div>
        </div>
      </div>
    );
  }

  _submit(values, dispatch) {
    console.log('values=',values)
    dispatch(sendingMessage(values))
    return new Promise((resolve, reject) => {
      return api_post('/messages',{
        body: JSON.stringify(values),
      })
      .then(res => res.json())
      .then(json => {
        if (json.error) {
          reject(Object.assign({}, json.errors, {
            _error: json.error,
          }))
        }
        resolve()
      })

    })
  }
}

Message.propTypes = {
  fields: PropTypes.object.isRequired,
  handleSubmit: PropTypes.func.isRequired,
  submitting: PropTypes.bool.isRequired
}

Message = reduxForm({
  form: 'message',
  fields: ['body', 'to']
})(Message)

export default connect(state => ({
  message: state.main.message,
}))(Message)
