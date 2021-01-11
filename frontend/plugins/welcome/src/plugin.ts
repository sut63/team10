import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';

import Table_patientrights from './components/Table_patientrights'
import create_Patientrights from './components/create_patientrights';
import Historytaking from './components/Historytaking';
import tableHistorytaking from './components/tableHistorytaking';
import createHistorytaking from './components/createHistorytaking';
import CreateBill from './components/createBill'
import createTreatment from './components/treatment';
import CreateDoctorinfo from './components/Doctorinfo';
import createPatientrecord from './components/createPatientrecord';
import Patientrecord from './components/Patientrecord';
import tablePatientrecord from './components/tablePatientrecord';
import Login from './components/Login';
import Logout from './components/Logout';
import reg from './components/reg';

import { Cookies } from 'react-cookie/cjs';//cookie
import React, { useEffect, FC } from 'react';




const cookies = new Cookies();
const Status = cookies.get('Status');
const Log = cookies.get('Log');




export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    if (Log == undefined) {
      router.registerRoute('/', Login);

    } else {
      router.registerRoute('/sign_out', Logout);

      if (Status == 'Root') {
        router.registerRoute('/reg', reg);
      }
      if (Status == 'Fin' || Status == 'Root' ) {
        router.registerRoute('/createBill', CreateBill);
      }
      if (Status == 'Med' || Status == 'Root' ) {
        router.registerRoute('/Table_patientrights',Table_patientrights);
        router.registerRoute('/create_Patientrights', create_Patientrights);
        router.registerRoute('/createPatientrecord', createPatientrecord);
        router.registerRoute('/Patientrecord', Patientrecord);
        router.registerRoute('/tablePatientrecord', tablePatientrecord);
      }
      if (Status == 'Nur' || Status == 'Root' ) {
        router.registerRoute('/Historytaking', Historytaking);
        router.registerRoute('/tableHistorytaking', tableHistorytaking);
        router.registerRoute('/createHistorytaking', createHistorytaking);
      }
      if (Status == 'Doc' || Status == 'Root' ) {
        router.registerRoute('/createTreatment', createTreatment);
      }
      
      if (Status == 'Reg' || Status == 'Root' ) {
        router.registerRoute('/Doctorinfo', CreateDoctorinfo);
      }

    
      router.registerRoute('/', WelcomePage);
    }





  },
});
