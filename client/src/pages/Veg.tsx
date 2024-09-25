import React from 'react';
import Vegan from '../components/Recepies/Vegan'
import Vegetarian from '../components/Recepies/Vegetarian'


const Veg: React.FC = () => {
    return (
        <React.Fragment>
            <Vegan />
            <Vegetarian />
        </React.Fragment>
    )
}
export default Veg;