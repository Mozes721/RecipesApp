import React from 'react';
import Vegan from '../components/recepies/Vegan'
import Vegetarian from '../components/recepies/Vegetarian'


const VegPage: React.FC = () => {
    return (
        <React.Fragment>
            <Vegan />
            <Vegetarian />
        </React.Fragment>
    )
}
export default VegPage;