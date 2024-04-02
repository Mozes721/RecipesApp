import { AuthenticationUserStates } from '../../types/global';
import axios from 'axios';

export function cacheUser(cacheUser: AuthenticationUserStates): Promise<axios.AxiosResponse<any> | void> {
    const serverCacheUrl = process.env.REACT_APP_SET_CACHE;

    if (!serverCacheUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_POST is not defined'));
    }

    return axios.post(serverCacheUrl, cacheUser)
        .then((response) => {
            return response.data;
        })
        .catch((error) => {
            console.error('Failed to cache user:', error);
            throw new Error('Failed to cache user. Please try again.');
        });
}