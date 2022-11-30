import axios from "axios";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

function Add() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [id, setID] = useState("");
  //const [state, setState] = React.useState();


  const navigate = useNavigate();
  const data = {
    name: name,
    email: email,
    phone: phone,
    id : id
  };


//const Page404 = () => <h1>Four:oh:four</h1>

  function submitForm(e) {
    e.preventDefault();
    axios.post("http://localhost:8080/user", data).then(navigate("/"))
    .catch(error => {
      navigate("/add-user")
    })
  }
  
  return (
    <div className="w-screen h-full flex flex-col justify-center items-center mt-16">
      <h2 className="text-2xl font-bold">Create Account</h2>
      <form className="w-[50%] h-full flex flex-col mt-2">
        <input
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-6 pl-6 mt-4"
          type="text"
          placeholder="Enter your name"
        />
        <input
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-6 pl-6 mt-4"
          type="email"
          placeholder="Enter your email"
        />
        <input
          value={phone}
          onChange={(e) => setPhone(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-6 pl-6 mt-4"
          type="phone"
          placeholder="Enter your phone no."
        />
         <input
          value={id}
          onChange={(e) => setID(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-6 pl-6 mt-4"
          type="id"
          placeholder="Enter your id"
        />
        <button
          className="bg-teal-600 outline-none font-bold border text-white border-zinc-400 py-4 pl-4 mt-4"
          type="submit"
          onClick={submitForm}
        >
          Create Account
        </button>
      </form>
    </div>
  );
}

export default Add;
