import React from 'react';
import  './login.css';
import axios from "axios"
import "./watch.css";
import {ip} from './global.js'

// Add Friends Functionality
// When we click button add friend button, users friend table will updated in the database 
function Friends({profile}){

     // states for login 
     //const [watched,setWatched] = useState('');
     const id = profile.Name;
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
          friendname:id
      })

        );
        console.log(JSON.stringify(response));

      
      }

      catch{
        console.log("error");
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
      


      {/* <div className="container">
         <div className="add-content">
          <div className="input-wrapper">
             
           </div>
        
            <ul className="results">
            {results.map((movie) => (
                 <li key={movie.id}>
                   <ResultCard movie={movie} />
                </li>
              ))}
             </ul>
      
         </div>
       </div> */}
       </div>
    

    )  
}
export default Friends;