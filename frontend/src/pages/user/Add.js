import axios from "axios";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { CountryDropdown, RegionDropdown, CountryRegionData } from 'react-country-region-selector';

function Add() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [id, setID] = useState("");
  const [surname, setSurname] = useState("");
  const [placeofbirth, setPlaceofbirth] = useState("");
  const [age, setAge] = useState("");
  const [weight, setWeight] = useState("");
  const [nationality, setNationality] = useState("");
  const [maritalStatus, setMaritalstatus] = useState("");
  const [address, setAddress] = useState("");
  const [sex, setSex] = useState("");


  const navigate = useNavigate();
  const data = {
    name: name,
    email: email,
    phone: phone,
    id : id ,
    surname      :surname,
	  placeofbirth :placeofbirth,
	  age          :age,
	  weight       :weight,
	  nationality  :nationality,
	  maritalStatus :maritalStatus,
	  address      :address,
	  sex          :sex,
  };


  function submitForm(e) {
    e.preventDefault();
    axios.post("http://localhost:8080/user", data).then(navigate("/consultants"))
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
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="text"
          placeholder="Enter your name"
        />
         <input
          value={surname}
          onChange={(e) => setSurname(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="text"
          placeholder="Enter Last Name"
        />
        <input
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="email"
          placeholder="Enter your email"
        />
        <input
          value={phone}
          onChange={(e) => setPhone(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="phone"
          placeholder="Enter your phone no."
        />
              <input
          value={age}
          onChange={(e) => setAge(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="age"
          placeholder="Enter your age no."
        />
         <input
          value={weight}
          onChange={(e) => setWeight(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="weight"
          placeholder="Enter your weight"
        />
        <input
          value={maritalStatus}
          onChange={(e) => setMaritalstatus(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="maritalStatus"
          placeholder="Enter your maritalStatus"
        />
         <input
          value={placeofbirth}
          onChange={(e) => setPlaceofbirth(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="placeofbirth"
          placeholder="Enter your placeofbirth"
        />
        <input
          value={nationality}
          onChange={(e) => setNationality(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="nationality"
          placeholder="Enter your nationality"
        />
        <input
          value={address}
          onChange={(e) => setAddress(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="address"
          placeholder="Enter your address no."
        />
         <input
          value={sex}
          onChange={(e) => setSex(e.target.value)}
          className="bg-white/10 outline-none font-normal border border-zinc-400 py-2 pl-6 mt-4"
          type="sex"
          placeholder="sex"
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
