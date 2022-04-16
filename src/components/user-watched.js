import React from 'react';
import  './login.css';
import axios from "axios"
import { Link } from 'react-router-dom';
import { useState, useContext } from 'react'
import AuthContext from '../context/AuthProvider';
import {ip} from './global.js'
import {FaCheck} from "react-icons/fa"

function UserWatched(props){

     // states for login 
     
     //const [watched,setWatched] = useState('');
     const id = props.movie;

// value={user} is added to make it a controlled input - we need to clear the fields once login has been performed
 
     const userWatched = async(e) => {

    const detailURL ="http://"+ip+":8080/updateUserLike/"

    //setWatched(1);

      try{
        e.preventDefault();

        const response = await axios.put(detailURL,
          JSON.stringify({username:"Srikant",
          movie:id,
          action:2,
          value:1
      })
          // ,
          // {
          // headers:{'Content-type':'application/json'},
          // withCredentials:true
          // }
        );

        console.log(JSON.stringify(response));
      

      }
      catch{

        console.log("error");

      }
    

     }


    return(
      

    

<div>
    
<form onSubmit={userWatched}>
  
  <div class="input-container">
 

    <button type="submit"><FaCheck/> Watched it? </button>

   
  </div>
  
</form>

</div>
)


    
}

export default UserWatched;