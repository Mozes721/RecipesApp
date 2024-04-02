import { store } from '../store/store';
import { AuthenticationUserStates } from '../types/global';
import { cacheUser } from "../api/Auth/cacheUser";
import { Simulate } from "react-dom/test-utils";
import error = Simulate.error;


export function storeUser(user: AuthenticationUserStates) {
    cacheUser(user)
        .then(cachedUser => {
            store.dispatch({
                type: 'SET_USER',
                payload: {
                    userID: user.userID,
                    email: user.email,
                    authToken: user.authToken
                },
            });
        })
        .catch(error => {
            console.error('Failed to cache user:', error);
            throw new Error('Failed to cache user. Please try again.');
        });
}