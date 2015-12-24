import fetch from 'isomorphic-fetch';

function api_request(method, url, options={}) {
  let sid = localStorage.getItem('sid');

  let myHeaders = new Headers();
  myHeaders.append('Sid', `${sid}`)
  // myHeaders.append('Accept', 'application/json')
  // myHeaders.append('Content-Type', 'application/json')
  myHeaders.append('Keep-Alive', true);
  options.method = method
  options.headers = myHeaders

  // let full_url = `https://appclouds.cn/v1/api${url}`
  let full_url = `/api${url}`;

  return fetch(full_url, options)
}

function api_get(url, options={}) {
  return api_request('GET', url, options)
}

function api_post(url, options={}) {
  return api_request('POST', url, options)
}

function api_put(url, options={}) {
  return api_request('PUT', url, options)
}

function api_delete(url, options={}) {
  return api_request('DELETE', url, options)
}

export { api_request, api_get, api_post, api_put, api_delete }
