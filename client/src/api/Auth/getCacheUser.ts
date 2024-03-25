import { AuthenticationUserStates } from '../../types/global';
import axios from 'axios';



export function getCacheUser(userID: string | undefined): Promise<AuthenticationUserStates | void> {
    const serverGetCacheUrl = process.env.REACT_APP_GET_CACHE;

    if (!serverGetCacheUrl) {
        return Promise.reject(new Error('REACT_APP_GET_CACHE is not defined'));
    }

    if (userID === "" || userID === undefined )  {
        return Promise.resolve();
    }

    return axios.get(`${serverGetCacheUrl}?userID=${userID}`)
        .then((response) => {
            const getCachedUser: AuthenticationUserStates = {
                authenticated: response.data.authenticated,
                authToken: response.data.idToken,
                userID: response.data.user.uid,
                email: response.data.user.email || ''
            };

            return getCachedUser;
        })

}
