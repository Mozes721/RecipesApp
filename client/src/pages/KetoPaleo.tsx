import React from 'react';
import Keto from '../components/recepies/Keto';
import Paleo from '../components/recepies/Paleo';

const KetoPaelo: React.FC = () => {
   return (
       <React.Fragment>
         <Keto />
         <Paleo />
       </React.Fragment>
   )
} 
export default KetoPaelo;