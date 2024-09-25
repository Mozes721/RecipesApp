import React from 'react';
import Whole from '../components/recepies/Whole'
import Popular from '../components/recepies/Popular'
const WholePopular: React.FC = () => {
    return (
        <React.Fragment>
            <Whole />
            <Popular />
        </React.Fragment>
    )
}
export default WholePopular;