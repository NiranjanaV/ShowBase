import React from 'react';
import  './login.css';
import { Link } from 'react-router-dom';
import { useState } from 'react'


function Login(){

     // states for login 
     const [user,setUser] = useState('');
     const [pwd,setPwd] = useState('');
     const [error,setErrMsg] = useState(' ');

     const[success,setSuccess]= useState(false);

// value={user} is added to make it a controlled input - we need to clear the fields once login has been performed
 
     const handleLoginSubmit = async(e) => {
     e.preventDefault();
     console.log(user,pwd);
     setUser('');
     setPwd('');
     setSuccess(true);

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

<form onSubmit={handleLoginSubmit}>
  
<h2>Sign in</h2>
  <div class="input-container">
    <input required type="text" onChange={(e)=>setUser(e.target.value)} value ={user} />
    <label>Email</label>
  </div>
  
  <div class="input-container">
    <input required type="password" onChange={(e)=>setPwd(e.target.value)} value ={pwd}/>
    <label>Password</label>
  </div>
  <div class="input-container">
    <input type="submit" value="Submit" />
  </div>
  <h4>New to ShowBase?<span> <Link to='/signup'>Sign up now.</Link></span></h4>
</form>

</div>

    )
}

export default Login;