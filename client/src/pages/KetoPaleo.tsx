import React, { useEffect, useState } from 'react';
import Keto from '../components/Recepies/Keto';
import Paelo from '../components/Recepies/Paleo';

const KetoPaelo: React.FC = () => {
   return (
       <React.Fragment>
         <Keto />
         <Paelo />
       </React.Fragment>
   )
}
export default KetoPaelo;