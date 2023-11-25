import store from '../store/store'

export function removeUser() {
    store.dispatch({
        type: 'LOGOUT',
    });
}