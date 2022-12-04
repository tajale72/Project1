import * as React from 'react';
import Box from '@mui/material/Box';
import { DataGrid } from '@mui/x-data-grid';
import {useEffect , useState} from 'react';
import axios from 'axios';
import SendIcon from '@mui/icons-material/Send';
import Button from '@mui/material/Button';
import DeleteIcon from '@mui/icons-material/Delete';
import './datagrid.css';


const columns = [
  { 
    field: 'id',
    headerName: 'ID', 
    width: 90 },
  {
    field: 'name',
    headerName: 'name',
    headerAlign: 'center',
    width: 150,
    editable: true,
  },
  {
    field: 'surname',
    headerName: 'Last name',
    headerAlign: 'center',
    width: 150,
    editable: true,
  },
  {
    field: 'age',
    headerName: 'Age',
    type: 'number',
    headerAlign: 'center',
    width: 110,
    editable: true,
  },
  {
    field: 'address',
    headerName: 'Address',
    description: 'This column has a value getter and is not sortable.',
    headerAlign: 'center',
    sortable: false,
    width: 160,
  },
  {
    field: 'actions',
    headerName: 'Acrions',
    headerAlign: 'center',
    type :  'tags',
    sortable: false,
    width: 160,
  },
];



export default function DataGridDemo() {

  const [data, setData] = useState([]);

  const getConsultantsData = async() => {
    await axios.get("http://localhost:8080/user").then((res) =>{
      setData(res.data.reverse());
    })
  }

  useEffect(() => {
    getConsultantsData();
  }, []);


    const rows = data.map((row) => ({
      id : row._id,
      name : row.name,
      surname : row.surname,
      age : row.age,
      address : row.address
    }));
  
  

  console.log("data", data)

  return (
  
  <>
    <Box sx={{ height: 600, width: '100%' , mx : 'auto' , paddingTop: 20 }}>
      <h1>Manage Consultants</h1>
      <DataGrid
      sx={{
        boxShadow: 2,
        border: 2,
        fontSize: 16,
        borderColor: 'rgb(168, 225, 244)',
        '& .MuiDataGrid-cell:hover': {
          color: 'primary.main',
        },
      }}
        rows={rows}
        getRowClassName={(params) =>
          params.indexRelativeToCurrentPage % 2 === 0 ? 'even' : 'odd'
        }
        columns={columns}
        pageSize={5}
        rowsPerPageOptions={[5]}
        checkboxSelection
        disableSelectionOnClick
        experimentalFeatures={{ newEditingApi: true }}
      />
    </Box>




   <div class="datagrid-button">
   <Button 
      variant="outlined" 
      endIcon={<SendIcon />}>
      View
      </Button>
      <Button 
      variant="outlined" 
      endIcon={<SendIcon />}>
      Edit
      </Button>
      <Button 
      variant="outlined" 
      startIcon={<DeleteIcon />}>
        Delete
      </Button>

   </div>
    
</>
    
  );
}