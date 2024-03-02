import { Recepie } from '../../types/global';
import axios from 'axios';

export function addNewRecepie(recipe: Recepie, token: string | undefined): Promise<{ status: number, message: string }> {
    const serverPostUrl = process.env.REACT_APP_SERVER_POST;

    if (!serverPostUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_POST is not defined'));
    }

    return axios.post(serverPostUrl, recipe, { headers:  {Authorization: `Bearer ${token}` } })
        .then((response) => {
            return {status: response.status, message: response.data.message}
         
        })
        .catch((error) => {
            return { status: error.response.status, message: error.response.data.message };
        });
}
