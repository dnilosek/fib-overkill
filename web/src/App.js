import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import OtherPage from './OtherPage.js';
import Fib from './Fib.js';

function App() {
  return (
    <Router>
    <div className="App">
      <header>
	<Link to="/">Home Updated</Link>
      </header>
      <div>
        <Route exact path="/" component={Fib} />
	<Route path="/otherpage" component={OtherPage} />
      </div>
    </div>
    </Router>
  );
}

export default App;
