import React from 'react'
import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Progressbar from './progressbar';
import ProgressStatus from './ProgressStatus';
import InputAdornment from '@mui/material/InputAdornment';
import TextField from '@mui/material/TextField';
import AccountCircle from '@mui/icons-material/AccountCircle';

import  './finance.css'

function Finance() {

  const [firstname, setName] = useState("");
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



  const data = {
    name: firstname,
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


  const namestyle = {
		padding: 10,
	}

  return (
    <>


          <ProgressStatus/>
         <Progressbar bgcolor="#ff00ff" progress={'25'}  height={50} />
    

      <div class="input">
          <div class="column">
      <TextField
       style={namestyle}
        id="input-with-icon-textfield"
        label="Name"
        value={firstname}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
      <TextField
       style={namestyle}
        id="input-with-icon-textfield"
        label="Lastname"
        value={surname}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
      <TextField
       style={namestyle}
        id="input-with-icon-textfield"
        label="Email"
        value={email}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
      <TextField
       style={namestyle}
        id="input-with-icon-textfield"
        label="phone"
        value={phone}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
      <TextField
       style={namestyle}
        id="input-with-icon-textfield"
        label="Age"
        value={age}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
          </div>
          <div class="column">
          <TextField
        style={namestyle}
        id="input-with-icon-textfield"
        label="weight"
        value={weight}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
       <TextField
        style={namestyle}
        id="input-with-icon-textfield"
        label="maritalstatus"
        value={maritalStatus}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
       <TextField
        style={namestyle}
        id="input-with-icon-textfield"
        label="placeofbirth"
        value={placeofbirth}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
       <TextField
        style={namestyle}
        id="input-with-icon-textfield"
        label="nationality"
        value={nationality}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
       <TextField
        style={namestyle}
        id="input-with-icon-textfield"
        label="address"
        value={address}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
         <TextField
          style={namestyle}
        id="input-with-icon-textfield"
        label="sex"
        value={sex}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <AccountCircle />
            </InputAdornment>
          ),
        }}
        variant="standard"
      />
          </div>
      </div>




{/* <div class="home-image">
<img src="https://images.pexels.com/photos/251225/pexels-photo-251225.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"></img>
</div>
   
           <div>
           <p>“You can’t connect the dots looking forward; you can only connect them looking backwards. 
            So you have to trust that the dots will somehow connect in your future.
             You have to trust in something – your gut, destiny, life, karma, whatever. 
             This approach has never let me down, and it has made all the difference in my life.”Steve Jobs"</p>
            </div> */}
    </>
  )
}

export default Finance

