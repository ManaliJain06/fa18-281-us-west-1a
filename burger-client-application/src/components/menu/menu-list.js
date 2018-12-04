/*
	UI Component to show all menu items
*/

import React, {Component} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as menuApi from './../../apis/menu-api';
import {updateMenuList} from './../../actions/menu-actions';
import '../../stylesheets/menu-list.css';
import '../../index.css';

class Menu extends Component{

  componentDidMount(){
    if(this.props.match && this.props.match.params && this.props.match.params.restaurantId){
      console.log("[Menu] componentDidMount restaurantId: ",this.props.match.params.restaurantId)
      menuApi.getRestaurantMenuItems(this.props.match.params.restaurantId).then((response)=>{
          if(response.status===200){
              response.json().then((data)=>{
                  console.log("[Menu] items: ", data);

                  this.props.updateMenuList(data);
              });
          }});

    }else{
      console.log("[Menu] componentDidMount !!! restaurantId missing !!!");
    }
  }

  getItems(items){
    return items.map((item)=>{
      return(
        <tr className = "menu-table-item-row">
            <td className = "menu-table-item-col">{item.name}</td>
            <td className = "menu-table-item-col">{item.description}</td>
            <td className = "menu-table-item-col">{item.calories}</td>
            <td className = "menu-table-item-col">{item.price}</td>
            <td className = "menu-table-item-col">
              <button> Add to cart </button>
            </td>
        </tr>
      )
    })
  }

  displayMenu(){
    if(this.props.menu.items && this.props.menu.items.length > 0){
      console.log("[Menu] displayMenuItems items: ",this.props.menu.items )
      return(
        <table className="table-menu">
          <tbody>
          <tr className = "menu-table-header-row">
              <th  className = "menu-table-item-col">Name</th>
              <th  className = "menu-table-item-col">Content</th>
              <th  className = "menu-table-item-col">Calories</th>
              <th  className = "menu-table-item-col">Price</th>
          </tr>
          {this.getItems(this.props.menu.items)}
          </tbody>
        </table>)

    }else{
      return null
    }
  }

render(){
  console.log("[Menu] render url param: ",this.props.match.params.restaurantId);
  return (
      <div className="menu-home">
      <div className="header">
          <div className="leftheader"> The Counter Custom burgers </div>
          <div className="rightheader">
              <div className="topnav">
                  <a >Home</a>
                  <a >Create Account</a>
                  <a >Login</a>
              </div>
          </div>
      </div>
          {this.displayMenu()}
      </div>
  )
}
}

function mapStateToProps(state) {
    console.log("[Menu] mapStateToProps");
    return{
        menu: state.menu
    }
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators({
        updateMenuList: updateMenuList
    }, dispatch)
}

export default connect(mapStateToProps, mapDispatchToProps)(Menu);
