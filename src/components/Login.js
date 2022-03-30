import React from 'react';
import  './login.css';
import { Link } from 'react-router-dom';


function Login(){
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

<form>
  
<h2>Sign in</h2>
  <div class="input-container">
    <input required type="text" />
    <label>Email</label>
  </div>
  
  <div class="input-container">
    <input required type="password" />
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