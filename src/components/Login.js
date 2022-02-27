import React from 'react';
import  './login.css';
import { Link } from 'react-router-dom';


function Login(){
    return(

        <div className="login">

<form>
<h2>Sign in</h2>
  <div class="input-container">
    <input required type="text" />
    <label>Email</label>
  </div>
  
  <div class="input-container">
    <input required type="text" />
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