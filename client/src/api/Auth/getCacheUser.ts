import { AuthenticationUserStates } from '../../types/global';
import axios from 'axios';



export function getCacheUser(userID: string | undefined): Promise<AuthenticationUserStates | void> {
    const serverGetCacheUrl = process.env.REACT_APP_GET_CACHE;

    if (!serverGetCacheUrl) {
        return Promise.reject(new Error('REACT_APP_GET_CACHE is not defined'));
    }


    return axios.get(`${serverGetCacheUrl}`, { params: { userID: userID }})
        .then((response) => {
            const getCachedUser: AuthenticationUserStates = {
                authenticated: response.data.Authenticated,
                authToken: response.data.AuthToken,
                userID: response.data.user.UserID,
                email: response.data.user.Email || ''
            };
            console.log(getCachedUser)

            return getCachedUser;
        })

}
