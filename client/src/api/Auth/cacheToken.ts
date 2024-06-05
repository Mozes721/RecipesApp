import axios from "axios";

export function cacheUserToken(userID: string | undefined, authToken: string): Promise<void> {
    const serverCacheUrl = process.env.REACT_APP_SET_CACHE;

    if (!serverCacheUrl) {
        return Promise.reject(new Error('REACT_APP_SERVER_POST is not defined'));
    }

    return axios.post(`${serverCacheUrl}`, { UserID: userID, AuthToken: authToken })
        .then(() => {})
        .catch((error) => {
            throw new Error('Failed to cache user. Please try again.');
        });
}
