
import React, {Component} from 'react';
// import {bindActionCreators} from 'redux';
// import {connect} from 'react-redux';
import {withRouter} from 'react-router-dom'
import '../../index.css';
import '../../stylesheets/payment.css';
import * as UserAPI from '../../apis/user-api'
import Header from '../header';

class UserSignUp extends Component {
    constructor(props) {
        super(props);
        this.state = {
            firstName:"",
            lastName:"",
            street: "",
            city: "",
            state: "",
            zipcode: "",
            password: "",
            email: ""
        };
        //this.handleRegister = this.handleRegister.bind(this);
    }

    componentDidMount = () => {
        console.log('componentDidMount ---');
        const user = localStorage.getItem('user');

        // check if the user is already signed in
        if (user != null) {
            alert("You are already logged in!");
            this.props.history.push('/');
        }
    };

    handleRegister(event) {
        event.preventDefault();
        const payload = {
            lastname : this.state.lastName,
            firstname : this.state.firstName,
            address : {
                street : this.state.street,
                city : this.state.city,
                state : this.state.state,
                zipcode : this.state.zipcode
            },
            password : this.state.password,
            email : this.state.email
        };
        UserAPI.callRegisterAPI(payload)
            .then((res) => {
                if (res.status === 201){
                    alert("user has been succesfully created");
                    this.props.history.push("/");
                }
                else if (res.status === 409){
                    alert("user with same email already exists");
                }
            })
            .catch((err) => {
                if (err.status === 409){
                    alert("user with same email already exists");
                }
                //console.log("some error occured :", err.error())
            })

    }

    render() {
        return (
            <div className="menu-home">
                <div className="outerdiv">
                    <Header/>
                    <div className="content payment">
                        <div className="card center">
                            <h2 id="center">Sign Up</h2>
                            {
                                <div>
                                    <form>
                                        <label>First Name</label>
                                        <input className="paymentInputText" type="text" id="firstname" name="FirstName" placeholder="First Name"
                                               onChange={
                                                   (event) => {
                                                       this.setState({
                                                           ...this.state,
                                                           firstName: event.target.value});
                                                   }
                                               }
                                        />
                                        <label>Last Name</label>
                                        <input className="paymentInputText" type="text" id="lastname" name="LastName" placeholder="Last Name"
                                               onChange={
                                                   (event) => {
                                                       this.setState({
                                                           ...this.state,
                                                           lastName: event.target.value});
                                                   }
                                               }
                                        />
                                        <label>street</label>
                                        <input className="paymentInputText" type="text" id="street" name="Street" placeholder="Street"
                                               onChange={
                                                   (event) => {
                                                       this.setState({
                                                           ...this.state,
                                                           street: event.target.value});
                                                   }
                                               }
                                        />
                                        <label>City</label>
                                        <input className="paymentInputText" type="text" id="city" name="City" placeholder="City"
                                               onChange={
                                                   (event) => {
                                                       this.setState({
                                                           ...this.state,
                                                           city: event.target.value});
                                                   }
                                               }
                                        />
                                        <label>State</label>
                                        <input className="paymentInputText" type="text" id="state" name="State" placeholder="State"
                                               onChange={
                                                   (event) => {
                                                       this.setState({
                                                           ...this.state,
                                                           state: event.target.value});
                                                   }
                                               }
                                        />
                                        <label>zipcode</label>
                                        <input className="paymentInputText" type="text" id="zipcode" name="Zipcode" placeholder="Zipcode"
                                               onChange={
                                                   (event) => {
                                                       this.setState({
                                                           ...this.state,
                                                           zipcode: event.target.value});
                                                   }
                                               }
                                        />
                                        <label>Email</label>
                                        <input className="paymentInputText" type="email" id="email" name="Email" placeholder="Email"
                                               onChange={
                                                   (event) => {
                                                       this.setState({
                                                           ...this.state,
                                                           email: event.target.value});
                                                   }
                                               }
                                        />
                                        <label>Password</label>
                                        <input className="paymentInputText" type="password" id="password" name="Password" placeholder="Password"
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
                                        <input type="button" className="payment_button" value="Create a new account"
                                               onClick={(e)=>{this.handleRegister(e)}}/>
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

// const mapStateToProps = (state) => {
//     return {
//     }
// };
//
// const mapDispatchToProps = (dispatch) => {
//     return {}
// };


const routerSignUp = withRouter(UserSignUp);

export default routerSignUp;
