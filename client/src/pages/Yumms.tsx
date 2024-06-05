import React, { useEffect, useState } from 'react';
import Index from '../components/Yumms/Index'
import { useSelector } from "react-redux";
import { getRecepies } from "../api/Yumms/Recepies";
import {AuthenticationUserStates, Recepie, Recepies} from '../types/global'

const Yumms: React.FC = () => {
    const [yumms1, setYumms1] = useState<Recepie[]>([]);
    const [yumms2, setYumms2] = useState<Recepie[]>([]);
    const [loading, setLoading] = useState(true);
    const userID = useSelector((state: AuthenticationUserStates) => state.userID);

    useEffect(() => {
        getYumms();
    }, []);

    const getYumms = async () => {
        try {
            const data = await getRecepies(userID);
            if (data) {
                const half = Math.ceil(data.recepies.length / 2)
                const firstHalf = data.recepies.slice(0, half);
                const secondHalf = data.recepies.slice(half);
                setYumms1(firstHalf);
                setYumms2(secondHalf);
            }
        } catch (error) {
            console.error('Failed to fetch recipes:', error);
        } finally {
            setLoading(false);
        }
    };

    if (loading) {
        return <p>Loading...</p>;
    }

    if (yumms1.length === 0) {
        return <h2>No recipes have been added yet.</h2>;
    }

    return (
        <React.Fragment>
            <Index recepies={yumms1} />
            { yumms2.length > 0 && <Index recepies={yumms2} />}
        </React.Fragment>
    )
}
export default Yumms;