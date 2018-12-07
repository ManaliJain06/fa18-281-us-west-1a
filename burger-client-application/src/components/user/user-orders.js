/*
	UI Component to list all the restaurants
*/
import React, {Component} from 'react';
import {withRouter} from 'react-router-dom';
import '../../index.css';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as UserAPI from './../../apis/user-api';
import {userOrdersList} from './../../actions/user-actions';
import Header from '../header';

class ListUserOrders extends Component {

    constructor(props) {
        super(props);
    }

    componentDidMount() {
        const user = JSON.parse(localStorage.getItem('user'));
        if (user == null) {
            alert("You are not logged in!");
            this.props.history.push('/login');
        }

        UserAPI.callUserOrdersAPI(user.id)
            .then((response) => {
                if (response.status === 200) {
                    let userOrders= response.data.filter(order => order.userid === user.id);
                    this.props.userOrdersList(userOrders);
                }
                else{
                    this.props.userOrdersList({});
                }
             })
            .catch((err) =>{
                if(err.status === 404){
                    console.log("data not found");
                }
            });
    }

    // filterData = (orders, userId) =>{
    //     let filteredData = orders.filter(order => order.userid === userId);
    //     return orders.filter(order => order.userid === userId);
    // };

    render() {
        {
            if (this.props.userOrdersData.ordersList && this.props.userOrdersData.ordersList.length > 0) {
                return (
                    <div className="outerdiv">
                        <Header showCart={{status: false}}/>
                        <div className="content">
                            <div className="card">
                                <table className="tableRestaurant" cellPadding={10}>
                                    <tbody>
                                    <tr>
                                        <th>Order Id</th>
                                        <th>Payment Id</th>
                                        <th>Amount</th>
                                        <th>Date</th>
                                    </tr>
                                    {
                                        this.props.userOrdersData.ordersList.map((res) => (

                                            <tr className="row">
                                                <td className="streetaddress">{res.orderid} <br></br></td>
                                                <td>{res.paymentid}</td>
                                                <td>{res.totalamount}</td>
                                                <td>{res.paymentdate}</td>
                                            </tr>
                                        ))
                                    }
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                )
            } else {
                return (
                    <Header showCart={{status: true}}/>
                )
            }
        }
    }
}

function mapStateToProps(state) {
    console.log("[user] mapStateToProps");
    return {
        userOrdersData: state.userOrders
    }
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators({
        userOrdersList: userOrdersList
    }, dispatch)
}

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(ListUserOrders));
