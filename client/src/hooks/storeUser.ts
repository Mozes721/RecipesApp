import store from '../store/store'

export function storeUser(userID: string, email: string | null , accessToken: string) {
    store.dispatch({
        type: 'SET_USER',
        payload: {
            userID,
            email,
            accessToken,
        },
    });
}