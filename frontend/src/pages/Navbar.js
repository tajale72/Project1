import React from 'react'
import { Link } from 'react-router-dom'
import  './Navbar.css'
import lgo from './Tfirm_logo.png'

function Navbar() {
  const navbarcolor = {
    height : 60,
		backgroundColor: 'rgb(168, 225, 244)',
	}

  return (
    <>
        <div className="h-16 flex items-center px-10 justify-between m-12" style={navbarcolor}>
        <Link to="/" class="logo"><img src={lgo} alt=""/></Link>
        <Link to={"/datagrid"} className="link"><h1>Consultants</h1></Link>
        <Link to={"/finance"} className="link"><h1>Finances</h1></Link>
        <Link to={"/consultants"} className="link"><h1>Consultants</h1></Link>
        <Link to={"/about"} className="link"><h1>About Us</h1></Link>
        <Link to={"/add-user"} className="link"><h1>Create Account</h1></Link>
        </div>
    </>
  )
}

export default Navbar