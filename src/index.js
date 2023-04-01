import React, { Component } from 'react';
import {createRoot} from 'react-dom/client'

import AppHeader from './AppHeader';
import AppFooter from './AppFooter';
import AppContent from './AppContent';

import 'bootstrap/dist/css/bootstrap.min.css'
import './index.css'

class App extends Component {

  constructor(props) {
    super(props);
    this.handlePostChange = this.handlePostChange.bind(this);
    this.state = {posts: []};
  }

  handlePostChange(posts) {
    this.setState({posts: posts});
  }

  // every react component must have a render() function
  render() {
    const myProps = {
      title: "Hola!!",
      subject: "123"
    };
    return (
      <div className="app">
        <AppHeader {...myProps} posts={this.state.posts} handlePostChange={this.handlePostChange}/>
         <AppContent handlePostChange = {this.handlePostChange} posts={this.state.posts}/>
         <AppFooter/>
      </div>
    );
  }
}

//ReactDOM.render(<App />, document.getElementById('root'));
const root = createRoot(document.getElementById('root'));
root.render(<App />);