import React, { Component } from 'react';
import ReactDOM, {createRoot} from 'react-dom/client'

import AppFooter from './AppFooter';
import AppContent from './AppContent';

import 'bootstrap/dist/css/bootstrap.min.css'
import './index.css'

class App extends Component {
  // every react component must have a render() function
  render() {
    return (
      <div className="app">
        <div>
          <h1>hello?</h1>
          <AppContent />
        </div>
        <AppFooter />
      </div>
    );
  }
}

//ReactDOM.render(<App />, document.getElementById('root'));
const root = createRoot(document.getElementById('root'));
root.render(<App />);