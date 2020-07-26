import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import Fib from './Fib.js';

function App() {
  return (
    <Router>
      <div className="App">
        <div>
          <Route exact path="/" component={Fib} />
        </div>
      </div>
    </Router>
  );
}

export default App;
