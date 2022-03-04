import React from 'react';
import { Link } from 'react-router-dom';
import { useLocation } from 'react-router-dom'
import { useEffect, useState } from "react"

function MovieDetail(props){


  const location = useLocation()
  const { from } = location.state
  console.log(from);

    return(


    <div>
      <p></p>
   
    </div>

    )
}

export default MovieDetail;