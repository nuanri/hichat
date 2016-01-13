import { c_message } from './constants'
import { api_get } from '../api'


export function sendingMessage(data) {
  return {
    type: c_message.SENDING,
    data,
  }
}

function list_message_success(data) {
  return {
    type: c_message.LIST_SUCCESS,
    data,
  }
}

function list_message_failed(error) {
  return {
    type: c_message.LIST_FAILURE,
    error,
  }
}

export function listMessage() {
  return dispatch => {
    return api_get('/messages')
    // .then(res => res.json())
    .then(json => {
      console.log('json = ', json)
      dispatch(list_message_success(json))
    })
    .catch(error => dispatch(list_message_failed(error)))
  }
}
