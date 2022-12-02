import React from 'react'
import { Link } from 'react-router-dom'
import  './Navbar.css'

function Navbar() {
  return (
    <>
        <div className="w-ful h-16 flex items-center px-10 justify-between bg-white-600">
        <div className="text-3xl text-black font-Montesarrat">T  FIRM</div>
        <Link to="/" className="link">Test</Link>
        <Link to={"/"} className="hover:border-2 hover:border-white hover:text-red hover:shadow-md rounded-lg bg-white text-black">Home</Link>
        <Link to={"/finance"} className=" hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white text-black">Finance</Link>
        {/* <Link to={"/bakery"} className=" hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white text-black">Bakery</Link> */}
        <Link to={"/consultants"} className=" hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white text-black">Consultants</Link>
        <Link to={"/about"} className=" hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white text-black">About Us</Link>
        <Link to={"/add-user"} className=" hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white  text-black">Create Account</Link>
        </div>
    </>
  )
}

export default Navbar