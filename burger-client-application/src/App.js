import React, { Component } from 'react';
import { Route, withRouter, Switch } from 'react-router-dom';
import {connect} from 'react-redux';
import Menu from './components/menu/menu-list';
import Homepage from './components/homepage';
import Order from './components/order/order-list';
import './App.css';
import ListRestaurant from "./components/restaurant/restaurant-list";
import PaymentOverview from "./components/payment/payment";
import PaymentSuccess from './components/payment/paymentSuccess';
import PaymentError from './components/payment/paymentError';
import UserSignIn from './components/user/user-signin'
import UserSignUp from './components/user/user-signup'

class App extends Component {

  render() {
    return (
      <div className="App">
          <Switch>
            <Route exact path= "/" render = {() => (<Homepage/>)}/>
            <Route exact path= "/menu/:restaurantId" render = {(match) => (<Menu {...match} showCart={{status:true}}/>)}/>
            <Route exact path= "/order/:orderId" render = {(match) => (<Order {...match} />)}/>
            <Route exact path="/listRestaurant/:zipcode" render={(match) => (<ListRestaurant {...match}/>)}/>
            <Route path="/payment" render={() => (<PaymentOverview />) } />
            <Route path="/paymentSuccess" render={() => (<PaymentSuccess />) } />
            <Route path="/paymentError" render={() => (<PaymentError />) } />
            <Route path="/login" render={() =>(<UserSignIn />)}/>
            <Route path="/signup" render={() =>(<UserSignUp />)}/>

          </Switch>
      </div>
    );
  }
}

export default withRouter(connect(null, null)(App));
