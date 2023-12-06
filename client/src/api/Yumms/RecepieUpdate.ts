import { Recepie } from '../../types/global';
import axios from 'axios';

export function updateRecepie(recipe: Recepie): Promise<void> {
    const serverUpdateUrl = process.env.REACT_APP_SERVER_PATCH_DELETE;

    if (!serverUpdateUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_PATCH_DELETE is not defined'));
    }

    return axios.post(serverUpdateUrl, recipe)
        .then(() => {})
        .catch((error) => {
            throw error;
        });
}