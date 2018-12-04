/*
	UI Component to list all the restaurants
*/
import React, {Component} from 'react';
import {withRouter} from 'react-router-dom';
import '../../index.css';

class ListRestaurant extends Component{

    constructor(props) {
        super(props);
    }

    render(){
        return (
            <div className="menu-home">
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
                                <tr className="row">
                                    <td>
                                        <div className="streetaddress">
                                            3055 Olin Avenue<br></br>
                                            San Jose,
                                            CA<br></br>
                                            Santana Row<br></br>
                                            Phone: (408) 423-9200
                                            <br></br>
                                        </div>
                                    </td>
                                    <td>
                                        Open today
                                        11am-10pm
                                        <br></br>
                                    </td>
                                    <td>
                                        7.8 mi.
                                    </td>
                                    <td>
                                        <input type="submit" value="Order" onClick={() => {
                                            this.props.history.push("/menu/1");}}/>
                                    </td>
                                </tr>

                                <tr className="row">
                                    <td>
                                        <div className="streetaddress">
                                            4000 Cross Avenue<br></br>
                                            San Jose,
                                            CA<br></br>
                                            Santana Row<br></br>
                                            Phone: (669) 423-9200
                                            <br></br>
                                        </div>
                                    </td>
                                    <td>
                                        Open today
                                        9am-9pm
                                        <br></br>
                                    </td>
                                    <td>
                                        1.2 mi.
                                    </td>
                                    <td>
                                        <input type="submit" value="Order" onClick={() => {
                                            this.props.history.push("/menu/1");}}/>
                                    </td>
                                </tr>

                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default withRouter(ListRestaurant);
