/*
	UI Component to list all the restaurants
*/
import React, {Component} from 'react';
import {withRouter} from 'react-router-dom';
import '../../index.css';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as restaurantApi from './../../apis/restaurant-api';
import {restaurantList} from './../../actions/restaurant-actions';
import Header from '../header';

class ListRestaurant extends Component{

    constructor(props) {
        super(props);
    }

    componentDidMount(){
        if(this.props.match && this.props.match.params && this.props.match.params.zipcode){
            restaurantApi.getRestaurants(this.props.match.params.zipcode).then((response)=>{
                if(response.status===200){
                    response.json().then((data)=>{
                        //action to save data
                        this.props.restaurantList(data);
                    });
                }});
        }else{
            console.log("error: zipcode not found");
        }
    }

    render(){
        {
            if(this.props.restaurants.restaurantList && this.props.restaurants.restaurantList.length>0){
                return (
                    <div className="outerdiv">
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

                                                <tr className = "row">
                                                    <td className = "streetaddress">{res.restaurantName} <br></br>
                                                        {res.AddressLine1} <br></br>
                                                        {res.city} {res.state}<br></br>
                                                        {res.AddressLine2} <br></br>
                                                        <div>"Phone: "</div> {res.phone} <br></br>
                                                    </td>
                                                    <td>{res.hours}</td>
                                                    <td>{res.distance}</td>
                                                    <td> <input type="submit" value="Order" onClick={() => {
                                                        this.props.history.push("/menu/"+res.restaurantId);}}/>
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
                    return null
                }
        }
    }
}

function mapStateToProps(state) {
    console.log("[Menu] mapStateToProps");
    return{
        restaurants: state.restaurant
    }
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators({
        restaurantList: restaurantList
    }, dispatch)
}

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(ListRestaurant));
