import React from 'react';
import './login.css';


function Login(){
    return(

        <div>

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
  <h4>New to ShowBase?<span> Sign up now.</span></h4>
</form>

</div>

    )
}

export default Login;