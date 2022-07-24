import React from 'react';
import axios from 'axios';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea } from '@mui/material';
//import { loginRequest } from "./Getusername";


export default class PersonList extends React.Component {
  state = {
    persons: []
  }

  componentDidMount() {
    const hostname = process.env("BACKEND_HOST")
    axios.get(`http://${hostname}:8080/user/romit`)
      .then(res => {
        const persons = res.data;
        this.setState({ persons });
      })
  }

  render() {
    return (
      <>
<div>
<Card sx={{ maxWidth: 345 }}>
      <CardActionArea>
        <CardMedia
          component="img"
          image="https://images.unsplash.com/photo-1551963831-b3b1ca40c98e"
        />
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
           User Info
          </Typography>
          <Typography variant="body2" color="text.secondary">
          <p>Name: {this.state.persons.map(person => person.name)}</p>
          <p>Age: {this.state.persons.map(person => person.age)}</p>
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
</div>
      </>
      
    )
  }
}