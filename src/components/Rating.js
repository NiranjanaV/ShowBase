import React from "react";
import { FaStar } from 'react-icons/fa';
import {ip} from './global.js';
import axios from "axios";
import { useState } from 'react';

const Rating = (props) =>{

    const id = props.movie;

   const [vote,setVote] = useState('');
   const [hover,setHover] = useState('');
   
   const prev = props.prev;
   console.log(prev);

    const handleVoteSubmit = async(val) => {

        console.log(val);
        setVote(val);
        


        const detailURL ="http://"+ip+":8080/updateUserLike/"
  
        try{
         // e.preventDefault();
  
          const response = await axios.put(detailURL,
            JSON.stringify({username:"Srikant",
                movie:id,
                action:1,
                value:parseInt(val, 10)
            })
           
          );
  
          console.log(JSON.stringify(response));
         
         
          
        }
        catch{
  
          console.log("error");
  
        }
    
  
       }

    return(
        
        <div className="star-container">
             <div className="caption">Rate :</div>
            {[...Array(10)].map((star,i)=>{
            
            const ratingValue = i+1;
                return(
            <label>
            <input type="radio" name="rating" onChange={()=>handleVoteSubmit(ratingValue)} value={ratingValue}/>
            <FaStar className="rating-stars" color={ratingValue<= (hover || vote || prev) ? "#ffc107": "#e4e5e9"  } size={38} onMouseEnter={()=>setHover(ratingValue)} onMouseLeave={()=>setHover(null)}/>
            </label>
                );
            })}
        </div>
       
    )

};

export default Rating;
