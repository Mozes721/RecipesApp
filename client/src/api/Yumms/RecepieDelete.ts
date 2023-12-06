import { Recepie } from '../../types/global';
import axios from 'axios';

export function deleteRecepie(recipe: Recepie): Promise<void> {
    const serverDeleteUrl = process.env.REACT_APP_SERVER_PATCH_DELETE;

    if (!serverDeleteUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_POST is not defined'));
    }

    return axios.delete(serverDeleteUrl, recipe)
        .then(() => {})
        .catch((error) => {
            throw error;
        });

}