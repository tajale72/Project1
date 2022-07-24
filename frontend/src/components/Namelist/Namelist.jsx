import React from 'react';
// import axios from 'axios';


function Namelist() { 
    // let response
    // axios.get(`http://localhost:8080/user/romit`)
    //       .then(res => {
    //         const response = res.data;
    //         this.setState({ response});
    //       })
    
    const names = ['hello', 'names', 'hi']
    const persons = [
        {
            id:1,
            name : 'romit',
            age: 25,
            skill: 'react'
        },
        {
            id:1,
            name : 'subu',
            age: 24,
            skill: 'food technology'
        },

    ]
// const responses = response.map((res) => <h2>{res.name}</h2>)
const namelist = names.map((name, index) => <h2>{index} {name}</h2>)
const personlist = persons.map(person => <h2>my name is {person.name}.</h2>)
return  <div>{namelist}{personlist}</div>
}

export default Namelist