import { Recepies } from '../../types/global';
import axios from 'axios';
export function getRecepies(userID: string | undefined): Promise<Recepies | void> {
    const serverGetUrl = process.env.REACT_APP_SERVER_GET;

    if (!serverGetUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_GET is not defined'));
    }
    if (!userID) {
        return Promise.reject(new Error('userID is not defined'));
    }


    return axios.get(serverGetUrl,  { params: { userID: userID } })
        .then((response) => {
            return response.data as Recepies
        })
        .catch((error) => {
            throw error;
        });
}