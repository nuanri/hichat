import { api_get } from '../api'
import { c_account } from './constants'

function read_account_failed(error) {
  return {
    type: c_account.READ_FAILURE,
    error,
  }
}

function read_account_success(body) {
  return {
    type: c_account.READ_SUCCESS,
    body,
  }
}

export function readAccountInfo() {
  return dispatch => {
    return api_get('/auth/userinfo')
    // .then(res => res.json())
    .then(function(res) {
      return res.json()
    })
    .then(json => {
      console.log('json = ', json)
      dispatch(read_account_success(json))
    })
    .catch(error => dispatch(read_account_failed(error)))
  }
}
