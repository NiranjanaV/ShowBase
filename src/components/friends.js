import React from 'react';
import  './login.css';
import axios from "axios"
import "./watch.css";
import {ip} from './global.js'
import useAuth from "../hooks/useAuth";

// Add Friends Functionality
// When we click button add friend button, users friend table will updated in the database 
function Friends({profile}){

     // states for login 
     //const [watched,setWatched] = useState('');
     const id = profile.Name;
   

     const { auth } = useAuth();
// value={user} is added to make it a controlled input - we need to clear the fields once login has been performed
 
  const fr = async(e) => {
   
    
    const detailURL ="http://"+ip+":8080/putFriends/"
    //setWatched(1);
      try{
        e.preventDefault();
      
        const response = await axios.put(detailURL,
          JSON.stringify({username: auth.user,
          friendname:id
      })

        );
       

      
      }

      catch{
        
      }
     }


    return(

      <div className="add">

<button
            className="button"
            onClick={fr}
          >
            Add 
          </button>


     
       </div>
    

    )  
}
export default Friends;