import React, { useEffect, useState } from 'react';
import Index from '../components/Yumms/Index'
import { useSelector } from "react-redux";
import { getRecepies } from "../api/Yumms/Recepies";
import {AuthenticationUserStates, Recepie} from '../types/global'

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
        return <div className="section py-0">
            <div className="container mb-16">
                <div className="has-mw-2xl mx-auto has-text-centered">
                    <h2 className="mt-8 mb-10 title is-2">Loading...</h2>
                    <progress className="progress is-small is-primary" max="100">15%</progress>
                </div>
            </div>
        </div>
    }

    if (yumms1.length === 0) {
        return <div className="section py-0">
            <div className="container mb-16">
                <div className="has-mw-2xl mx-auto has-text-centered">
                    <h2 className="mt-8 mb-10 title is-2">No recipes have been added yet.</h2>
                    <progress className="progress is-small is-primary" max="100">15%</progress>
                </div>
            </div>
        </div>
    }

    return (
        <React.Fragment>
            <Index recepies={yumms1} />
            { yumms2.length > 0 && <Index recepies={yumms2} />}
        </React.Fragment>
    )
}
export default Yumms;