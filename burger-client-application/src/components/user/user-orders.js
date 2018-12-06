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
        this.props.history.push('/');
        const user = JSON.parse(localStorage.getItem('user'));
        if (user == null) {
            alert("You are not logged in!");
            this.props.history.push('/login');
        }

        UserAPI.callUserOrdersAPI(user.id)
            .then((response) => {
                if (response.status === 200) {
                    this.props.userOrdersList(response.data);
                }
        });
    }

    render() {
        {
            if (this.props.userOrdersData.userOrders && this.props.userOrdersData.userOrders.length > 0) {
                return (
                    <div className="outerdiv">
                        <Header showCart={{status: true}}/>
                        <div className="content">
                            <div className="card">
                                <table className="tableRestaurant">
                                    <tbody>
                                    <tr>
                                        <th>Name</th>
                                        <th>Open Hours</th>
                                        <th>Distance</th>
                                        <th>Option</th>
                                    </tr>
                                    {
                                        this.props.restaurants.restaurantList.map((res) => (

                                            <tr className="row">
                                                <td className="streetaddress">{res.restaurantName} <br></br>
                                                    {res.addressLine1} <br></br>
                                                    {res.city} {res.state}<br></br>
                                                    {res.addressLine2} <br></br>
                                                    <div>Phone:</div>
                                                    {res.phone} <br></br>
                                                </td>
                                                <td>{res.hours}</td>
                                                <td>{res.distance}</td>
                                                <td><input className="home-page-button" type="submit" value="Order"
                                                           onClick={() => {
                                                               this.props.history.push("/menu/" + res.restaurantId);
                                                           }}/>
                                                </td>
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
