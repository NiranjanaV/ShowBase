import React from 'react';
import  './login.css';
import axios from "axios"
import { Link} from 'react-router-dom';
import { useState, useContext } from 'react'
import AuthContext from '../context/AuthProvider';
import {ip} from './global.js'
import {useNavigate} from 'react-router-dom';
// import "./styles.css";


function Login(props){

     // states for login 
  
     //console.log(props);
    // console.log(props.history);
     const navigate = useNavigate();
    // let history = useHistory();
     const {setAuth} = useContext(AuthContext);
     const [user,setUser] = useState('');
     const [pwd,setPwd] = useState('');
     const [error,setErrMsg] = useState(' ');

     const[success,setSuccess]= useState(false);

// value={user} is added to make it a controlled input - we need to clear the fields once login has been performed
 
     const handleLoginSubmit = async(e) => {

      const detailURL ="http://"+ip+":8080/authenticateUser/";

      try{
        e.preventDefault();

        const response = await axios.put(detailURL,
          JSON.stringify({Username: user,Password: pwd})
        );

        console.log(JSON.stringify(response.data.Return));
        if(response.data.Return=='1'){
        setAuth({user,pwd});
        setUser('');
        setPwd('');
        console.log(user,pwd);
        setSuccess(true);
        navigate('/');
        }else{
        setSuccess(false);
        setErrMsg("Wrong Credentials please try again!")
        }
  
      
      }
      catch{

        console.log("error");

      }
    

     }


    return(
      <>

      {success ? (

      <div className="login">
        <h1> You are logged in. </h1>
        </div>

      ):(

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

<form className='login-form' onSubmit={handleLoginSubmit}>
  <h4>{error}</h4>
<h2>Sign in</h2>
  <div class="input-container">
    <input required type="text" onChange={(e)=>setUser(e.target.value)} value ={user} />
    <label>Username</label>
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
{/* <section id="footer">
    <div className="footer container">
      <div className="brand">
        <h1>ShowBase</h1>
      </div>
      <h2>Software Engineering Project</h2>
      <div className="social-icon">
        <div className="social-item">
          <a href="https://github.com/NiranjanaV/ShowBase"><img src="https://img.icons8.com/bubbles/100/000000/github.png" /></a>
        </div>
      </div>
    </div>
  </section> */}
</div>

)}
</>

    )
    
}

export default Login;