import React from 'react';
import './login.css';

function Login(){
    return(

        <div>

<form>
<h2>Sign in</h2>
  <div class="input-container">
    <input required type="text" />
    <label>Name</label>
  </div>
  
  <div class="input-container">
    <input required type="text" />
    <label>Email</label>
  </div>
  <div class="input-container">
    <input type="submit" value="Submit" />
  </div>
</form>
</div>

    )
}

export default Login;