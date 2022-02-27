import React from 'react';

function SignUp(){

    return(
        <div className="login">

        <form>
        <h2>Sign up to ShowBase</h2>
          <div class="input-container">
            <input required type="text" />
            <label>Email</label>
          </div>
          
          <div class="input-container">
            <input required type="password" />
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