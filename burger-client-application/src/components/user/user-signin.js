
import React, {Component} from 'react';
// import {bindActionCreators} from 'redux';
// import {connect} from 'react-redux';
import {withRouter} from 'react-router-dom'
import '../../index.css';
import '../../stylesheets/payment.css';
import * as UserAPI from '../../apis/user-api'
import Header from '../header';

class UserSignIn extends Component {
    constructor(props) {
        super(props);
        this.state = {
            email:"",
            password:""
        };
        this.handleLogin = this.handleLogin.bind(this);
    }

    componentDidMount = () => {
        console.log('componentDidMount ---');
        const user = JSON.parse(localStorage.getItem('user'));

        // check if the user is already signed in
        if (user != null) {
            alert("You are already logged in!");
            this.props.history.push('/');
        }
    };

    handleLogin(event) {
        event.preventDefault();
        const payload = {
            email: this.state.email,
            password: this.state.password
        };
        UserAPI.callLoginApi(payload)
            .then((res) => {
                if (res.status === 200){
                    let user = {
                        email:res.data.email,
                        id: res.data.id,
                        name: res.data.firstName
                    };
                    localStorage.setItem('user', JSON.stringify(user));
                    this.props.history.push('/');
                }else if(res.status === 401){
                    alert("password does not match with given username")
                }
            })
            .catch((err) => {
                if(err.status === 401){
                    alert("password does not match with given username")
                }
                console.log("some error occured :", err)
            })

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

// const mapStateToProps = (state) => {
//     return {
//     }
// };
//
// const mapDispatchToProps = (dispatch) => {
//     return {}
// };


const routerSignIn = withRouter(UserSignIn);

export default routerSignIn;

