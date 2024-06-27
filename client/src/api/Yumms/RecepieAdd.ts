import { Recepie } from '../../types/global';
import axios from 'axios';

export function addNewRecepie(recipe: Recepie, userID: string | undefined): Promise<{ status: number, message: string }> {
    const serverPostUrl = process.env.REACT_APP_SERVER;

    if (!serverPostUrl) {
        throw new Error('REACT_APP_SERVER is not defined');
    }
    return axios.post(serverPostUrl, recipe, { params: { userID: userID }})
        .then((response) => {
            return {status: response.status, message: response.data.message}

        })
        .catch((error) => {
            return { status: error.response.status, message: error.response.data.message };
        });
}

