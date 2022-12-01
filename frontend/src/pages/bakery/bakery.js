import React from 'react'
import logo from './images/back.png';
import  './bakery.css'




function Bakery() {
  return (
    <>
      <header class="header">
        <div class="logoContent">
            <a href="#" class="logo" width="10px"><img src={logo} alt=""/></a>
        </div>
        <div class="icon">
            <i class="fas fa-search" id="search"></i>
            <i class="fas fa-bars" id="menu-bar"></i>
        </div>
    </header>
       
    </>
  )
}

export default Bakery