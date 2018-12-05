/*
	UI Component to show all menu items
*/

import React, {Component} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as menuApi from './../../apis/menu-api';
import * as orderApi from './../../apis/order-api';
import {updateMenuList} from './../../actions/menu-actions';
import {updateCart} from './../../actions/order-actions';
import '../../stylesheets/menu-list.css';
import '../../index.css';
import uuidv4 from "uuid";
import Header from '../header';

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

  addItem(item){
    console.log("[Menu] add Item:",item)
    var requestPayload = {}

    if(!localStorage.getItem('orderId')){
      requestPayload.orderId = uuidv4();
      localStorage.setItem('orderId',requestPayload.orderId );
    }else{
      requestPayload.orderId = localStorage.getItem('orderId');
    }
    requestPayload.itemName = item.name;
    requestPayload.itemId = item.id
    requestPayload.price = item.price
    requestPayload.description = item.description
    requestPayload.calories = item.calories


    //[ToDo] error handling needs to be done

    orderApi.addOrderItem(requestPayload).then((response)=>{
        if(response.status===200){
            response.json().then((data)=>{
                console.log("[Menu] Item Ordered: ", data);
                this.props.updateCart(data);
            });
        }});
  }

  getItems(items){
    return items.map((item)=>{
      return(
        <tr className = "menu-table-item-row">
            <td className = "menu-table-item-col">{item.name}</td>
            <td className = "menu-table-item-col">{item.description}</td>
            <td className = "menu-table-item-col">{item.calories}</td>
            <td className = "menu-table-item-col">$ {item.price}</td>
            <td className = "menu-table-item-col">
            <span onClick={()=>{this.addItem(item)}}> <i className="fas fa-cart-plus"
              style={{fontSize:24,color:"#72182a",cursor:"pointer"}}/>
            </span>

            </td>
        </tr>
      )
    })
  }

  displayMenu(){
    if(this.props.menu.items && this.props.menu.items.length > 0){
      console.log("[Menu] displayMenuItems items: ",this.props.menu.items )
      return(
        <div className = "menu-item-div">
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
        </table>
        </div>)

    }else{
      return  null
    }
  }

render(){
  console.log("[Menu] render url param: ",this.props.match.params.restaurantId);
  return (
      <div className="menu-home">
          <Header showCart={{status:true}}/>
          {this.displayMenu()}
      </div>
  )
}
}

function mapStateToProps(state) {
    console.log("[Menu] mapStateToProps");
    return{
        menu: state.menu,
        order:state.order
    }
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators({
        updateMenuList: updateMenuList,
        updateCart:updateCart
    }, dispatch)
}

export default connect(mapStateToProps, mapDispatchToProps)(Menu);
