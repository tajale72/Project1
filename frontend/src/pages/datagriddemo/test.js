import React , {useEffect , useState} from 'react';

import {DataGrid} from "@mui/x-data-grid";
import axios from 'axios';
 

function DataGridforConsultants() {
  const [data, setData] = useState([]);

  const getConsultantsData = async() => {
    await axios.get("http://localhost:8080/user").then((res) =>{
      setData(res.data.data);
    })
  }

  useEffect(() => {
    getConsultantsData();
  }, []);

  const columns = [
    {field : "_id" , headerName : "ID",  width: 90},
    {field : "name" , headerName : "NAME",  width: 90},
    {field : "email" , headerName : "EMAIL",  width: 90},
    {field : "age" , headerName : "AGE",  width: 90},
    {field : "address" , headerName : "ADDRESS",  width: 90},

  ];

  const rows = data.map((row) => ({
    _id : row._id,
    name : row.name,
    email : row.email,
    age : row.age,
    address : row.address
  }));



  console.log("data", data)

  return <div style={{height : 500, width: "100%"}}>
     <DataGrid
     rows = {rows}
     columns = {columns}
     pageSize = {10}
     rowsPerPageOptions = {[10]}
     />
  </div>
}

export default DataGridforConsultants;