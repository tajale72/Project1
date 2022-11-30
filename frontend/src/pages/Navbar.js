import React from 'react'
import { Link } from 'react-router-dom'

function Navbar() {
  return (
    <>
        <div className="w-ful h-16 flex items-center px-10 justify-between bg-white-600">
        <Link to={"/"} className="hover:bg-teal-600 hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white text-black">Home</Link>
        <Link to={"/finance"} className="hover:bg-teal-600 hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white text-black">Finance</Link>
        <Link to={"/bakery"} className="hover:bg-teal-600 hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white text-black">Bakery</Link>
        <Link to={"/consultants"} className="hover:bg-teal-600 hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white text-black">Consultants</Link>
        <div className="text-3xl text-black font-Montesarrat">Consulting | Portal</div>
        <Link to={"/add-user"} className="hover:bg-teal-600 hover:border-2 hover:border-white hover:text-teal-200 hover:shadow-md rounded-lg bg-white  text-black">Create Account</Link>
        </div>
    </>
  )
}

export default Navbar