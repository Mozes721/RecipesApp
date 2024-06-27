import axios from 'axios';

export function deleteRecepie(userID: string | undefined, recordID: string | undefined): Promise<boolean>  {
    const serverDeleteUrl = process.env.REACT_APP_SERVER;

    if (!serverDeleteUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER is not defined'));
    }

    return axios.delete(`${serverDeleteUrl}/${recordID}`, { params: { userID: userID } })
        .then(response => {
            return response.status === 200;
        })
        .catch((error) => {
            console.log(error)
            return false
        });
}