import { store } from '../store/store';
import { getCacheUser } from "../api/Auth/getCacheUser";


export function getUserDetails(userID: string | undefined) {
    return getCacheUser(userID)
        .then(user => {
            store.dispatch({
                type: 'GET_USER',
                payload: {
                     user
                },
            });
        })
}
