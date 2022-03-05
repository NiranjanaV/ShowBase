import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

import {BrowserRouter} from 'react-router-dom'


  

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
    <App />
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// Sample Integration Between front end and back end

// {
// const URL ="http://192.168.0.206:8080";
// getData();

// function getData(){
//   console.log("in get function");
// fetch(URL + '/search').then(response => 
//   response.json().then(data => ({
//       data: data,
//       status: response.status
//   })
// ).then(res => {
//   console.log(res.status, res.data);
//  // document.getElementById("minVal").innerHTML=res.data;

// }));}

// }