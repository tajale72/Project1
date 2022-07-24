import * as React from 'react';
import TextField from '@mui/material/TextField';


function Submit() {
    return (
        <>
       <TextField id="name" label="Name" variant="outlined" />
       <TextField id="age" label="Age" variant="outlined" />
       <TextField id="gpa" label="Gpa" variant="outlined" />
       <TextField id="address" label="Address" variant="outlined" />
       </>
    )
  }
  
export default Submit;