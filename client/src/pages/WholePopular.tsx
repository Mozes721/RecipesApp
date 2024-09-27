import React from 'react';
import Whole from '../components/Recepies/Whole'
import Popular from '../components/Recepies/Popular'
const WholePopular: React.FC = () => {
    return (
        <React.Fragment>
            <Whole />
            <Popular />
        </React.Fragment>
    )
}
export default WholePopular;