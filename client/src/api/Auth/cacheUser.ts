import { AuthenticationUserStates } from '../types/global';
import axios from 'axios';

export function cacheUser(cacheUser: AuthenticationUserStates): Promise<void> {
    const serverCacheUrl = process.env.REACT_APP_CACHE;

    if (!serverCacheUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_POST is not defined'));
    }

    return axios.post(serverCacheUrl, cacheUser)
        .then((response) => {
            
         
        })
        .catch((error) => {
            console.log(error)
        });
}