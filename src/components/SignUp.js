import React from 'react';
import { Link } from 'react-router-dom';
import axios from "axios"
import { useState } from 'react'


function SignUp(){

  const [user,setUser] = useState('');
  const [pwd,setPwd] = useState('');

  const handleRegSubmit = async(e) => {

    const detailURL ="http://70.171.43.6:8080/userReg/";

    try{
      e.preventDefault();
      console.log(user,pwd);

      const response = await axios.put(detailURL,
        JSON.stringify({Username: user,Password: pwd})
        // ,
        // {
        // headers:{'Content-type':'application/json'},
        // withCredentials:true
        // }
      );

      console.log(JSON.stringify(response));
      //setAuth(user,pwd);
      setUser('');
      setPwd('');
     
     // setSuccess(true);
     console.log("registered");

    }
    catch{

      console.log("error");

    }
  

   }



    return(
        <div className="login">
<section id="header">
       <div className="header container">
    <div className="nav-bar">
      <div className="brand">
        <a href="#hero">
        <Link to='/'><h1><span>S</span>how <span>B</span>ase</h1></Link>
        </a>
      </div>
    </div>
  </div>
  </section>
        <form  onSubmit={handleRegSubmit}>
        <h2>Sign up to ShowBase</h2>
          <div class="input-container">
            <input required type="text" onChange={(e)=>setUser(e.target.value)} value ={user}/>
            <label>Email</label>
          </div>
          
          <div class="input-container">
            <input required type="password"  onChange={(e)=>setPwd(e.target.value)} value ={pwd}/>
            <label>Password</label>
          </div>


          <div class="input-container">
            <input required type="password" />
            <label>Reenter your Password</label>
          </div>


          <div class="input-container">
            <input type="submit" value="Submit" />
          </div>
         
        </form>
        
        </div>
    )
}

export default SignUp;