import React, {Component} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import {withRouter} from 'react-router-dom';
import '../index.css';

class Header extends Component{


  handleOrderCartClick(){
    console.log("[Header] handleOrderCartClick")
    let orderId = localStorage.getItem('orderId');
    if(orderId){
      this.props.history.push("/order/"+orderId)
    }else{
      alert("No orders available in the cart !!!")
    }


  }
  showCart(){
    console.log("[Header]showCart: ",this.props.showCart)
    if(this.props.showCart && this.props.showCart.status){
      return (
          <span style = {{cursor:"pointer"}} onClick={()=>{this.handleOrderCartClick()}}>
              <i className="fa fa-shopping-cart">
                Cart Item: {this.props.order?this.props.order.orderCount:0}
              </i>
          </span>
      )
    }else{
      return null
    }
  }

  handleLogout = () => {
      localStorage.removeItem("user");
      this.props.history.push("/");
  };


  render(){
    console.log("[Header] render: ");
    let loginLogout;
    let user = localStorage.getItem("user");
    if(user != null){
        loginLogout = <span style = {{cursor:"pointer"}} onClick={()=>{this.props.history.push("/login")}}>Sign in</span>;
    }
    else{
        loginLogout = <span style = {{cursor:"pointer"}} onClick={()=>{this.handleLogout()}}>Logout</span>;
    }
    return (
      <div className="header">
          <div className="leftheader"> The Counter Custom burgers </div>
          <div className="rightheader">
              <div className="topnav">
                  <span style = {{cursor:"pointer"}} onClick={()=>{this.props.history.push("/")}}>Home</span>
                  <span style = {{cursor:"pointer"}} onClick={()=>{this.props.history.push("/signup")}}>Sign up</span>
                  {loginLogout}
                  {this.showCart()}

              </div>
          </div>
      </div>
    )
  }
}

function mapStateToProps(state) {
    console.log("[Header] mapStateToProps");
    return{
        order:state.order
    }
}



export default withRouter(connect(mapStateToProps, null)(Header));
