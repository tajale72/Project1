import './App.css';
import Create from './components/create';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
function App() {
  return (
   
      <div className="main">
        <h2 className="main-header">React Crud Operations</h2>
        <BrowserRouter>
        <Switch>
          <Route path="/create">
            <Create />
          </Route>
        </Switch>
      </BrowserRouter>
      </div>
  
    
  );
}

export default App;