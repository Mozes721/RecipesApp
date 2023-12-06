import {Recepie, User} from '../../types/global';
import axios from 'axios';

export function getRecepie(recipe: Recepie): Promise<void> {
    const servergetUrl = process.env.REACT_APP_SERVER_GET;

    if (!servergetUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_GET is not defined'));
    }

    return axios.get(servergetUrl, recipe)
        .then(() => {})
        .catch((error) => {
            throw error;
        });
}