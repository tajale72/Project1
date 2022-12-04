import axios from "axios";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Progressbar from './progressbar';
import ProgressStatus from './ProgressStatus';
import InputAdornment from '@mui/material/InputAdornment';
import TextField from '@mui/material/TextField';
import AccountCircle from '@mui/icons-material/AccountCircle';
import SendIcon from '@mui/icons-material/Send';
import Button from '@mui/material/Button';
import './User.css'

import testimage from './client1.jpg'
import { Upload } from "@mui/icons-material";

function Users() {
  const { id } = useParams();

  const [user, setUser] = useState([]);

  useEffect(() => {
    axios.get(`http://localhost:8080/user/${id}`).then((res) => {
      setUser(res.data);
    });
  }, []);

  const namestyle = {
		padding: 10,
	}

  console.log(user);
  return (
    <>

      <div className="h-full w-full flex flex-col mt-32 justify-center items-center">
          <div className="py-4 inline-block min-w-full sm:px-6 lg:px-8">
          <ProgressStatus/>
         <Progressbar bgcolor="#ff00ff" progress={user.status}  height={50} />
          </div>
      </div>

      <div class="user-input">
          <div class="column">
          <TextField
       style={namestyle}
        id="input-with-icon-textfield"
        label="Name"
        value={user.name}
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
        value={user.surname}
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
        value={user.email}
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
        value={user.phone}
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
        value={user.age}
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
        value={user.weight}
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
        value={user.maritalstatus}
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
        value={user.placeofbirth}
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
        value={user.nationality}
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
        value={user.address}
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
        value={user.sex}
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

    <div>
      <img src="http://localhost:3000/static/media/client1.5979b53460bc3716bdfa.jpg"></img>

    </div>


<div>
<Button 
      variant="outlined" 
      endIcon={<SendIcon />}>
      Continue
      </Button>
</div>
    


  
      </>
  );
}

export default Users;
