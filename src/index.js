import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

import {BrowserRouter} from 'react-router-dom'

import { AuthProvider } from './context/AuthProvider';



  

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
    <AuthProvider>
    <App />
    </AuthProvider>
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// Sample Integration Between front end and back end

// {
// const URL ="http://"+ip+":8080";
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