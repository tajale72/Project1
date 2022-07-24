import React from 'react';
//import axios from 'axios';

function Home(props) {
    return (
        <div className="home">
            <h1>hello {props.name}</h1>
            {props.children}
        </div>
      )
    
  }

export default Home;


