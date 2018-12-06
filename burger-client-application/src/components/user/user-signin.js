/*
	UI Component for Sign In page
*/
/*
	UI Component to show payment web page
*/
import React, {Component} from 'react';
import {bindActionCreators} from 'redux';
import {connect} from 'react-redux';
import {withRouter} from 'react-router-dom'
import '../../index.css';
import '../../stylesheets/payment.css';

import Header from '../header';

class UserSignIn extends Component {
    constructor(props) {
        super(props);
        this.state = {
            email:"",
            password:""
        };
        this.handleLogin = this.handleLogin.bind(this);
        this.handleBackButton = this.handleBackButton.bind(this);

    }

    componentDidMount = () => {
        console.log('componentDidMount ---')
        const user = localStorage.getItem('user');

        // check if the user is already signed in
        if (user != null) {
            this.props.history.push('/');
        }
    };

    handleLogin(event) {
        event.preventDefault();
        const payload = {
            email: this.state.email,
            orderId: this.state.password
        };
        this.props.PaymentCreate(data, this.props.history);
    }

    handleBackButton(event) {
        event.preventDefault();
        console.log(`handleButton go back home`);
        this.props.history.push('/');
    }

    render() {
        return (
            <div className="menu-home">
                <div className="outerdiv">
                    <Header/>
                    <div className="content payment">
                        <div className="card center">
                            <h2 id="center">Sign In</h2>
                            {
                                <div>
                                    <form>
                                        <label>Email</label>
                                        <input className="paymentInputText" type="text" id="email" name="Email" placeholder="Email"
                                            onChange={
                                                (event) => {
                                                    this.setState({
                                                        ...this.state,
                                                        email: event.target.value});
                                                }
                                            }
                                        />
                                        <label>Password</label>
                                        <input className="paymentInputText" type="password" id="pass" name="password" placeholder="Password"
                                               onChange={
                                                   (event) => {
                                                       this.setState({
                                                           ...this.state,
                                                           password: event.target.value});
                                                   }
                                               }
                                        />
                                    </form>
                                    <div className="btn-container">
                                        <input type="button" className="payment_button" value="Login to your account"
                                               onClick={this.handleLogin}/>
                                    </div>
                                </div>
                            }
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return {
        data: state.payment.data,
        orderDetail: state.payment.orderDetail,

    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        PaymentGetAll: () => {
            dispatch(paymentActions.axiosGetAll());
        },
        PaymentCreate: (data, router) => {
            dispatch(paymentActions.axiosCreatePayment(data, router));
        },
        PaymentGetOrderDetail: (orderId) => {
            dispatch(paymentActions.axiosGetOrder(orderId));
        },
    }
};


const connectedSignIn = withRouter(connect(mapStateToProps, mapDispatchToProps)(UserSignIn));

export default connectedSignIn;

