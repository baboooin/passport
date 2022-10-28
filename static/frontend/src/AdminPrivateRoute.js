import { getToken, logout } from './Auth'; 
// import Api from './Api';
import React, {useState, useEffect} from 'react'
import {Route} from 'react-router-dom';



const AdminPrivateRoute = ({component: Component, ...rest}) => {
    useEffect(() => {
        // Here I make a fetch call to validate the token in the API
        // Api.checkToken(getToken() || '').then(res => {
            // if (!res.auth) {
                logout();
            // }
        // }
        // );
    })

    return (
        <Route
            {...rest}
            render={props =>
                isAuthenticated() ? (
                    <Component {...props} />
                ) : (
                    <SignIn />
                )
            }
        />
    );
};