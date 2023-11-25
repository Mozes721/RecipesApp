import store from '../store/store';

export function getUserDetails() {
    store.dispatch({
        type: 'GET_USER',
    });
}
