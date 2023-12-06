import { Recepie } from '../../types/global';
import axios from 'axios';

export function addRecepie(recipe: Recepie): Promise<void> {
    const serverPostUrl = process.env.REACT_APP_SERVER_POST;

    if (!serverPostUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_POST is not defined'));
    }

    return axios.post(serverPostUrl, recipe)
        .then(() => {})
        .catch((error) => {
            throw error;
        });
}
