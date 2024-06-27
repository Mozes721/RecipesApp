import { ReviewedRecepie } from '../../types/global';
import axios from 'axios';

export function updateRecepie(recipeReview: ReviewedRecepie, userID: string | undefined): Promise<boolean> {
    const serverUpdateUrl = process.env.REACT_APP_SERVER;

    if (!serverUpdateUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER is not defined'));
    }

    return axios.patch(serverUpdateUrl, recipeReview,  { params: { userID: userID }})
        .then(response => {
            return response.status === 200;
        })
        .catch((error) => {
            console.log(error)
            return false
        });
}