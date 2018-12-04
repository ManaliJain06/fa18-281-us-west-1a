import React, {Component} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import '../index.css';

class Header extends Component{


  handleOrderCartClick(){

  }


  render(){
    console.log("[Header] render: ");
    return (
      <div className="header">
          <div className="leftheader"> The Counter Custom burgers </div>
          <div className="rightheader">
              <div className="topnav">
                  <a >Home</a>
                  <a >Create Account</a>
                  <a >Login</a>
                  <span onClick={()=>{this.handleOrderCartClick()}}>
                    <i className="fa fa-shopping-cart" style={{fontSize:24}}>
                      Car Item: {this.props.order.orderCount}
                    </i></span>
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



export default connect(mapStateToProps, null)(Header);
