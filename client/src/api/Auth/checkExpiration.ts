import axios from "axios";

export function checkSessionExpiration(userID: string | undefined): Promise<boolean>  {
    const serverExpirationUrl = process.env.REACT_APP_GET_EXPIRATION;

    if (!serverExpirationUrl) {
        return Promise.reject(new Error('REACT_APP_GET_EXPIRATION is not defined'));
    }

    return axios.get(serverExpirationUrl,  { params: { userID: userID } })
        .then((response) => {
            return response.data as boolean;
        })
        .catch((error) => {
            return false;
        });

}
