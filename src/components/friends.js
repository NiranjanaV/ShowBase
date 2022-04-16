import React from 'react';
import  './login.css';
import axios from "axios"
import "./watch.css";
import {ip} from './global.js'

function Friends({friend}){

     // states for login 
     //const [watched,setWatched] = useState('');
     const id = friend;
     console.log(id);
// value={user} is added to make it a controlled input - we need to clear the fields once login has been performed
 
  const fr = async(e) => {
    console.log(e);
    
    const detailURL ="http://"+ip+":8080/putFriends/"
    //setWatched(1);
      try{
        e.preventDefault();
      
        const response = await axios.put(detailURL,
          JSON.stringify({username:"Srikant",
          friendname:"Swaathi"
      })

        );
        console.log(JSON.stringify(response));

      
      }

      catch{
        console.log("error");
      }
     }


    return(
        <div>

<form onSubmit={fr}>
  
  <div class="input-container">
    <input type="submit" value="Add friend" />
  </div>
  
</form>
 


    </div>
    
    )  
}
export default Friends;