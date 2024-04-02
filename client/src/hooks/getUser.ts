import { store } from '../store/store';
import { getCacheUser } from "../api/Auth/getCacheUser";


export function getUserDetails(userID: string | undefined) {
    console.log("Checking")
    const userData = getCacheUser(userID);
        store.dispatch({
            type: 'GET_USER',
            payload: {
                user: userData
            },
        });

}
