import  axios  from 'axios';

export function postUserID(userID: string) {
    const backendUrl = process.env.REACT_APP_BACKEND_URL;

    axios.post(`${backendUrl}/userID`, {
        userId: userID,
    })
        .then((response) => {
            const token = response.data;
            console.log("Received token:", token);
        })
        .catch((error) => {
            console.error(error);
        });
}

