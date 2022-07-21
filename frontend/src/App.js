import Button from '@mui/material/Button';
import SendIcon from '@mui/icons-material/Send';
import PersonList from './components/PersonList'
import ReactDOM from 'react-dom/client';

function Get() {
  const root = ReactDOM.createRoot(document.getElementById('users'));
  root.render(
  <PersonList />
);
}

function App() {
 
  return (
    <>
    <div className="App">
       <Button onClick={Get} variant="outlined" endIcon={<SendIcon />}>GET</Button> &nbsp;
       <Button variant="outlined" endIcon={<SendIcon />}>POST</Button> &nbsp;
       <Button variant="outlined" endIcon={<SendIcon />}>UPDATE</Button> &nbsp;
       <Button variant="outlined" endIcon={<SendIcon />}>DELETE</Button> &nbsp;
    </div>
    <br></br> <br></br>
    </>
  );
}



export default App;
