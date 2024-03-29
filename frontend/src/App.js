import { Route, Routes } from 'react-router-dom';
import './App.css';
import Bakery from './pages/bakery/bakery';
import Finance from './pages/finance/finance';
import Home from './pages/home/home';
import Consultants from './pages/consultants/consultants';
//import Consultants from './pages/Consultants';
import About from './pages/about/about';
import Navbar from './pages/Navbar';
import Add from './pages/user/Add';
import Edit from './pages/user/Edit';
import Users from './pages/user/Users';
import DataGridforConsultants from './pages/datagriddemo/datagriddemo';
import InputWithIcon from './pages/InputWithIcon/InputWithIcon';


function App() {
  return (
    <div className="App">
      <Routes>
      <Route  path="/datagrid" exact element={<DataGridforConsultants/>} />
      <Route  path="/userinput" exact element={<InputWithIcon/>} />
      <Route  path="/finance" exact element={<Finance/>} />
     <Route  path="/bakery" exact element={<Bakery/>} />
        <Route  path="/" exact element={<Home/>} />
        <Route  path="/consultants" exact element={<Consultants/>} />
        <Route  path="/about" exact element={<About/>} />
        <Route  path="/users/:id" exact element={<Users/>} />
        <Route  path="/add-user" exact element={<Add/>} />
        <Route  path="/edit-user/:id" exact element={<Edit/>} />
      </Routes>
    </div>
  );
}

export default App;
