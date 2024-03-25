import { store } from '../store/store';
import { AuthenticationUserStates } from '../types/global';
import { cacheUser } from "../api/Auth/cacheUser";
import { Simulate } from "react-dom/test-utils";
import error = Simulate.error;


export function storeUser(user: AuthenticationUserStates) {
    const {authToken, userID, email} =  user
    cacheUser(user).then(r => {})
        .catch(error)
    store.dispatch({
        type: 'SET_USER',
        payload: {
            userID,
            email,
            authToken,
        },
    });
}