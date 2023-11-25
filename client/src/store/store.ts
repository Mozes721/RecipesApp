import { createStore, Reducer } from 'redux';
import AuthenticationUserStates from '../types/global.d';


const initialState: AuthenticationUserStates = {
    authenticated: false,
    authToken: '',
    userID: '',
    email: '',
};

// Define your reducer function
const authReducer: Reducer<AuthenticationUserStates, any> = (
    state = initialState,
    action: any
) => {
    switch (action.type) {
        case 'SET_USER':
            console.log("SET_USER action dispatched");
            return {
                authenticated: true,
                authToken: action.payload.accessToken,
                userID: action.payload.userID,
                email: action.payload.email,
            };
        case 'GET_USER':
            return {
                ...state,
            };
        case 'LOGOUT':
            return initialState;
        default:
            return state;
    }
};

// Create the Redux store
const store = createStore(authReducer);

export default store;
