import { store } from '../store/store'

export function storeUser(userID: string, email: string | null ) {
    store.dispatch({
        type: 'SET_USER',
        payload: {
            userID,
            email,
        },
    });
}