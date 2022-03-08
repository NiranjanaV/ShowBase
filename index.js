import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './src/App';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter} from 'react-router-dom'

{
  getData();
  const URL ="192.168.0.206:8080";
  function getData(){
    console.log("in get function");
  fetch(URL + '/articles').then(response => 
    response.json().then(data => ({
        data: data,
        status: response.status
    })
  ).then(res => {
    console.log(res.status, res.data);
   // document.getElementById("minVal").innerHTML=res.data;
  
  }));
  }

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
    <App />
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
}
