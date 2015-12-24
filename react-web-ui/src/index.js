import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import { Router, Route } from 'react-router'
import { createHistory } from 'history'
import { syncReduxAndRouter } from 'redux-simple-router'


import AuthSignUpForm from './components/signupform'
import App from './containers/app'

import configureStore from './store/configureStore'

const store = configureStore()
const history = createHistory()

syncReduxAndRouter(history, store)

console.log('App = ', App)
ReactDOM.render(
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={App}>
        <Route path="auth/signup" component={AuthSignUpForm}/>
      </Route>
    </Router>
  </Provider>,
  document.getElementById('root')
)
