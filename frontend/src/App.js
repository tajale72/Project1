
import { Component } from 'react';
import Home from './components/Home/Home'
import Welcome from './components/Welcome/Welcome';
import Namelist from './components/Namelist/Namelist';
// import Button from '@mui/material/Button';
// import Get from './components/Home/Home'



class App extends Component {
 
  render() {
    return (
      <div className="App">
       
          <Home/>
          <Welcome name="romit" super="superman" />
          <Namelist />
      
      </div>
    )
  }
}



export default App;
