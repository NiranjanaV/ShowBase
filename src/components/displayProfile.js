import React, { useContext } from "react";
import { GlobalContext } from "../context/GlobalState";
import Moment from "react-moment";
import "./watch.css";
import Friends from './friends';


export const DisplayProfile = ({profile}) => {
return (
    <div className="resultscard">

      <div className="movieinfo">
        <div className="header">
          <h3 className="movietitle">{profile.Name}</h3>
        </div>

        <Friends profile={profile} />

        </div>
      </div>
    
  );
};