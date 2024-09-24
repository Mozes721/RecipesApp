import React from 'react';
import Keto from '../components/Recepies/Keto';
import Paleo from '../components/Recepies/Paleo';

const KetoPaelo: React.FC = () => {
   return (
       <React.Fragment>
         <Keto />
         <Paleo />
       </React.Fragment>
   )
} 
export default KetoPaelo;