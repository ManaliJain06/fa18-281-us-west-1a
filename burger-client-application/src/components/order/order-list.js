/*
	UI Component to show order
*/

/*
	UI Component to show all menu items
*/

import React, {Component} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as orderApi from './../../apis/order-api';
import {updateCart} from './../../actions/order-actions';
import '../../stylesheets/menu-list.css';
import uuidv4 from "uuid";
import Header from '../header';

class Order extends Component{

  componentDidMount(){
    if(this.props.match && this.props.match.params && this.props.match.params.orderId){
      console.log("[Order] componentDidMount orderId: ",this.props.match.params.orderId)
      orderApi.getOrderItems(this.props.match.params.orderId).then((response)=>{
          if(response.status===200){
              response.json().then((data)=>{
                  console.log("[Order] order: ", data);

                  this.props.updateCart(data);
              });
          }});

    }else{
      console.log("[Order] componentDidMount !!! restaurantId missing !!!");
    }
  }

  deleteItem(item){
    let orderId = localStorage.getItem('orderId');
    console.log("[Order] deleteItem item ", item, "  orderId: ",orderId);
    if(orderId){
      orderApi.deleteOrderItem(orderId, item.itemId).then((response)=>{
          if(response.status===200){
              response.json().then((data)=>{
                  console.log("[Order] deleteItem success data: ", data);

                  this.props.updateCart(data);
              });
          }});

    }else{
      alert("Could not find a existing order !!!")
    }

  }

  getOrderItems(items){
      return items.map((item)=>{
        return(
          <tr className = "menu-table-item-row">
              <td className = "menu-table-item-col">{item.itemName}</td>
              <td className = "menu-table-item-col">{item.description}</td>
              <td className = "menu-table-item-col">{item.price}</td>
              <td className = "menu-table-item-col">
                <span onClick={()=>{this.deleteItem(item)}}><i className="fa fa-close"
                  style={{fontSize:24,color:"#72182a",cursor:"pointer"}}/>
                </span>
              </td>
          </tr>
        )
      })


  }

  displayOrder(){
    if(this.props.order.items && this.props.order.items.length > 0){
      console.log("[Order] displayOrderItems items: ",this.props.order.items )
      return(
        <table className="table-menu">
          <tbody>
          <tr className = "menu-table-header-row">
              <th  className = "menu-table-item-col">Name</th>
              <th  className = "menu-table-item-col">Content</th>
              <th  className = "menu-table-item-col">Price</th>
          </tr>
          {this.getOrderItems(this.props.order.items)}
          <tr className = "order-table-total-row">
            <td className="order-table-total-amount"><span><i></i></span></td>
            <td className = "order-table-total-text">Total Amount </td>
            <td className = "order-table-total-amount"> {this.props.order && this.props.order.totalAmount?this.props.order.totalAmount:0}</td>
          </tr>
          </tbody>

        </table>)

    }else{
      return (<h2> No items in the cart !!! </h2>)
    }
  }

render(){
  console.log("[Order] render url param: ",this.props.match.params.orderId);
  return (
      <div className="menu-home">
          <Header/>
          {this.displayOrder()}
      </div>
  )
}
}

function mapStateToProps(state) {
    console.log("[Order] mapStateToProps");
    return{
        order:state.order
    }
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators({
        updateCart:updateCart
    }, dispatch)
}

export default connect(mapStateToProps, mapDispatchToProps)(Order);
