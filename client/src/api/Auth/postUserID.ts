import  axios  from 'axios';

export function postUserID(userID: unknown): void {
    
        const backendUrl = process.env.REACT_APP_BACKEND_URL;
        console.log(userID);
        axios.post(`${backendUrl}/userID`, {
            userID: userID,
        })
            .then((response) => {
                const token = response.data;
                console.log("Received token:", token);
            })
            .catch((error) => {
                console.error(error);
            });
    }   


