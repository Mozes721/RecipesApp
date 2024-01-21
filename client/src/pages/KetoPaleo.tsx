import React, { useEffect, useState } from 'react';
import Keto from '../components/recepies/Keto';
import Paelo from '../components/recepies/Paleo';

const KetoPaelo: React.FC = () => {
   return (
       <React.Fragment>
         <Keto />
         <Paelo />
       </React.Fragment>
   )
}
export default KetoPaelo;