import { Recepies } from '../../types/global';
import axios from 'axios';

export function getRecepie(userID: string | undefined): Promise<Recepies | void> {
    const servergetUrl = process.env.REACT_APP_SERVER_GET;

    if (!servergetUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_GET is not defined'));
    }

     const url = `${servergetUrl}/${userID}`;

    return axios.get(url, { params: { userID: userID }})
        .then((response) => {

            return response.data as Recepies
        })
        .catch((error) => {
            throw error;
        });
}