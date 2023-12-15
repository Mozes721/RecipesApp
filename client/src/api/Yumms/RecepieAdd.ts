import { Recepie } from '../../types/global';
import { GoogleAuthProvider, signInWithPopup, Auth } from 'firebase/auth';
import axios from 'axios';

export function addNewRecepie(recipe: Recepie, token: string | undefined): Promise<void> {
    const serverPostUrl = process.env.REACT_APP_SERVER_POST;

    if (!serverPostUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_POST is not defined'));
    }
    console.log(token)

    return axios.post(serverPostUrl, recipe, { headers:  {Authorization: `Bearer ${token}` } })
        .then(() => {})
        .catch((error) => {
            throw error;
        });
}
