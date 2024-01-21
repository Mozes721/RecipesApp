import React from 'react';
import Vegan from '../components/recepies/Vegan'
import Vegetarian from '../components/recepies/Vegetarian'


const Veg: React.FC = () => {
    return (
        <React.Fragment>
            <Vegan />
            <Vegetarian />
        </React.Fragment>
    )
}
export default Veg;