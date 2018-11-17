import React, { Component } from 'react';
import { Route, withRouter, Switch } from 'react-router-dom';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import Menu from './components/menu/menu-list';
import './App.css';


class App extends Component {

  render() {
    return (
      <div className="App">
      <Switch>
          <Route path= "/" render = {() => (<Menu/>)}/>
      </Switch>
      </div>
    );
  }
}

export default withRouter(connect(null, null)(App));
